package kruskal

import (
	"github.com/ghjan/algorithms/hashtable/inthashtable"
	"os"
	"strings"
	"fmt"
	"testing"
)

func TestSolveHashingHard(t *testing.T) {
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"hashinghard_case_1.txt"}
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		table := inthashtable.CreateTableForHashingHard(filename, true)
		graph := BuildGraphFromHashtable(table)
		if result, inVertexes, err := graph.TopologicalSort(nil); err == nil {
			for _, item := range result {
				fmt.Println(item)
			}
			for _, inv := range inVertexes {
				fmt.Println(inv)
			}
		}
	}

}
