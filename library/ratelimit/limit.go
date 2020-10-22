package ratelimit

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/segmentio/ksuid"
)

// https://github.com/jpg013/go_rate_limiter

type Token struct {
	ID        string
	CreatedAt time.Time
	ExpiredAt time.Time
}

type tokenFactory func() *Token

func NewToken() *Token {
	return &Token{
		ID:        ksuid.New().String(),
		CreatedAt: time.Now().UTC(),
		ExpiredAt: time.Time{},
	}
}

func (t *Token) NeedReset(resetAfter time.Duration) bool {
	if time.Since(t.CreatedAt) >= resetAfter {
		return true
	}

	return false
}

func (t *Token) IsExpired() bool {
	now := time.Now().UTC()
	return t.ExpiredAt.Before(now)
}

type RateLimiter interface {
	Acquire() (*Token, error)
	Release(*Token)
}

type Config struct {
	Throttle time.Duration

	Limit            int
	TokenResetsAfter time.Duration

	FixedInterval time.Duration
}

type Manager struct {
	errorChan    chan error
	releaseChan  chan *Token
	outChan      chan *Token
	inChan       chan struct{}
	needToken    int64
	activeTokens map[string]*Token
	limit        int
	makeToken    tokenFactory
}

func NewManager(conf *Config) *Manager {
	m := &Manager{
		errorChan:    make(chan error),
		outChan:      make(chan *Token),
		inChan:       make(chan struct{}),
		activeTokens: make(map[string]*Token),
		releaseChan:  make(chan *Token),
		needToken:    0,
		limit:        conf.Limit,
		makeToken:    NewToken,
	}

	if conf.TokenResetsAfter > 0 {
		m.runResetTokenTask(conf.TokenResetsAfter)
	}

	return m
}

func (m *Manager) incNeedToken() {
	atomic.AddInt64(&m.needToken, 1)
}

func (m *Manager) decNeedToken() {
	atomic.AddInt64(&m.needToken, -1)
}

func (m *Manager) awaitingToken() bool {
	return atomic.LoadInt64(&m.needToken) > 0
}

func (m *Manager) isLimitExceeded() bool {
	if len(m.activeTokens) >= m.limit {
		return true
	}

	return false
}

func (m *Manager) Acquire() (*Token, error) {
	go func() {
		m.inChan <- struct{}{}
	}()

	select {
	case t := <-m.outChan:
		return t, nil
	case err := <-m.errorChan:
		return nil, err
	}
}

func (m *Manager) Release(token *Token) {
	if token.IsExpired() {
		go func() {
			m.releaseChan <- token
		}()
	}
}

func (m *Manager) releaseToken(token *Token) {
	if token == nil {
		log.Print("unable to release nil token")
		return
	}

	if _, ok := m.activeTokens[token.ID]; !ok {
		log.Printf("unable to release token %s - not in use", token)
		return
	}

	delete(m.activeTokens, token.ID)

	if m.awaitingToken() {
		m.decNeedToken()
		go m.tryGenerateToken()
	}
}

func (m *Manager) tryGenerateToken() {
	if m.makeToken == nil {
		panic("tokenMethod")
	}

	if m.isLimitExceeded() {
		m.incNeedToken()
		return
	}

	token := m.makeToken()
	m.activeTokens[token.ID] = token

	go func() {
		m.outChan <- token
	}()
}

func (m *Manager) runResetTokenTask(resetAfter time.Duration) {
	go func() {
		ticker := time.NewTicker(resetAfter)
		for range ticker.C {
			for _, token := range m.activeTokens {
				if token.NeedReset(resetAfter) {
					go func(t *Token) {
						m.releaseChan <- t
					}(token)
				}
			}
		}
	}()
}

func (m *Manager) releaseExpiredTokens() {
	for _, token := range m.activeTokens {
		for token.IsExpired() {
			go func(t *Token) {
				m.releaseChan <- t
			}(token)
		}
	}
}
