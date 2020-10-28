package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 域名管理
type Domain struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Hostname   string             `bson:"hostname,omitempty" json:"hostname"`     // 域名
	Enabled    bool               `bson:"enabled,omitempty" json:"enabled"`       // 是否生效
	IsInEffect bool               `bson:"isInEffect,omitempty" json:"isInEffect"` // 是否可用
	Created    *time.Time         `bson:"created,omitempty" json:"created"`       // 创建时间
}

// 记录扫码信息
type SpringBoard struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	IP     string             `bson:"ip,omitempty" json:"ip"`         // 扫码时IP
	IPv4   string             `bson:"IPv4,omitempty" json:"IPv4"`     // 扫码时IP
	UA     string             `bson:"UA,omitempty" json:"UA"`         // 扫码时UA
	OS     string             `bson:"OS,omitempty" json:"OS"`         // 扫码时OS
	Device string             `bson:"device,omitempty" json:"device"` // 扫码Devicd
	Group  string             `bson:"group,omitempty" json:"group"`
}

type ContentType string

const (
	IMAGE       ContentType = "IMAGE"       // eg: 分享的头像
	TEXT        ContentType = "TEXT"        // eg: 分享的昵称
	QRCODE_TEXT ContentType = "QRCODE_TEXT" // eg: 分享的二维码
)

type AlignType string

const (
	LEFT   AlignType = "left"
	RIGHT  AlignType = "right"
	CENTER AlignType = "center"
)

type Watermark struct {
	From        string      `bson:"from,omitempty" json:"from"`               // 内容来源，用于计算动态内容
	Type        ContentType `bson:"type,omitempty" json:"type"`               // 内容类型
	Font        string      `bson:"font,omitempty" json:"font"`               // 字体
	FontSize    int32       `bson:"fontSize,omitempty" json:"fontSize"`       // 字体大小
	LineSpacing int         `bson:"lineSpacing,omitempty" json:"lineSpacing"` // 文本间距
	Color       string      `bson:"color,omitempty" json:"color"`             // 文本颜色
	Align       AlignType   `bson:"align,omitempty" json:"align"`             // 文本对齐方式
	Bold        bool        `bson:"bold,omitempty" json:"bold"`               // 字体是否加粗
	Position    struct {
		X int32 `bson:"x,omitempty" json:"x"` // x坐标
		Y int32 `bson:"y,omitempty" json:"y"` // y坐标
		W int32 `bson:"w,omitempty" json:"w"` // 宽度
		H int32 `bson:"h,omitempty" json:"h"` // 高度
	}
}

// 水印图片
type WatermarkImage struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Group      string             `bson:"group,omitempty" json:"group"`           // Group
	SubGroup   string             `bson:"subGroup,omitempty" json:"subGroup"`     // subGroup
	ImgTitle   string             `bson:"imgTitle,omitempty" json:"imgTitle"`     // 图片title
	ImgUrl     string             `bson:"imgUrl,omitempty" json:"imgUrl"`         // 图片地址
	Watermarks []*Watermark       `bson:"watermarks,omitempty" json:"watermarks"` // watermark
	Created    *time.Time         `bson:"created,omitempty" json:"created"`       // 创建时间
	Updated    *time.Time         `bson:"updated,omitempty" json:"updated"`       // 更新时间
}
