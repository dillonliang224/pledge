package mongodb

import (
	"context"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type params struct {
	sort       interface{}
	skip       *int64
	limit      *int64
	multi      bool
	upsert     *bool
	projection interface{}
	ctx        context.Context
	cancel     context.CancelFunc
}

type Param func(*params)

func evaluateParam(param []Param) *params {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	ps := &params{
		ctx:    ctx,
		cancel: cancel,
	}
	for _, p := range param {
		p(ps)
	}
	return ps
}

func Context(ctx context.Context) Param {
	return func(o *params) {
		o.ctx = ctx
	}
}

func Sort(sort interface{}) Param {
	return func(o *params) {
		o.sort = sort
	}
}

func Skip(skip int64) Param {
	return func(o *params) {
		o.limit = &skip
	}
}

func Limit(limit int64) Param {
	return func(o *params) {
		o.limit = &limit
	}
}

func Multi() Param {
	return func(o *params) {
		o.multi = true
	}
}

func Upsert() Param {
	return func(o *params) {
		u := true
		o.upsert = &u
	}
}

func Projection(projection interface{}) Param {
	return func(o *params) {
		o.projection = projection
	}
}

func Select(fields []string) Param {
	return func(o *params) {
		pj := bson.M{}
		for _, f := range fields {
			pj[f] = true
		}
		o.projection = pj
	}
}

func Exclude(fields []string) Param {
	return func(o *params) {
		pj := bson.M{}
		for _, f := range fields {
			pj[f] = false
		}
		o.projection = pj
	}
}
