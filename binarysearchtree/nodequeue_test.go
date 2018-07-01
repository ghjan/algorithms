package binarysearchtree

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"strconv"
)

func initQueue() *NodeItemQueue {
	var queueLocal NodeItemQueue
	if queueLocal.items == nil {
		queueLocal = NodeItemQueue{}
		queueLocal.New()
	}
	var i int
	for i = 0; i < 15; i++ {
		newNode := &Node{i, strconv.Itoa(i), nil, nil}
		queueLocal.Enqueue(newNode)
	}

	if size := queueLocal.Size(); size != i {
		fmt.Printf("Error:Wrong count, expected %d and got %d\n", i, size)
	}
	return &queueLocal
}

func TestEnqueue(t *testing.T) {

	queueLocal := initQueue()

	if size := queueLocal.Size(); size != 15 {
		t.Errorf("Wrong count, expected %d and got %d", 15, size)
	}
}

func TestDequeue(t *testing.T) {
	queueLocal := initQueue()
	if size := len(queueLocal.items); size != 15 {
		t.Errorf("Wrong count, expected %d and got %d", 15, size)
	}

	item0 := queueLocal.Dequeue()
	assert.EqualValues(t, item0.key, 0, fmt.Sprintf("Wrong value:%d, expected 0", item0.key))
	assert.EqualValues(t, item0.value, "0", fmt.Sprintf("Wrong value:%s, expected 0", item0.value))
	queueLocal.Dequeue() //1
	queueLocal.Dequeue() //2
	item3 := queueLocal.Dequeue()
	assert.EqualValues(t, item3.key, 3, fmt.Sprintf("Wrong value:%d, expected 3", item3.key))
	assert.EqualValues(t, item3.value, "3", fmt.Sprintf("Wrong value:%s, expected 3", item3.value))
	if size := len(queueLocal.items); size != (15 - 4) {
		t.Errorf("Wrong count, expected 0 and got %d", size)
	}

	if queueLocal.IsEmpty() {
		t.Errorf("IsEmpty should return false")
	}
}
