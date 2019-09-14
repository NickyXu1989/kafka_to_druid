package tools

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	interval time.Duration
	capacity int64
	availability int64
	ticker *time.Ticker
	tokenMutex *sync.Mutex

}


func NewTokenBucket(interval time.Duration, capacity int64) *TokenBucket {
	if interval < 0 {
		panic(fmt.Sprintf("interval < 0"))
	}

	if capacity < 0 {
		panic(fmt.Sprintf("capacity < 0"))
	}

	tb := &TokenBucket{
		interval: interval,
		capacity: capacity,
		availability: capacity,
		ticker: time.NewTicker(interval),
		tokenMutex: &sync.Mutex{},
	}

	go tb.runDaemon()
	return tb

}


func (tb *TokenBucket) runDaemon() {
	for range tb.ticker.C {
		tb.tokenMutex.Lock()
		if tb.availability < tb.capacity {
			tb.availability ++
		}
		tb.tokenMutex.Unlock()
	}
}

func (tb *TokenBucket) Take (need int64) bool{
	tb.tokenMutex.Lock()
	defer tb.tokenMutex.Unlock()
	if need < tb.availability {
		tb.availability = tb.availability - need
		return true
	}
	return false
}

