package core

import (
	"log"
	"sync"
	"time"
)

type TaskDispatcher struct {
	accountPool *AccountPool
	taskChan    chan *Task
	wg          sync.WaitGroup
	mu          sync.Mutex
}

func NewTaskDispatcher(pool *AccountPool) *TaskDispatcher {
	return &TaskDispatcher{
		accountPool: pool,
		taskChan:    make(chan *Task, 10000),
	}
}

func (d *TaskDispatcher) AddTask(task *Task) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.taskChan <- task
}

func (d *TaskDispatcher) Run(concurrency int) {
	// 启动工作池
	for i := 0; i < concurrency; i++ {
		d.wg.Add(1)
		go d.worker()
	}

	d.wg.Wait()
	close(d.taskChan)
}

func (d *TaskDispatcher) worker() {
	defer d.wg.Done()

	for task := range d.taskChan {
		acc := d.accountPool.GetAccount()
		log.Printf("使用账号 %s 处理任务: %s %s", acc.ID, task.Method, task.URL)

		// 带重试的执行
		retry := 0
		maxRetries := 3
		for retry < maxRetries {
			if err := ExecuteRequest(task, acc, d); err != nil {
				log.Printf("请求失败 (尝试 %d/%d): %v", retry+1, maxRetries, err)
				retry++
				time.Sleep(time.Duration(retry) * time.Second)
			} else {
				break
			}
		}
	}
}
