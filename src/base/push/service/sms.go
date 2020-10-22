package service

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"git.dillonliang.cn/micro-svc/pledge/src/base/push/model"
	"git.dillonliang.cn/micro-svc/pledge/src/base/push/service/sms"
)

func (s *Service) SendSms(group string, mobile int64, smsType int) error {
	smsSigns := s.dao.GetSmsSigns(group, smsType)

	var platformName string
	var sign *model.SmsSign
	if len(*smsSigns) > 1 {
		platformName, sign = s.getSmsPlatform(smsSigns)
	} else if len(*smsSigns) == 1 {
		sign := (*smsSigns)[0]
		platformName = sign.Platform
	} else {
		return errors.New("NOT_SUPPORT")
	}

	if sign == nil {
		return errors.New("NO_SIGN")
	}

	var sm sms.Platform
	sm = getPlatform(sign)

	code := randomCode(sign.CodeNum)
	template := sign.Template
	content := strings.Replace(template, "${code}", code, 1)
	returnCode, err := sm.SendSms(mobile, content)
	if err != nil {
		return err
	}

	data := &model.SmsCode{
		ID:         primitive.NewObjectID(),
		Group:      group,
		Mobile:     mobile,
		Code:       code,
		Content:    content,
		Created:    time.Now(),
		Platform:   platformName,
		ExpireTime: time.Now().Add(30 * time.Minute),
		SmsType:    smsType,
		ReturnCode: returnCode,
	}
	return s.dao.CreateSmsCode(data)
}

func (s *Service) CheckSmsCode(mobile int64, code string, smsType int) (bool, error) {
	record, err := s.dao.GetSmsRecordByMobileAndType(mobile, smsType)
	if err != nil {
		return false, err
	}

	if record == nil {
		return false, errors.New("NO_RECORD")
	}

	if record.Code == code {
		if time.Now().Before(record.ExpireTime) {
			return true, nil
		} else {
			return false, errors.New("CODE_EXPIRE")
		}
	}

	return false, errors.New("BAD_CODE")
}

func (s *Service) getSmsPlatform(smsSigns *[]model.SmsSign) (string, *model.SmsSign) {
	zz253W := s.cfg.SMS.ZZ253.Weight
	momnetsW := s.cfg.SMS.Montnets.Weight
	var weight int

	m := make(map[string]*model.SmsSign)
	for _, k := range *smsSigns {
		sign := k
		m[k.Platform] = &sign
		if k.Platform == "zz253" {
			weight += zz253W
		}
		if k.Platform == "montnets" {
			weight += momnetsW
		}
	}

	platforms := make([]string, 0)
	for i := 0; i < zz253W; i++ {
		platforms = append(platforms, "zz253")
	}

	for i := 0; i < momnetsW; i++ {
		platforms = append(platforms, "montnets")
	}

	// 随机索引
	rand.Seed(time.Now().Unix())
	index := rand.Intn(weight)

	platformName := platforms[index]
	return platformName, m[platformName]
}

func getPlatform(sign *model.SmsSign) sms.Platform {
	var sm sms.Platform
	switch sign.Platform {
	case "zz253":
		sm = &sms.ZZ253{sign.UserName, sign.Password}
	case "montnets":
		sm = &sms.Montnets{sign.UserName, sign.Password}
	default:
		sm = nil
	}

	return sm
}

func randomCode(codeLength int) string {
	format := strings.Replace("%0Nv", "N", strconv.Itoa(codeLength), 1)
	return fmt.Sprintf(format, rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(math.Pow10(codeLength))))
}
