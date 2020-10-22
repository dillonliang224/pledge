package discovery

import (
	"context"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc/resolver"
)

const (
	schema = "etcd"
)

type Resolver struct {
	schema      string
	Endpoints   []string // etcd节点
	DialTimeout int

	closeCh      chan struct{}
	watchCh      clientv3.WatchChan
	client       *clientv3.Client
	keyPrefix    string
	srvAddrsList []resolver.Address // 服务地址列表

	cc     resolver.ClientConn
	logger *zap.Logger
}

func NewResolver(endpoints []string, logger *zap.Logger) *Resolver {
	return &Resolver{
		schema:      schema,
		Endpoints:   endpoints,
		DialTimeout: 3,
		logger:      logger,
	}
}

func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r.cc = cc
	r.keyPrefix = BuildPrefix(Server{Name: target.Endpoint})

	if _, err := r.start(); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *Resolver) Scheme() string {
	return r.schema
}

func (r *Resolver) ResolveNow(options resolver.ResolveNowOptions) {

}

func (r *Resolver) Close() {
	r.closeCh <- struct{}{}
}

func (r *Resolver) start() (chan<- struct{}, error) {
	var err error
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   r.Endpoints,
		DialTimeout: time.Duration(r.DialTimeout) * time.Second,
	})

	if err != nil {
		return nil, err
	}
	r.client = client

	resolver.Register(r)

	r.closeCh = make(chan struct{})
	go r.watch()

	return r.closeCh, nil
}

func (r *Resolver) watch() {
	ticker := time.NewTicker(time.Minute)
	r.watchCh = r.client.Watch(context.Background(), r.keyPrefix, clientv3.WithPrefix())

	for {
		select {
		case <-r.closeCh:
			return
		case res, ok := <-r.watchCh:
			if ok {
				if ok {
					r.update(res.Events)
				}
			}
		case <-ticker.C:
			// TODO SYNC
		}
	}
}

func (r *Resolver) update(events []*clientv3.Event) {
	for _, ev := range events {
		var info Server
		var err error

		switch ev.Type {
		case clientv3.EventTypePut:
			info, err = ParseValue(ev.Kv.Value)
			if err != nil {
				continue
			}

			addr := resolver.Address{Addr: info.Addr, Metadata: info.Weight}
			if !Exist(r.srvAddrsList, addr) {
				r.srvAddrsList = append(r.srvAddrsList, addr)
				r.cc.UpdateState(resolver.State{Addresses: r.srvAddrsList})
			}
		case clientv3.EventTypeDelete:
			info, err = SplitPath(string(ev.Kv.Key))
			if err != nil {
				continue
			}

			addr := resolver.Address{Addr: info.Addr, Metadata: info.Weight}
			if s, ok := Remove(r.srvAddrsList, addr); ok {
				r.srvAddrsList = s
				r.cc.UpdateState(resolver.State{Addresses: r.srvAddrsList})
			}
		}
	}
}

func (r *Resolver) sync() {

}
