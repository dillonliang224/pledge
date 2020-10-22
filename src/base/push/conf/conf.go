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

	SMS struct {
		ZZ253 struct {
			Weight int
		}

		Montnets struct {
			Weight int
		}
	}
}

func Load(app string) (conf *Config) {
	conf = new(Config)
	if err := config.Load(app, conf); err != nil {
		panic(fmt.Sprintf("config load failed: %v", err))
	}
	return
}
