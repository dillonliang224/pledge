package sms

type Platform interface {
	SendSms(mobile int64, content string) (string, error)
}
