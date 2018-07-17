package main

import (
	"os"
	"strings"

	"github.com/ghjan/algorithms/stack"
)

func main() {

	GOPATH := os.Getenv("GOPATH")
	f := "treeranverseagain_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")

	stack.SolveTreeTranverseAgain(filename)
}
