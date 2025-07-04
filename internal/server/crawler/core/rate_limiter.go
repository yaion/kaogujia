package core

import (
	"sync"
	"time"
)

// 令牌桶速率限制器
type RateLimiter struct {
	limit      int           // 每分钟请求数
	tokens     int           // 当前令牌数
	maxTokens  int           // 最大令牌数
	lastRefill time.Time     // 上次补充令牌时间
	refillRate time.Duration // 补充速率
	mu         sync.Mutex
}

func NewRateLimiter(limitPerMinute int) *RateLimiter {
	refillRate := time.Minute / time.Duration(limitPerMinute)
	return &RateLimiter{
		limit:      limitPerMinute,
		tokens:     limitPerMinute,
		maxTokens:  limitPerMinute,
		lastRefill: time.Now(),
		refillRate: refillRate,
	}
}

func (rl *RateLimiter) Limit() int {
	return rl.limit
}

func (rl *RateLimiter) Wait() {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	// 补充令牌
	now := time.Now()
	elapsed := now.Sub(rl.lastRefill)
	tokensToAdd := int(elapsed / rl.refillRate)

	if tokensToAdd > 0 {
		rl.tokens += tokensToAdd
		if rl.tokens > rl.maxTokens {
			rl.tokens = rl.maxTokens
		}
		rl.lastRefill = now
	}

	// 如果令牌不足，等待
	if rl.tokens <= 0 {
		waitTime := rl.refillRate - elapsed
		if waitTime > 0 {
			time.Sleep(waitTime)
		}
		rl.tokens = 0
		rl.lastRefill = time.Now()
	}

	// 消耗一个令牌
	rl.tokens--
}
