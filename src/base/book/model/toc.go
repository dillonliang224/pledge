package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Chapter struct {
	ID    primitive.ObjectID `bson:"id,omitempty" json:"id"`
	Title string             `bson:"title,omitempty" json:"title"`
	Order int32              `bson:"order,omitempty" json:"order"`
}

type Toc struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Book          primitive.ObjectID `bson:"book,omitempty" json:"book"`
	Source        string             `bson:"source,omitempty" json:"source"`
	IsFirsthand   bool               `bson:"isFirsthand,omitempty" json:"isFirsthand"`
	Chapters      []Chapter          `bson:"chapters,omitempty" json:"chapters"`
	LastChapter   string             `bson:"lastChapter,omitempty" json:"lastChapter"`
	ChaptersCount int32              `bson:"chaptersCount,omitempty" json:"chaptersCount"`
	Updated       time.Time          `bson:"updated,omitempty" json:"updated"`
}

const TableToc = "tocs"
