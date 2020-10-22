package discovery

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
)

type Register struct {
	Endpoints   []string
	DialTimeout int

	leaseId     clientv3.LeaseID
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse
	closeCh     chan struct{}

	srvInfo Server
	srvTTL  int64
	client  *clientv3.Client
	logger  *zap.Logger
}

// base on etcd
func NewRegister(endpoints []string, logger *zap.Logger) *Register {
	return &Register{
		Endpoints:   endpoints,
		DialTimeout: 3,
		logger:      logger,
	}
}

func (r *Register) Register(srvInfo Server, ttl int64) (chan<- struct{}, error) {
	var err error

	if strings.Split(srvInfo.Addr, ":")[0] == "" {
		return nil, errors.New("invalid ip")
	}

	// 连接etcd
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   r.Endpoints,
		DialTimeout: time.Duration(r.DialTimeout) * time.Second,
	})

	if err != nil {
		return nil, err
	}

	r.client = client
	r.srvInfo = srvInfo
	r.srvTTL = ttl

	// 注册节点信息
	if err = r.register(); err != nil {
		return nil, err
	}

	r.closeCh = make(chan struct{})
	// 保持连接
	go r.keepAlive()

	return r.closeCh, err
}

// 注册节点
func (r *Register) register() error {
	// 注册节点
	leaseCtx, cancel := context.WithTimeout(context.Background(), time.Duration(r.DialTimeout)*time.Second)
	defer cancel()

	// 创建一个租约
	leaseResp, err := r.client.Grant(leaseCtx, r.srvTTL)
	if err != nil {
		return err
	}

	r.leaseId = leaseResp.ID
	keepAliveCh, err := r.client.KeepAlive(context.Background(), r.leaseId)
	if err != nil {
		return err
	}

	r.keepAliveCh = keepAliveCh

	// 把自身的服务信息写入etcd
	data, err := json.Marshal(r.srvInfo)
	if err != nil {
		return err
	}

	_, err = r.client.Put(context.Background(), BuildPrefix(r.srvInfo), string(data), clientv3.WithLease(r.leaseId))

	return err
}

func (r *Register) unRegister() error {
	_, err := r.client.Delete(context.Background(), BuildPrefix(r.srvInfo))
	return err
}

// 停止服务注册
func (r *Register) Stop() {
	r.closeCh <- struct{}{}
}

func (r *Register) keepAlive() {
	ticker := time.NewTicker(time.Duration(r.srvTTL) * time.Second)
	for {
		select {
		case <-r.closeCh:
			if err := r.unRegister(); err != nil {
				r.logger.Error("unRegister failed", zap.Error(err))
			}

			if err := r.revoke(); err != nil {
				r.logger.Error("revoke failed", zap.Error(err))
			}
			return
		case res := <-r.keepAliveCh:
			if res == nil {
				if err := r.register(); err != nil {
					r.logger.Error("register failed", zap.Error(err))
				}
			}
		case <-ticker.C:
			if r.keepAliveCh == nil {
				if err := r.register(); err != nil {
					r.logger.Error("register failed", zap.Error(err))
				}
			}
		}
	}
}

func (r *Register) revoke() error {
	_, err := r.client.Revoke(context.Background(), r.leaseId)
	if err != nil {
		return err
	}
	return nil
}

func (r *Register) GetServerInfo() (Server, error) {
	resp, err := r.client.Get(context.Background(), BuildPrefix(r.srvInfo))
	if err != nil {
		return r.srvInfo, err
	}

	info := Server{}
	if resp.Count >= 1 {
		if err := json.Unmarshal(resp.Kvs[0].Value, &info); err != nil {
			return info, err
		}
	}

	return info, nil
}
