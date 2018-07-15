package main

import (
	"os"
	"strings"
	"github.com/ghjan/algorithms/binarytree"
)

func solveIsomorphic() {
	fileNames := [...]string{"binarytree_isomophic_case_1.txt", "binarytree_isomophic_case_2.txt"}

	GOPATH := os.Getenv("GOPATH")

	for _, f := range fileNames {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		binarytree.SolveBTIsomorphic(filename)
	}
}

func main() {
	solveIsomorphic()
}
