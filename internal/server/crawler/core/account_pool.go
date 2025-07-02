package core

import (
	"sync"
	"time"
)

type Account struct {
	ID       string
	Proxy    string
	Cookies  []interface{}
	LastUsed time.Time
	mu       sync.Mutex
}

type AccountPool struct {
	accounts   []*Account
	interval   time.Duration
	currentIdx int
	mu         sync.Mutex
}

func NewAccountPool(accounts []*Account, interval time.Duration) *AccountPool {
	return &AccountPool{
		accounts: accounts,
		interval: interval,
	}
}

func (p *AccountPool) GetAccount() *Account {
	p.mu.Lock()
	defer p.mu.Unlock()

	for i := 0; i < len(p.accounts); i++ {
		p.currentIdx = (p.currentIdx + 1) % len(p.accounts)
		acc := p.accounts[p.currentIdx]

		acc.mu.Lock()
		defer acc.mu.Unlock()

		if time.Since(acc.LastUsed) >= p.interval {
			acc.LastUsed = time.Now()
			return acc
		}
	}

	// 所有账号都在冷却中，等待最快要可用的
	next := time.Until(p.accounts[p.currentIdx].LastUsed.Add(p.interval))
	time.Sleep(next)
	return p.GetAccount()
}
