package mongodb

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Model struct {
	collection *mongo.Collection
}

var (
	CollectionError = errors.New("collection not found")
)

func (m *Model) FindOne(result, query interface{}, param ...Param) error {
	if m.collection == nil {
		return CollectionError
	}

	ps := evaluateParam(param)

	err := m.collection.FindOne(ps.ctx, query, &options.FindOneOptions{
		Sort:       ps.sort,
		Projection: ps.projection,
	}).Decode(result)
	return err
}

// 反射的应用呀
func (m *Model) Find(result, query interface{}, param ...Param) error {
	if m.collection == nil {
		return CollectionError
	}

	ps := evaluateParam(param)

	resultv := reflect.ValueOf(result)
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		panic("result argument must be a slice address")
	}
	slicev := resultv.Elem()
	elemt := slicev.Type().Elem()

	cursor, err := m.collection.Find(ps.ctx, query, &options.FindOptions{
		Sort:       ps.sort,
		Skip:       ps.skip,
		Limit:      ps.limit,
		Projection: ps.projection,
	})
	if err != nil {
		return err
	}
	defer cursor.Close(ps.ctx)

	for cursor.Next(ps.ctx) {
		elemp := reflect.New(elemt)
		if err := cursor.Decode(elemp.Interface()); err == nil {
			slicev = reflect.Append(slicev, elemp.Elem())
		}
	}
	if err := cursor.Err(); err != nil {
		return err
	}

	resultv.Elem().Set(slicev)
	return nil
}

func (m *Model) Insert(doc interface{}, param ...Param) (interface{}, error) {
	if m.collection == nil {
		return nil, CollectionError
	}

	ps := evaluateParam(param)
	result, err := m.collection.InsertOne(ps.ctx, doc)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (m *Model) Update(query, update interface{}, param ...Param) (*mongo.UpdateResult, error) {
	if m.collection == nil {
		return nil, CollectionError
	}

	ps := evaluateParam(param)

	if ps.multi {
		updateRes, err := m.collection.UpdateMany(ps.ctx, query, update, &options.UpdateOptions{Upsert: ps.upsert})
		return updateRes, err
	}

	updateRes, err := m.collection.UpdateOne(ps.ctx, query, update, &options.UpdateOptions{Upsert: ps.upsert})
	return updateRes, err
}

func (m *Model) Delete(query interface{}, param ...Param) error {
	if m.collection == nil {
		return CollectionError
	}

	ps := evaluateParam(param)

	_, err := m.collection.DeleteMany(ps.ctx, query)
	return err
}

func (m *Model) Count(query interface{}, param ...Param) int64 {
	if m.collection == nil {
		return 0
	}

	ps := evaluateParam(param)
	count, _ := m.collection.CountDocuments(ps.ctx, query)
	return count
}

func (m *Model) Aggregate(pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, context.Context, error) {
	if m.collection == nil {
		return nil, nil, CollectionError
	}

	ps := evaluateParam(nil)
	cursor, err := m.collection.Aggregate(ps.ctx, pipeline, opts...)
	return cursor, ps.ctx, err
}
