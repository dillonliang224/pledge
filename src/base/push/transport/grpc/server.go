package grpc

import (
	"context"
	"net"

	"google.golang.org/grpc"

	pb "git.dillonliang.cn/micro-svc/pledge/src/base/push/api"
	"git.dillonliang.cn/micro-svc/pledge/src/base/push/conf"
	"git.dillonliang.cn/micro-svc/pledge/src/base/push/service"
)

func Start(cfg *conf.Config, service *service.Service) (*grpc.Server, error) {
	srv := grpc.NewServer()
	pb.RegisterPushServer(srv, &server{as: service})
	// 开启反射
	// reflection.Register(srv)

	lis, err := net.Listen("tcp", cfg.Port.GRPC)
	if err != nil {
		return nil, err
	}

	go func() {
		if err := srv.Serve(lis); err != nil {
			panic(err)
		}
	}()

	return srv, nil
}

type server struct {
	as *service.Service
}

func (s server) SendSms(ctx context.Context, req *pb.SendSmsReq) (*pb.SendSmsRes, error) {
	err := s.as.SendSms(req.Group, req.Mobile, int(req.Type))
	ok := true
	if err != nil {
		ok = false
	}
	return &pb.SendSmsRes{
		Ok: ok,
	}, err
}

func (s server) CheckSmsCode(ctx context.Context, req *pb.CheckSmsCodeReq) (*pb.CheckSmsCodeRes, error) {
	ok, err := s.as.CheckSmsCode(req.Mobile, req.Code, int(req.Type))
	if err != nil {
		return &pb.CheckSmsCodeRes{
			Ok: false,
		}, err
	}

	return &pb.CheckSmsCodeRes{
		Ok: ok,
	}, err
}
