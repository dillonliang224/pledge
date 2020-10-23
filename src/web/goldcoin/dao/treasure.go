package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"git.dillonliang.cn/micro-svc/pledge/library/mongodb"
	"git.dillonliang.cn/micro-svc/pledge/src/web/goldcoin/model"
)

func (d *Dao) FindUserTreasureRecords() {

}

func (d *Dao) FindTreasureProductById(ctx context.Context, id primitive.ObjectID) (*model.TreasureProduct, error) {
	product := new(model.TreasureProduct)
	query := bson.M{
		"_id": id,
	}
	err := d.schema[model.TableTreasureProduct].FindOne(product, query, mongodb.Context(ctx))
	return product, err
}

func (d *Dao) FindTreasureRecords(ctx context.Context, activityId primitive.ObjectID, productId primitive.ObjectID,
	period string) ([]*model.TreasureRecord, error) {
	records := make([]*model.TreasureRecord, 0)
	query := bson.M{
		"activityId": activityId,
		"productId":  productId,
		"period":     period,
	}
	err := d.schema[model.TableTreasureRecord].Find(&records, query,
		mongodb.Context(ctx),
		// mongodb.Select([]string{"user", "codes", "created"}),
		mongodb.Sort(bson.M{"_id": -1}),
		mongodb.Limit(10))
	return records, err
}
