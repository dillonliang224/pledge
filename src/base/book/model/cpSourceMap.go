package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CpSourceMap struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Book     primitive.ObjectID `bson:"book,omitempty" json:"book"`
	Source   string             `bson:"source,omitempty" json:"source"`
	RemoteId string             `bson:"remoteId,omitempty" json:"remoteId"`
	Updated  *time.Time         `bson:"updated,omitempty" json:"omitempty"`
}

const TableCpSourceMap = "cpSourceMaps"
