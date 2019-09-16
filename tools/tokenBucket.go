package tools

import (
	"fmt"
	"sync"
	"time"
)

type TokenBucket struct {
	interval time.Duration
	capacity int
	availability int
	limitPerSec int
	ticker *time.Ticker
	tokenMutex *sync.Mutex

}


func NewTokenBucket(limit int, capacity int) *TokenBucket {
	if limit < 0 {
		panic(fmt.Sprintf("interval < 0"))
	}

	if capacity < 0 {
		panic(fmt.Sprintf("capacity < 0"))
	}

	tb := &TokenBucket{
		interval: time.Second,
		capacity: capacity,
		availability: capacity,
		limitPerSec: limit,
		ticker: time.NewTicker(time.Second),
		tokenMutex: &sync.Mutex{},
	}

	go tb.runDaemon()
	return tb

}


func (tb *TokenBucket) runDaemon() {
	for range tb.ticker.C {
		tb.tokenMutex.Lock()
		if tb.availability < tb.capacity {
			tb.availability = tb.availability + tb.limitPerSec
			if tb.availability > tb.capacity{
				tb.availability = tb.capacity
			}
		}
		tb.tokenMutex.Unlock()
	}
}

func (tb *TokenBucket) Take (need int) bool{
	tb.tokenMutex.Lock()
	defer tb.tokenMutex.Unlock()
	if need < tb.availability {
		tb.availability = tb.availability - need
		return true
	}
	return false
}

