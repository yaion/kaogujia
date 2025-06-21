package crawler

import (
	"sort"
	"sync"
	"time"
)

type Account struct {
	Username  string
	Password  string
	Cookies   map[string]string
	Proxy     string        // 专属代理
	RateLimit time.Duration // 请求间隔
	LastUsed  time.Time     // 最后使用时间
}

type AccountManager struct {
	accounts []*Account
	lock     sync.Mutex
}

func (m *AccountManager) GetAccount() *Account {
	m.lock.Lock()
	defer m.lock.Unlock()

	// 选择最近未使用的账号
	sort.Slice(m.accounts, func(i, j int) bool {
		return m.accounts[i].LastUsed.Before(m.accounts[j].LastUsed)
	})

	acc := m.accounts[0]
	acc.LastUsed = time.Now()
	return acc
}
