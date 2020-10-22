package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"git.dillonliang.cn/micro-svc/pledge/library/mongodb"
	"git.dillonliang.cn/micro-svc/pledge/src/web/goldcoin/model"
)

func (d *Dao) FindUserTreasureRecord() {

}

func (d *Dao) FindTreasureProductById(ctx context.Context, id primitive.ObjectID) (*model.TreasureProduct, error) {
	product := new(model.TreasureProduct)
	query := bson.M{
		"_id": id,
	}
	err := d.schema[model.TableTreasureProduct].FindOne(product, query, mongodb.Context(ctx))
	return product, err
}
