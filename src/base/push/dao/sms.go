package dao

import (
	"go.mongodb.org/mongo-driver/bson"

	"git.dillonliang.cn/micro-svc/pledge/library/mongodb"
	"git.dillonliang.cn/micro-svc/pledge/src/base/push/model"
)

func (d *Dao) CreateSmsCode(data interface{}) error {
	_, err := d.smsCode.Insert(data)
	return err
}

func (d *Dao) GetSmsRecordByMobileAndType(mobile int64, smsType int) (*model.SmsCode, error) {
	var records []model.SmsCode
	err := d.smsCode.Find(&records, bson.M{
		"mobile": mobile,
		"type":   smsType,
	}, mongodb.Sort(bson.M{"_id": -1}), mongodb.Limit(1))

	if err != nil {
		return nil, err
	}

	if len(records) > 0 {
		return &records[0], nil
	}
	return nil, nil
}

func (d *Dao) GetSmsSigns(group string, smsType int) *[]model.SmsSign {
	var signs []model.SmsSign
	filter := bson.M{
		"group": group,
		"type":  smsType,
	}
	_ = d.smsSign.Find(&signs, filter)
	return &signs
}
