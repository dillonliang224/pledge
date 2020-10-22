package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Platform string

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"_id"`
	Nickname string             `bson:"nickname" json:"nickname"`
	Avatar   string             `bson:"avatar" json:"avatar"`
	Gender   string             `bson:"gender" json:"gender"`
	Exp      int                `bson:"exp" json:"exp"`
	Lv       int                `bson:"lv" json:"lv"`
	Bind     map[Platform]*Bind `bson:"bind" json:"bind"`
}

type Bind struct {
	Uid     string `bson:"uid,omitempty" json:"uid"`
	Name    string `bson:"name,omitempty" json:"name"`
	UnionId string `bson:"unionId,omitempty" json:"unionId"`
	Mobile  int64  `bson:"mobile,omitempty" json:"mobile"`
}

const TableUser = "dl-users"
