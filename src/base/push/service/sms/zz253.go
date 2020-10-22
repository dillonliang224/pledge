package sms

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type ZZ253 struct {
	UserName string
	Password string
}

func (p *ZZ253) SendSms(mobile int64, content string) (string, error) {
	data := map[string]interface{}{
		"account":  p.UserName,
		"password": p.Password,
		"phone":    mobile,
		"msg":      content,
		"report":   false,
	}

	b, _ := json.Marshal(data)
	_, err := p.doRequest("http://smssh1.253.com/msg/send/json", b)
	if err != nil {
		return "", err
	}

	return "0", nil
}

func (p *ZZ253) doRequest(url string, data []byte) ([]byte, error) {
	body := strings.NewReader(string(data[:]))
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	res, err := http.DefaultClient.Do(req)
	if res != nil {
		defer res.Body.Close()
	}
	if err != nil {
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}
