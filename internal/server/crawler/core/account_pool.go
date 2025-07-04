package core

import (
	"log"
	"sync"
	"time"
)

// 账号配置
type AccountConfig struct {
	ID        string `yaml:"id"`
	Proxy     string `yaml:"proxy"`
	RateLimit int    `yaml:"rate_limit"` // 每分钟请求数
}

// 账号结构
type Account struct {
	ID        string
	Name      string
	Password  string
	Proxy     string
	Cookies   []interface{}
	LastUsed  time.Time
	RateLimit *RateLimiter
	mu        sync.Mutex
}

type AccountPool struct {
	accounts   []*Account
	currentIdx int
	mu         sync.Mutex
}

func NewAccountPool(configs []AccountConfig) *AccountPool {
	pool := &AccountPool{}
	for _, config := range configs {
		account := &Account{
			ID:       config.ID,
			Proxy:    config.Proxy,
			Cookies:  nil,
			LastUsed: time.Now().Add(-5 * time.Minute),
		}

		// 初始化速率限制器
		if config.RateLimit > 0 {
			account.RateLimit = NewRateLimiter(config.RateLimit)
			log.Printf("账号 %s 速率限制: %d 请求/分钟", config.ID, config.RateLimit)
		}

		pool.accounts = append(pool.accounts, account)
	}
	return pool
}

func (p *AccountPool) GetAccount() *Account {
	p.mu.Lock()
	defer p.mu.Unlock()

	startIdx := p.currentIdx
	for {
		p.currentIdx = (p.currentIdx + 1) % len(p.accounts)
		acc := p.accounts[p.currentIdx]

		acc.mu.Lock()
		elapsed := time.Since(acc.LastUsed)

		// 计算账号冷却时间（至少1秒）
		cooldown := time.Second
		if acc.RateLimit != nil {
			cooldown = time.Minute / time.Duration(acc.RateLimit.Limit())
		}

		if elapsed >= cooldown {
			acc.LastUsed = time.Now()
			acc.mu.Unlock()
			return acc
		}
		acc.mu.Unlock()

		// 如果转了一圈都没找到可用账号，等待最快要可用的账号
		if p.currentIdx == startIdx {
			var minTime time.Duration = time.Minute
			for _, acc := range p.accounts {
				acc.mu.Lock()
				cooldown := time.Second
				if acc.RateLimit != nil {
					cooldown = time.Minute / time.Duration(acc.RateLimit.Limit())
				}

				remaining := cooldown - time.Since(acc.LastUsed)
				if remaining < minTime {
					minTime = remaining
				}
				acc.mu.Unlock()
			}

			if minTime > 0 {
				time.Sleep(minTime)
			}
			break
		}
	}

	// 递归调用，此时应该可以获取到账号
	return p.GetAccount()
}
