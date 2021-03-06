package main

import (
	"os"
	"strings"

	"github.com/ghjan/algorithms/graph/kruskal"
)

func main() {
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"howlongtake_case_1.txt", "howlongtake_case_2.txt"}
	isZeroIndex := true
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		kruskal.SolveHowLong(filename, isZeroIndex)

	}
}
