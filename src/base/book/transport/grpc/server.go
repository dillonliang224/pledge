package grpc

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "git.dillonliang.cn/micro-svc/pledge/src/base/book/api"
	"git.dillonliang.cn/micro-svc/pledge/src/base/book/config"
	"git.dillonliang.cn/micro-svc/pledge/src/base/book/service"
)

func Start(cfg *config.Config, service *service.Service) (*grpc.Server, error) {
	srv := grpc.NewServer()
	pb.RegisterBooksServer(srv, &server{as: service})
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

func (s *server) FindById(ctx context.Context, req *pb.FindByIdReq) (*pb.BookResp, error) {
	resp, err := s.as.FindById(ctx, req.Id)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return resp, nil
}
