package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CpChapter struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Cp          string             `bson:"cp,omitempty" json:"cp"`
	CpChapterId string             `bson:"cpChapterId,omitempty" json:"cpChapterId"`
	Title       string             `bson:"title,omitempty" json:"title"`
	Order       int32              `bson:"order,omitempty" json:"order"`
	WordCount   int32              `bson:"wordCount,omitempty" json:"wordCount"`
	Content     string             `bson:"content,omitempty" json:"content"`
	Updated     *time.Time         `bson:"updated,omitempty" json:"updated"`
}

const TableCpChapter = "cpChapters"
