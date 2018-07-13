package main

import (
	"os"
	"strings"
	"github.com/ghjan/algorithms/graph/kruskal"
)

func main() {
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"crucialpath_case_1.txt"}
	isZeroIndex := false
	isDebug := false
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		kruskal.SolveCrucialPath(filename, isZeroIndex, isDebug)
	}
}
