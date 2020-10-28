package kafka

import (
	"testing"
	"time"
)

func TestConsume(t *testing.T) {
	Consume()
	time.Sleep(5 * time.Second)
}
