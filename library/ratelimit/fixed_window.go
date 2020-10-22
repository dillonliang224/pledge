package ratelimit

import (
	"errors"
	"time"
)

type FixedWindowInterval struct {
	startTime time.Time
	endTime   time.Time
	interval  time.Duration
}

func (w *FixedWindowInterval) setWindowTime() {
	w.startTime = time.Now().UTC()
	w.endTime = time.Now().UTC().Add(w.interval)
}

func (w *FixedWindowInterval) run(cb func()) {
	go func() {
		ticker := time.NewTicker(w.interval)
		w.setWindowTime()
		for range ticker.C {
			cb()
			w.setWindowTime()
		}
	}()
}

func NewFixedWindowRateLimit(conf *Config) (RateLimiter, error) {
	if conf.FixedInterval == 0 {
		return nil, errors.New("invalid interval")
	}

	if conf.Limit == 0 {
		return nil, errors.New("invalid limit")
	}

	m := NewManager(conf)
	w := &FixedWindowInterval{interval: conf.FixedInterval}

	m.makeToken = func() *Token {
		t := NewToken()
		t.ExpiredAt = w.endTime
		return t
	}

	await := func() {
		go func() {
			for {
				select {
				case <-m.inChan:
					m.tryGenerateToken()
				case token := <-m.releaseChan:
					m.releaseToken(token)
				}
			}
		}()
	}

	w.run(m.releaseExpiredTokens)
	await()
	return m, nil
}
