package scheduler

import "go-spider/engine"

// 简单调度器
type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

func (s *SimpleScheduler) Submit(
	r engine.Request) {
	go func() { s.workerChan <- r }()
}
