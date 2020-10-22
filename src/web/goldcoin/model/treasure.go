package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ActivityState int

const (
	CLOSE ActivityState = iota
	OPEN
)

// 天天夺宝活动
type TreasureActivity struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`               // 活动ID
	State      ActivityState      `bson:"state,omitempty" json:"state"`           // 活动状态
	Title      string             `bson:"title,omitempty" json:"title"`           // 活动名
	StartTime  *time.Time         `bson:"startTime,omitempty" json:"startTime"`   // 开始时间
	EndTime    *time.Time         `bson:"endTime,omitempty" json:"endTime"`       // 结束时间
	BuyTime    string             `bson:"buyTime,omitempty" json:"buyTime"`       // 每日购买时间
	BuyEndTime string             `bson:"buyEndTime,omitempty" json:"buyEndTime"` // 每日截止购买时间
	DrawTime   string             `bson:"drawTime,omitempty" json:"drawTime"`     // 每日开奖时间
	Products   []string           `bson:"products,omitempty" json:"products"`     // 商品ID列表
	Multiple   float32            `bson:"multiple,omitempty" json:"multiple"`     // 回收倍数，用于开奖
	Created    *time.Time         `bson:"created,omitempty" json:"created"`       // 创建时间
	Updated    *time.Time         `bson:"updated,omitempty" json:"updated"`       // 更新时间
}

// 天天夺宝商品
type TreasureProduct struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`                     // 商品ID
	Title         string             `bson:"title,omitempty" json:"title"`                 // 商品名
	Order         int32              `bson:"order,omitempty" json:"order"`                 // 商品顺序
	Image         string             `bson:"image,omitempty" json:"image"`                 // 商品主图
	Cost          int32              `bson:"cost,omitempty" json:"cost"`                   // 商品成本（金币）
	OfficialPrice int32              `bson:"officialPrice,omitempty" json:"officialPrice"` // 官方售价（元）
	Price         int32              `bson:"price,omitempty" json:"price"`                 // 夺宝商品单价（金币）
	ProdDesc      string             `bson:"prodDesc,omitempty" json:"prodDesc"`           // 商品描述
	ProdVideo     string             `bson:"prodVideo,omitempty" json:"prodVideo"`         // 商品视频
	MaxLimit      int32              `bson:"maxLimit,omitempty" json:"maxLimit"`           // 最大夺宝码数量
	Updated       *time.Time         `bson:"updated,omitempty" json:"updated"`             // 更新时间
}

type RecordWinState int

const (
	InProgress RecordWinState = iota
	Win
	UnWin = -1
)

// 天天夺宝记录
type TreasureRecord struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"_id"`               // 记录ID
	ActivityId primitive.ObjectID `bson:"activityId,omitempty" json:"activityId"` // 活动ID
	ProductId  primitive.ObjectID `bson:"productId,omitempty" json:"productId"`   // 商品ID
	UserId     string             `bson:"userId,omitempty" json:"userId"`         // 用户ID
	Codes      []string           `bson:"codes,omitempty" json:"codes"`           // 夺宝码
	Gold       int32              `bson:"gold,omitempty" json:"gold"`             // 花费金币数
	Period     string             `bson:"period,omitempty" json:"period"`         // 期数
	WinState   RecordWinState     `bson:"winState,omitempty" json:"winState"`     // 中奖状态
	OpenTime   *time.Time         `bson:"openTime,omitempty" json:"openTime"`     // 开奖时间
	Mock       bool               `bson:"mock,omitempty" json:"mock"`             // mock
	Created    *time.Time         `bson:"created,omitempty" json:"created"`       // 创建时间
}

type WinnerRewardState int

const (
	UnReward WinnerRewardState = 0
	Reward   WinnerRewardState = 0
)

// 天天夺宝中奖名单(冗余record部分字段，好做往期揭晓)
type TreasureWinner struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id"`                 // _
	ActivityId  primitive.ObjectID `bson:"activityId,omitempty" json:"activityId"`   // 活动ID
	ProductId   primitive.ObjectID `bson:"productId,omitempty" json:"productId"`     // 商品ID
	UserId      string             `bson:"userId,omitempty" json:"userId"`           // 用户ID
	Nickname    string             `bson:"nickname,omitempty" json:"nickname"`       // 头像
	Avatar      string             `bson:"avatar,omitempty" json:"avatar"`           // 昵称
	Code        string             `bson:"code,omitempty" json:"code"`               // 中奖夺宝码
	Period      string             `bson:"period,omitempty" json:"period"`           // 期数
	RewardState WinnerRewardState  `bson:"rewardState,omitempty" json:"rewardState"` // 奖品发放状态
	OpenTime    *time.Time         `bson:"openTime,omitempty" json:"openTime"`       // 开奖时间
	Mock        bool               `bson:"mock,omitempty" json:"mock"`               // mock
	Created     *time.Time         `bson:"created,omitempty" json:"created"`         // 创建时间
}

// 天天夺宝期数
type TreasurePeriod struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"_id"`               // _
	ActivityId    primitive.ObjectID `bson:"activityId,omitempty" json:"activityId"` // 活动ID
	ProductId     primitive.ObjectID `bson:"productId,omitempty" json:"productId"`   // 商品ID
	Period        string             `bson:"period,omitempty" json:"period"`         // 活动期号
	OpenTime      *time.Time         `bson:"openTime,omitempty" json:"openTime"`     // 开奖时间
	IsOpen        bool               `bson:"isOpen,omitempty" json:"isOpen"`         // 是否开奖
	RealPeoPleNum int32              `bson:"peopleReal,omitempty" json:"peopleReal"` // 真实用户数
	MockPeopleNum int32              `bson:"peopleMock,omitempty" json:"peopleMock"` // mock用户数
}

// 天天夺宝免费机会
type TreasureFreeChance struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id"`       // _
	UserId string             `bson:"userId,omitempty" json:"userId"` // 用户ID
	Date   int32              `bson:"date,omitempty" json:"date"`     // 日期
	Nums   int32              `bson:"nums,omitempty" json:"nums"`     // 已获得免费次数
	Used   int32              `bson:"used,omitempty" json:"used"`     // 已用次数
}

// 天天夺宝机器人
type TreasureRobot struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"_id"`       // _
	UserId string             `bson:"userId,omitempty" json:"userId"` // 用户ID
}

const (
	TableTreasureActivity   = "TreasureActivity"
	TableTreasureProduct    = "TreasureProduct"
	TableTreasureRecord     = "TreasureRecord"
	TableTreasureWinner     = "TreasureWinner"
	TableTreasurePeriod     = "TreasurePeriod"
	TableTreasureFreeChance = "TreasureFreeChance"
	TableTreasureRobot      = "TreasureRobot"
)
