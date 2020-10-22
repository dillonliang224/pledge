package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const TableAccount = "accounts"

type Account struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	User primitive.ObjectID `bson:"user,omitempty" json:"user"`
	Code int64              `bson:"code,omitempty" json:"code"`

	Name    string `bson:"name,omitempty" json:"name"`       // 真实姓名
	CardNo  string `bson:"cardNo,omitempty" json:"cardNo"`   // 身份证号
	Phone   string `bson:"phone,omitempty" json:"phone"`     // 手机号
	Address string `bson:"address,omitempty" json:"address"` // 用户收货地址

	Updated *time.Time `bson:"updated,omitempty" json:"updated"` // 更新时间
	Created time.Time  `bson:"created,omitempty" json:"created"` // 创建时间
}
