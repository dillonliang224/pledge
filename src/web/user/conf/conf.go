package conf

import (
	"fmt"

	"git.dillonliang.cn/micro-svc/pledge/library/config"
)

type Config struct {
	config.Common
	Log struct {
		Level string
	}
}

func Load(app string) (conf *Config) {
	conf = new(Config)
	if err := config.Load(app, conf); err != nil {
		panic(fmt.Sprintf("config load failed: %v", err))
	}
	return
}
