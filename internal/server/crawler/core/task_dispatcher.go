package core

import (
	"sync"
)

type TaskDispatcher struct {
	tasks       []*Task
	accountPool *AccountPool
	wg          sync.WaitGroup
}

func NewTaskDispatcher(pool *AccountPool) *TaskDispatcher {
	return &TaskDispatcher{
		accountPool: pool,
	}
}

func (d *TaskDispatcher) AddTask(task *Task) {
	d.tasks = append(d.tasks, task)
}

func (d *TaskDispatcher) Run(concurrency int) error {
	taskChan := make(chan *Task, len(d.tasks))

	// 填充任务通道
	for _, task := range d.tasks {
		taskChan <- task
	}
	close(taskChan)

	// 启动工作池
	for i := 0; i < concurrency; i++ {
		d.wg.Add(1)
		go d.worker(taskChan)
	}

	d.wg.Wait()
	return nil
}

func (d *TaskDispatcher) worker(taskChan <-chan *Task) {
	defer d.wg.Done()

	for task := range taskChan {
		acc := d.accountPool.GetAccount()
		if err := ExecuteRequest(task, acc); err != nil {
			// 错误处理
		}
	}
}
