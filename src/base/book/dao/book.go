package dao

import (
	"context"

	"git.dillonliang.cn/micro-svc/pledge/src/base/book/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (dao *Dao) FindBookInfo(ctx context.Context, bookId primitive.ObjectID) (*model.Book, error) {
	query := bson.M{
		"_id": bookId,
	}

	var book model.Book
	if err := dao.book.FindOne(&book, query); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &book, nil
}
