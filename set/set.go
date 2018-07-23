// Package set creates a ItemSet data structure for the Item type
package set

import "github.com/cheekybits/genny/generic"

// Item the type of the Set
type Item generic.Type

// ItemSet the set of Items
type ItemSet struct {
	items map[Item]bool
}

// Add adds a new element to the Set. Returns a pointer to the Set.
func (s *ItemSet) Add(t Item) *ItemSet {
	if s.items == nil {
		s.items = make(map[Item]bool)
	}
	_, ok := s.items[t]
	if !ok {
		s.items[t] = true
	}
	return s
}

// Clear removes all elements from the Set
func (s *ItemSet) Clear() {
	s.items = make(map[Item]bool)
}

// Delete removes the Item from the Set and returns Has(Item)
func (s *ItemSet) Delete(item Item) bool {
	_, ok := s.items[item]
	if ok {
		delete(s.items, item)
	}
	return ok
}

// Has returns true if the Set contains the Item
func (s *ItemSet) Has(item Item) bool {
	_, ok := s.items[item]
	return ok
}

// Items returns the Item(s) stored
func (s *ItemSet) Items() []Item {
	var items []Item
	for i := range s.items {
		items = append(items, i)
	}
	return items
}

// Size returns the size of the set
func (s *ItemSet) Size() int {
	return len(s.items)
}

//Intersect 两个集合的交集
func (s ItemSet) Intersect(target ItemSet) (ItemSet) {
	var result ItemSet
	for i := range s.items {
		if target.Has(i) {
			result.Add(i)
		}
	}
	return result

}

//Minus 两个集合的差 s-target
func (s ItemSet) Minus(target ItemSet) (ItemSet) {
	var result ItemSet
	for i := range s.items {
		if !target.Has(i) {
			result.Add(i)
		}
	}
	return result

}

//Union 两个集合的并集
func (s ItemSet) Union(target ItemSet) (ItemSet) {
	var result ItemSet
	for i := range s.items {
		result.Add(i)
	}
	for j := range target.items {
		if !result.Has(j) {
			result.Add(j)
		}
	}
	return result

}

//IntSet 转换为整数集合
func IntSet(items []int) (ItemSet) {
	var result ItemSet
	for _, item := range items {
		result.Add(item)
	}
	return result
}

//StringSet 转换为整数集合
func StringSet(items []string) (ItemSet) {
	var result ItemSet
	for _, item := range items {
		result.Add(item)
	}
	return result
}
