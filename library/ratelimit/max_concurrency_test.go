package ratelimit

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewMaxCurrencyRateLimit(t *testing.T) {
	r, err := NewMaxCurrencyRateLimit(&Config{
		Limit:            5,
		TokenResetsAfter: 10 * time.Second,
	})

	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	doWork := func(id int) {
		token, err := r.Acquire()
		fmt.Printf("ratelimit token %s acquired at %s...\n", token.ID, token.CreatedAt)
		if err != nil {
			panic(err)
		}

		n := rand.Intn(5)
		fmt.Printf("Worker %d Sleeping %d seconds...\n", id, n)
		time.Sleep(time.Duration(n) * time.Second)
		fmt.Printf("Worker %d Done\n", id)
		r.Release(token)
		wg.Done()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doWork(i)
	}

	wg.Wait()

	assert.True(t, true, "true is true1111111111111111111111111")
}
