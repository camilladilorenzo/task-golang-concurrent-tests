package main

import (
	"sync"
)

type MeasuredWorker struct {
	Worker
	mu    sync.Mutex
	value int
}

func (m *MeasuredWorker) Work() {
	m.Worker.Work()
	m.mu.Lock()
	m.value++
	m.mu.Unlock()
}

func (m *MeasuredWorker) Value() int {
	return m.value
}
