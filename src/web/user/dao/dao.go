package dao

import (
	"context"

	"git.dillonliang.cn/micro-svc/pledge/library/mongodb"
	"git.dillonliang.cn/micro-svc/pledge/src/web/user/conf"
	"git.dillonliang.cn/micro-svc/pledge/src/web/user/model"
)

type Dao struct {
	mongoConn *mongodb.Connection
	user      *mongodb.Model
}

func New(cfg *conf.Config) (dao *Dao) {
	conn, err := mongodb.Connect(cfg.Mongodb.URL)
	if err != nil {
		panic("conn db error")
	}

	dao = &Dao{
		mongoConn: conn,
		user:      conn.Model(model.TableUser),
	}

	return
}

func (d *Dao) Ping(ctx context.Context) error {
	return d.mongoConn.HealthCheck()
}

func (d *Dao) Close() {

}
