package graph

import (
	"sync"
	"fmt"
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

func GetPath(path map[GNode]*GNode, source *GNode, target *GNode) string {
	var stack GNodeStack
	stack.New()
	stack.Push(*target)
	for pathPrev := path[*target]; pathPrev != nil; pathPrev = path[*pathPrev] {
		if pathPrev == nil {
			break
		}
		stack.Push(*pathPrev)
	}
	result := ""
	for node := stack.Pop(); node != nil; node = stack.Pop() {
		result += fmt.Sprintf("%d ", node.value)
		if stack.IsEmpty() {
			break
		}
	}
	return result
}
