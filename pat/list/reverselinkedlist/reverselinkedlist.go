package main

import (
	"os"
	"strings"
	"github.com/ghjan/algorithms/linkedlist"
)

func main() {
	fileName := "reverselinkedlist_case_1.txt"
	GOPATH := os.Getenv("GOPATH")
	filename := strings.Join([]string{GOPATH, "bin", fileName}, "/")
	linkedlist.SolveReverseLinkedList(filename)
}
