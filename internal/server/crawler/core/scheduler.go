package core

import (
	"github.com/robfig/cron/v3"
	"log"
)

type Scheduler struct {
	cron       *cron.Cron
	dispatcher *TaskDispatcher
	handlerMap map[string]ResponseHandler
}

func NewScheduler(dispatcher *TaskDispatcher, handlerMap map[string]ResponseHandler) *Scheduler {
	return &Scheduler{
		cron:       cron.New(),
		dispatcher: dispatcher,
		handlerMap: handlerMap,
	}
}

func (s *Scheduler) AddSchedule(name, cronExpr string, taskSpecs []TaskSpec) {
	_, err := s.cron.AddFunc(cronExpr, func() {
		log.Printf("执行定时任务: %s", name)
		for _, spec := range taskSpecs {
			handler, ok := s.handlerMap[spec.Handler]
			if !ok {
				log.Printf("未知的处理器: %s", spec.Handler)
				continue
			}

			task := &Task{
				URL:     spec.URL,
				Method:  spec.Method,
				Headers: spec.Headers,
				Body:    spec.Body,
				Handler: handler,
				Meta:    spec.Meta,
			}

			s.dispatcher.AddTask(task)
			log.Printf("添加任务: %s %s (处理器: %s)", spec.Method, spec.URL, spec.Handler)
		}
	})

	if err != nil {
		log.Printf("添加定时任务失败: %s, 错误: %v", name, err)
	} else {
		log.Printf("已添加定时任务: %s, Cron: %s", name, cronExpr)
	}
}

func (s *Scheduler) Start() {
	s.cron.Start()
	log.Println("定时任务调度器已启动")

	// 输出下次执行时间
	for _, entry := range s.cron.Entries() {
		log.Printf("任务 %d 下次执行时间: %v", entry.ID, entry.Next)
	}
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
	log.Println("定时任务调度器已停止")
}
