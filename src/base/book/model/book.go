package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Title         string             `bson:"title,omitempty" json:"title"`
	Author        string             `bson:"author,omitempty" json:"author"`
	Cover         string             `bson:"cover,omitempty" json:"cover"`
	ShortIntro    string             `bson:"shortIntro,omitempty" json:"shortIntro"`
	LongIntro     string             `bson:"longIntro,omitempty" json:"longIntro"`
	LastChapter   *string            `bson:"lastChapter,omitempty" json:"lastChapter"`
	ChaptersCount int64              `bson:"chaptersCount,omitempty" json:"chaptersCount"`
	IsSerial      bool               `bson:"isSerial,omitempty" json:"isSerial"`
	Tags          []string           `bson:"tags,omitempty" json:"tags"`
	Updated       *time.Time         `bson:"updated,omitempty" json:"updated"`
}

const TableBook = "books"
