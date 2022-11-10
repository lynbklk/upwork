package worker

import (
	"testcookie/pkg/headless"
)

// Worker defines methods that standard worker should have
type Worker interface {
	Start()
	Stop()
}

func New(headlessBrowser *headless.HeadlessBrowser) *Manager {
	return &Manager{
		HeadlessBrowser: headlessBrowser,
	}
}

// Manager manages and passes shared properties to all workers
type Manager struct {
	HeadlessBrowser *headless.HeadlessBrowser
	workers         []Worker
}

func (manager *Manager) add(workers ...Worker) {
	manager.workers = append(manager.workers, workers...)
}

func (manager *Manager) register() {
	loginWorker := &loginWorker{
		Name:            "AHX:LOGIN",
		HeadlessBrowser: manager.HeadlessBrowser,
		StopCh:          make(chan bool),
	}

	manager.add(loginWorker)
}

// StartAll starts all worker in different goroutines
func (manager *Manager) StartAll() {
	manager.register()
	for _, worker := range manager.workers {
		go worker.Start()
	}
}

// StopAll stops all worker in different goroutines
func (manager *Manager) StopAll() {
	for _, worker := range manager.workers {
		worker.Stop()
	}
}
