package crawler

import (
	"kaogujia/pkg/config"
	"sort"
	"sync"
	"time"
)

type Account struct {
	Username      string
	Password      string
	Authorization string
	Proxy         string        // 专属代理
	RateLimit     time.Duration // 请求间隔
	LastUsed      time.Time     // 最后使用时间
}

type AccountPool struct {
	accounts []*Account
	lock     sync.Mutex
}

type MultiAccountManager struct {
	pools map[string]*AccountPool // key: website name
}

func NewMultiAccountManager(cfg *config.AppConfig) *MultiAccountManager {
	manager := &MultiAccountManager{
		pools: make(map[string]*AccountPool),
	}

	for _, website := range cfg.Websites {
		var accounts []*Account
		for _, accCfg := range website.Accounts {
			accounts = append(accounts, &Account{
				Username:      accCfg.Username,
				Password:      accCfg.Password,
				Authorization: accCfg.Authorization,
				//Cookies:   accCfg.Cookies,
				//Proxy:     accCfg.Proxy,
				//RateLimit: accCfg.RateLimit,
			})
		}
		manager.pools[website.Name] = &AccountPool{
			accounts: accounts,
		}
	}
	return manager
}

func (m *MultiAccountManager) GetAccount(website string) *Account {
	pool, ok := m.pools[website]
	if !ok {
		return nil
	}

	pool.lock.Lock()
	defer pool.lock.Unlock()

	// 选择最近未使用的账号
	sort.Slice(pool.accounts, func(i, j int) bool {
		return pool.accounts[i].LastUsed.Before(pool.accounts[j].LastUsed)
	})

	acc := pool.accounts[0]
	acc.LastUsed = time.Now()
	return acc
}

func (m *MultiAccountManager) SetAccount(website string, acc *Account) {
	pool, ok := m.pools[website]
	if !ok {
		return
	}

	pool.lock.Lock()
	defer pool.lock.Unlock()

	// 账号使用完毕，将其放回账号池
	acc.LastUsed = time.Now()
}
