package service

import (
	"context"

	"git.dillonliang.cn/micro-svc/pledge/src/web/goldcoin/conf"
	"git.dillonliang.cn/micro-svc/pledge/src/web/goldcoin/dao"
)

type Service struct {
	dao *dao.Dao
}

func New(c *conf.Config) *Service {
	return &Service{
		dao: dao.New(c),
	}
}

func (s *Service) Ping(ctx context.Context) error {
	return s.dao.Ping(ctx)
}

func (s *Service) Close() {
	s.dao.Close()
}
