package linkedlist

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func buildDigitalItemLinkedList(filename string) (*DigitalItemLinkedList, int32, int32, error) {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, -1, -1, nil
	}
	defer fi.Close()

	var list DigitalItemLinkedList
	br := bufio.NewReader(fi)
	var N, K int32
	index := 0
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if index == 0 { //第一行
			headInfoLit := strings.Split(string(a), " ")
			temp, _ := strconv.ParseInt(headInfoLit[0], 10, 32)
			headIndex := int32(temp)
			temp, _ = strconv.ParseInt(headInfoLit[1], 10, 32)
			N = int32(temp)
			temp, _ = strconv.ParseInt(headInfoLit[2], 10, 32)
			K = int32(temp)
			list = DigitalItemLinkedList{}
			list.head = headIndex
		} else {
			processAddNode(list, string(a))
		}
		index++
	}
	return &list, N, K, nil
}
func TestDigitalItemLinkedList_AddNode(t *testing.T) {
	list := DigitalItemLinkedList{}
	list.Init()
	nodeInfoSlice := []string{"00000 4 99999", "00100 1 12309", "68237 6 -1", "33218 3 00000", "99999 5 68237", "12309 2 33218"}
	list.head = 100
	for _, nodeInfo := range nodeInfoSlice {
		processAddNode(list, nodeInfo)
	}
	assert.Equal(t, int32(100), list.Head())
	assert.Equal(t, int32(6), list.Size())
	assert.Equal(t, false, list.IsEmpty())

}

func processAddNode(list DigitalItemLinkedList, nodeInfo string) int32 {
	nodeInfoList := strings.Split(nodeInfo, " ")
	var pos, content, next int32
	temp, _ := strconv.ParseInt(nodeInfoList[0], 10, 32)
	pos = int32(temp)
	temp, _ = strconv.ParseInt(nodeInfoList[1], 10, 32)
	content = int32(temp)
	temp, _ = strconv.ParseInt(nodeInfoList[2], 10, 32)
	next = int32(temp)
	list.AddNode(content, pos, next)
	return pos
}

func TestDigitalItemLinkedList_Reverse(t *testing.T) {
	fileName := "reverselinkedlist_case_1.txt"

	GOPATH := os.Getenv("GOPATH")
	filename := strings.Join([]string{GOPATH, "bin", fileName}, "/")
	if list, _, K, err := buildDigitalItemLinkedList(filename); err == nil {
		list.Reverse(K)
		list.browse(func(nodeIndex int32) {
			node := list.Nodes[nodeIndex]
			fmt.Printf("%s %d %s", convertToAddress(nodeIndex), node.content, convertToAddress(node.next))
		})
	} else {
		fmt.Println(err)
	}
}

var ZeroString = []string{"00000", "0000", "000", "00", "0"}

func convertToAddress(index int32) string {
	result := strconv.Itoa(int(index))
	if len(result) < 5 {
		result = ZeroString[len(result)] + result
	}
	return result
}
