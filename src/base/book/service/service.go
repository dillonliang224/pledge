package service

import (
	"git.dillonliang.cn/micro-svc/pledge/src/base/book/config"
	"git.dillonliang.cn/micro-svc/pledge/src/base/book/dao"
	"go.uber.org/zap"
)

type Service struct {
	dao    *dao.Dao
	logger *zap.Logger
}

func New(cfg *config.Config, logger *zap.Logger) *Service {
	return &Service{
		dao:    dao.New(cfg),
		logger: logger,
	}
}

func (s *Service) Close() {
	s.dao.Close()
}
