package ratelimit

import (
	"errors"
	"time"
)

func NewThrottleRateLimiter(conf *Config) (RateLimiter, error) {
	if conf.Throttle == 0 {
		return nil, errors.New("throttle duration muse be greate than zero")
	}

	m := NewManager(conf)

	await := func(throttle time.Duration) {
		ticker := time.NewTicker(throttle)
		go func() {
			for ; true; <-ticker.C {
				<-m.inChan
				m.tryGenerateToken()
			}
		}()
	}

	await(conf.Throttle)

	return m, nil
}
