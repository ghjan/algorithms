package stack

import (
	"sync"

	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type ItemStack struct {
	items []Item
	lock  sync.RWMutex
}

// 创建栈
func (s *ItemStack) New() *ItemStack {
	s.items = []Item{}
	return s
}

// 入栈
func (s *ItemStack) Push(t Item) {
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

// 出栈
func (s *ItemStack) Pop() *Item {
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	s.lock.Unlock()
	return &item
}

//IsEmpty 是否为空栈
func (s *ItemStack) IsEmpty() bool {
	return len(s.items) == 0
}

//Size 返回堆栈大小
func (s *ItemStack) Size() int {
	s.lock.Lock()
	defer s.lock.Unlock()
	return len(s.items)
}

//Peek 返回最后一个，区别于Pop函数，Peek仅仅偷看一下，元素继续保留在堆栈中
func (s *ItemStack) Peek() *Item {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.items) <= 0 {
		return nil
	} else {
		return &s.items[len(s.items)-1]
	}
}
