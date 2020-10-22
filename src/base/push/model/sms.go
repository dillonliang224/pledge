package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SmsCode struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Group      string             `bson:"group,omitempty" json:"group"`           // group
	Mobile     int64              `bson:"mobile,omitempty" json:"mobile"`         // 手机号
	Code       string             `bson:"code,omitempty" json:"code"`             // 验证码
	Content    string             `bson:"content,omitempty" json:"content"`       // 短信内容
	Platform   string             `bson:"platform,omitempty" json:"platform"`     // 使用的第三方平台
	SmsType    int                `bson:"type" json:"type"`                       // 短信类型
	ExpireTime time.Time          `bson:"expireTime,omitempty" json:"expireTime"` // 过期时间
	Created    time.Time          `bson:"created,omitempty" json:"created"`       // 创建时间
	ReturnCode string             `bson:"returnCode" json:"returnCode"`           // 第三方流失号
	UserId     string             `bson:"userId,omitempty" json:"userId"`         // 用户ID
}

type SmsSign struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Group    string             `bson:"group,omitempty" json:"group"`       // group，appName太局限
	Platform string             `bson:"platform,omitempty" json:"platform"` // 使用的第三方平台
	SmsType  int                `bson:"type,omitempty" json:"type"`         // 短信类型
	CodeNum  int                `bson:"codeNum,omitempty" json:"codeNum"`   // 短信验证码长度
	Sign     string             `bson:"sign,omitempty" json:"sign"`         // 短信签名
	Template string             `bson:"template,omitempty" json:"template"` // 短信模版
	UserName string             `bson:"userName,omitempty" json:"userName"` // 短信密码
	Password string             `bson:"password,omitempty" json:"password"` // 短信密码
}

const TableSmsCode = "dl-smscodes"
const TableSmsSign = "dl-smssigns"
