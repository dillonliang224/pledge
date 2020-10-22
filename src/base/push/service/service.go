package service

import (
	"go.uber.org/zap"

	"git.dillonliang.cn/micro-svc/pledge/src/base/push/conf"
	"git.dillonliang.cn/micro-svc/pledge/src/base/push/dao"
)

type Service struct {
	dao    *dao.Dao
	logger *zap.Logger
	cfg    *conf.Config
}

func New(cfg *conf.Config, logger *zap.Logger) *Service {
	return &Service{
		dao:    dao.New(cfg),
		logger: logger,
		cfg:    cfg,
	}
}

func (s *Service) Close() {
	s.dao.Close()
}
