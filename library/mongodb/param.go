package mongodb

import (
	"context"
	"time"
)

type params struct {
	sort       interface{}
	skip       *int64
	limit      *int64
	multi      bool
	upsert     *bool
	projection interface{}
	ctx        context.Context
}

type Param func(*params)

func evaluateParam(param []Param) *params {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	ps := &params{
		ctx: ctx,
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
