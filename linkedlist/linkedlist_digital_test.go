package linkedlist

import (
"os"
"strings"
"testing"

"github.com/stretchr/testify/assert"
)

func TestDigitalItemLinkedList_AddNode(t *testing.T) {
	list := DigitalItemLinkedList{}
	list.Init()
	nodeInfoSlice := []string{"00000 4 99999", "00100 1 12309", "68237 6 -1", "33218 3 00000", "99999 5 68237", "12309 2 33218"}
	list.head = 100
	for _, nodeInfo := range nodeInfoSlice {
		list.ProcessAddNode(nodeInfo)
	}
	assert.Equal(t, int32(100), list.head)
	assert.Equal(t, int32(6), list.nodes[68237].content)
	assert.Equal(t, int32(-1), list.nodes[68237].next)
	assert.Equal(t, int32(99999), list.nodes[0].next)
	assert.Equal(t, int32(6), list.size)
	assert.Equal(t, false, list.IsEmpty())

}

func TestDigitalItemLinkedList_Reverse(t *testing.T) {
	fileName := "reverselinkedlist_case_1.txt"
	GOPATH := os.Getenv("GOPATH")
	filename := strings.Join([]string{GOPATH, "bin", fileName}, "/")
	SolveReverseLinkedList(filename)
}
