package scheduler

import "demoCrawler/engine"

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

//func (s *SimpleScheduler) ConfigureMasterWorkerChan(c chan engine.Request) {
//	s.workerChan = c
//}

//
func (s *SimpleScheduler) Submit(r engine.Request) {
	// ***注意***
	go func() {
		s.workerChan <- r
	}()
}
