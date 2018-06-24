package linkedlist_int

import (
	"testing"
	"fmt"
)

var list ItemLinkedList_int
func TestAppend(t *testing.T) {
	if !list.IsEmpty() {
		t.Errorf("Linked list should be empty")
	}
	list.Append(1)
	if list.IsEmpty() {
		t.Errorf("Linked list should not be empty")
	}
	if size := list.Size(); size != 1 {
		t.Errorf("Wrong count, expected 1 but got %d", size)
	}

	list.Append(2)
	list.Append(3)
	if size := list.Size(); size != 3 {
		t.Errorf("Wrong count, expected 3 but got %d", size)
	}
}

func TestRemoveAt(t *testing.T) {
	_, err := list.RemoveAt(1) // 删除 second
	if err != nil {
		t.Errorf("Unexcepted error: %s", err)
	}
	if size := list.Size(); size != 2 {
		t.Errorf("Wrong count, expected 2 but got %d", size)
	}
}

func TestInsert(t *testing.T) {
	// 测试插入到链表中间
	err := list.Insert(2, 2)
	if err != nil {
		t.Errorf("Unexcepted error: %s", err)
	}
	if size := list.Size(); size != 3 {
		t.Errorf("Wrong count, expected 3 but got %d", size)
	}
	// 测试插入到链表两侧
	err = list.Insert(0, 0)
	if err != nil {
		t.Errorf("Unexcepted error: %s", err)
	}
}

func TestIndexOf(t *testing.T) {
	if i := list.IndexOf(0); i != 0 {
		t.Errorf("Excepted postion 0 but got: %d", i)
	}
	if i := list.IndexOf(1); i != 1 {
		t.Errorf("Excepted postion 1 but got: %d", i)
	}
	if i := list.IndexOf(2); i != 2 {
		t.Errorf("Excepted postion 2 but got: %d", i)
	}
	if i := list.IndexOf(3); i != 3 {
		t.Errorf("Excepted postion 3 but got: %d", i)
	}
}

func TestHead(t *testing.T) {
	head := list.head
	content := fmt.Sprint(head.content)
	if content != "0" {
		t.Errorf("Excepted `zero` but got: %s", content)
	}
}
