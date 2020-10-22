package sms

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/axgle/mahonia"
)

type Montnets struct {
	UserName string
	Password string
}

type MontnetsRes struct {
	Result int   `json:"result"`
	MsgId  int64 `json:"msgid"`
}

func (p *Montnets) SendSms(mobile int64, content string) (string, error) {
	ts := time.Now().Format("0102150405")
	pwd := p.cryptPwd(ts)

	data := map[string]interface{}{
		"userid":    p.UserName,
		"pwd":       pwd,
		"mobile":    strconv.FormatInt(mobile, 10),
		"content":   p.formatContent(content),
		"timestamp": ts,
	}

	b, _ := json.Marshal(data)
	resBody, err := p.doRequest("http://61.145.229.26:8086/sms/v2/std/single_send", b)
	if err != nil {
		return "", err
	}

	var res MontnetsRes
	err = json.Unmarshal(resBody, &res)
	if err != nil {
		return "", errors.New("Montnets unmarshall json err")
	}

	return strconv.Itoa(int(res.MsgId)), nil
}

func (p *Montnets) cryptPwd(ts string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(p.UserName + "00000000" + p.Password + ts))
	encryptPwd := md5Ctx.Sum(nil)
	return hex.EncodeToString(encryptPwd[:])
}

func (p *Montnets) formatContent(content string) string {
	content = strings.TrimSpace(content)
	gbk := mahonia.NewEncoder("gbk").ConvertString(content)
	v := url.Values{}
	v.Set("aa", gbk)
	str := v.Encode()
	arr := strings.Split(str, "=")
	return arr[1]
}

func (p *Montnets) doRequest(url string, data []byte) ([]byte, error) {
	body := strings.NewReader(string(data[:]))
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Connection", "Close")
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
