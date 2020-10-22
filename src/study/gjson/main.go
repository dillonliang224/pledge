package main

import (
	"encoding/json"
	"os"
	"time"
)

type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
}

func (u *MyUser) MarshalJSON() ([]byte, error) {
	type Alias MyUser
	return json.Marshal(&struct {
		LastSeen int64 `json:"lastSeen"`
		*Alias
	}{
		Alias:    (*Alias)(u),
		LastSeen: u.LastSeen.Unix(),
	})
}

func (u *MyUser) UnmarshalJSON(data []byte) error {
	type Alias MyUser

	aux := &struct {
		LastSeen int64 `json:"lastSeen"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	u.LastSeen = time.Unix(aux.LastSeen, 0)
	return nil
}

// https://colobu.com/2020/03/19/Custom-JSON-Marshalling-in-Go/
// 自定义json序列化方法
func main() {
	_ = json.NewEncoder(os.Stdout).Encode(&MyUser{
		1, "dillon", time.Now(),
	})
}
