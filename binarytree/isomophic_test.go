package binarytree

import (
	"testing"
	"os"
	"strings"
	"fmt"
)

func TestBinaryTree_Isomorphic(t *testing.T) {
	fileNames := [...]string{"binarytree_isomophic_case_1.txt"} // , "binarytree_isomophic_case_2.txt"

	GOPATH := os.Getenv("GOPATH")

	for _, f := range fileNames {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		SolveBTIsomorphic(filename)
		fmt.Println()
	}
}
