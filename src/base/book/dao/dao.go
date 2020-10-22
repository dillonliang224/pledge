package dao

import (
	"context"

	"git.dillonliang.cn/micro-svc/pledge/library/mongodb"
	"git.dillonliang.cn/micro-svc/pledge/src/base/book/config"
	"git.dillonliang.cn/micro-svc/pledge/src/base/book/model"
)

type Dao struct {
	mongoConn   *mongodb.Connection
	book        *mongodb.Model
	cpSourceMap *mongodb.Model
	cpChapter   *mongodb.Model
	toc         *mongodb.Model
}

func New(cfg *config.Config) (dao *Dao) {
	conn, err := mongodb.Connect(cfg.Mongodb.URL)
	if err != nil {
		panic("conn db error")
	}

	dao = &Dao{
		mongoConn:   conn,
		book:        conn.Model(model.TableBook),
		cpSourceMap: conn.Model(model.TableCpSourceMap),
		cpChapter:   conn.Model(model.TableCpChapter),
		toc:         conn.Model(model.TableToc),
	}

	return
}

func (d *Dao) Ping(ctx context.Context) error {
	return d.mongoConn.HealthCheck()
}

func (d *Dao) Close() {

}
