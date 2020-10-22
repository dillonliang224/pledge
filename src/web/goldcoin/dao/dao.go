package dao

import (
	"context"

	"git.dillonliang.cn/micro-svc/pledge/library/mongodb"
	"git.dillonliang.cn/micro-svc/pledge/src/web/goldcoin/conf"
	"git.dillonliang.cn/micro-svc/pledge/src/web/goldcoin/model"
)

var tables = []string{
	model.TableTreasureActivity,
	model.TableTreasureProduct,
	model.TableTreasureRecord,
	model.TableTreasureWinner,
	model.TableTreasurePeriod,
	model.TableTreasureFreeChance,
	model.TableTreasureRobot,
}

type Dao struct {
	schema map[string]*mongodb.Model
}

func New(c *conf.Config) *Dao {
	db, err := mongodb.Connect(c.Mongodb.URL)
	if err != nil {
		panic(err)
	}

	schema := make(map[string]*mongodb.Model)

	for _, name := range tables {
		schema[name] = db.Model(name)
	}

	return &Dao{
		schema: schema,
	}
}

func (d *Dao) Ping(ctx context.Context) (err error) {
	return
}

func (d *Dao) Close() {
}
