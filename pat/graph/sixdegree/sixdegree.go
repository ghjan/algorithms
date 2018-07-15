package main

import (
	"os"
	"strings"
	"github.com/ghjan/algorithms/graph/sixdegree"
)

func solveSixDegree() {

	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"sixdegree_case_1.txt"}

	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		N, _, bitMatrix := sixdegree.BuildGraphForSixDegree(filename)
		sixdegree.SolveSixDegree(N, bitMatrix)
	}
}

func main() {
	solveSixDegree()
}
