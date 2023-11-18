package autosolve

import "sync"

type task struct {
	TaskId    string `json:"taskId"`
	CreatedAt int64  `json:"createdAt"`
	Status    string `json:"status"`
	Token     string `json:"token"`
}

type taskMap struct {
	mu   sync.RWMutex
	data map[string]task
}

func newTaskMap() *taskMap {
	return &taskMap{
		data: make(map[string]task),
	}
}

func (s *taskMap) Set(k string, v task) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[k] = v
}

func (s *taskMap) Get(k string) (task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[k]
	return val, ok
}

func (s *taskMap) Delete(k string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, k)
}

func (s *taskMap) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.data)
}

func (s *taskMap) ForEach(f func(string, task)) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for key, val := range s.data {
		f(key, val)
	}
}
