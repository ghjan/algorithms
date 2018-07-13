package main

import (
	"os"
	"strings"
	"github.com/ghjan/algorithms/graph/kruskal"
)

func SolveConnectedComponents() {
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"connectedcomponents_case_1.txt"}
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		isZeroBased := true
		kruskal.SolveConnectedComponents(filename, isZeroBased)
	}
}

func main() {
	SolveConnectedComponents()
}
