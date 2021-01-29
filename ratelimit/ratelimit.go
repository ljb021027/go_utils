package main

import (
	"sync"
	"time"
)

type RateLimit interface {
	Wait(count int64)
	Take(count int64) time.Duration
}

type RedisRateLimit struct {

	// startTime holds the moment when the bucket was
	// first created and ticks began.
	startTime time.Time

	// capacity holds the overall capacity of the bucket.
	capacity int64

	// quantum holds how many tokens are added on
	// each tick.
	quantum int64

	// fillInterval holds the interval between each tick.
	fillInterval time.Duration

	// mu guards the fields below it.
	mu sync.Mutex

	// availableTokens holds the number of available
	// tokens as of the associated latestTick.
	// It will be negative when there are consumers
	// waiting for tokens.
	availableTokens int64

	// latestTick holds the latest tick for which
	// we know the number of tokens in the bucket.
	latestTick int64
}

func (tb *RedisRateLimit) Wait(count int64) {
	if d := tb.Take(count); d > 0 {
		//tb.Sleep(d)
	}
}

func (tb *RedisRateLimit) Take(count int64) time.Duration {

}

func (tb *RedisRateLimit) adjustavailableTokens(tick int64) {
	lastTick := tb.latestTick
	tb.latestTick = tick
	if tb.availableTokens >= tb.capacity {
		return
	}
	tb.availableTokens += (tick - lastTick) * tb.quantum
	if tb.availableTokens > tb.capacity {
		tb.availableTokens = tb.capacity
	}
	return
}

func (tb *RedisRateLimit) take(now time.Time, count int64) (time.Duration, bool) {
	if count <= 0 {
		return 0, true
	}

	tick := tb.currentTick(now)
	tb.adjustavailableTokens(tick)
	avail := tb.availableTokens - count
	if avail >= 0 {
		tb.availableTokens = avail
		return 0, true
	}
	// Round up the missing tokens to the nearest multiple
	// of quantum - the tokens won't be available until
	// that tick.

	// endTick holds the tick when all the requested tokens will
	// become available.
	endTick := tick + (-avail+tb.quantum-1)/tb.quantum
	endTime := tb.startTime.Add(time.Duration(endTick) * tb.fillInterval)
	waitTime := endTime.Sub(now)
	//if waitTime > maxWait {
	//	return 0, false
	//}
	tb.availableTokens = avail
	return waitTime, true

}

func (tb *RedisRateLimit) currentTick(now time.Time) int64 {
	return int64(now.Sub(tb.startTime) / tb.fillInterval)
}
