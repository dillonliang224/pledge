package dao

import (
	"context"

	"git.dillonliang.cn/micro-svc/pledge/library/mongodb"
	"git.dillonliang.cn/micro-svc/pledge/src/base/push/conf"
	"git.dillonliang.cn/micro-svc/pledge/src/base/push/model"
)

type Dao struct {
	mongoConn *mongodb.Connection
	smsCode   *mongodb.Model
	smsSign   *mongodb.Model
}

func New(cfg *conf.Config) (dao *Dao) {
	conn, err := mongodb.Connect(cfg.Mongodb.URL)
	if err != nil {
		panic("conn db error")
	}

	dao = &Dao{
		mongoConn: conn,
		smsCode:   conn.Model(model.TableSmsCode),
		smsSign:   conn.Model(model.TableSmsSign),
	}

	return
}

func (d *Dao) Ping(ctx context.Context) error {
	return d.mongoConn.HealthCheck()
}

func (d *Dao) Close() {

}
