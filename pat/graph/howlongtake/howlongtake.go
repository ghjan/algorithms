package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/ghjan/algorithms/graph/kruskal"
)

func main() {
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"howlongtake_case_1.txt", "howlongtake_case_2.txt"}
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		howLongTest(filename)
	}
}
func howLongTest(filename string) {
	graph := kruskal.BuildGraphForToplogicalSort(filename)
	if howLong, err := kruskal.SolveEarliest(graph); err == nil {
		fmt.Println(howLong)
	} else {
		fmt.Println(err)
	}
}
