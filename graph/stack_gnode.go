package graph

import (
	"sync"
)

type GNodeStack struct {
	items []GNode
	lock  sync.RWMutex
}

// 创建栈
func (s *GNodeStack) New() *GNodeStack {
	s.items = []GNode{}
	return s
}

// 入栈
func (s *GNodeStack) Push(t GNode) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// 出栈
func (s *GNodeStack) Pop() *GNode {
	s.lock.Lock()
	if s.IsEmpty() {
		return nil
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.lock.Unlock()
	return &item
}

//IsEmpty 是否为空栈
func (s *GNodeStack) IsEmpty() bool {
	return len(s.items) == 0
}
