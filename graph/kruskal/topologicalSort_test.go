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
		graph := BuildGraphFromHashTable(table)
		if result, inVertexes, err := graph.TopologicalSortConditional(nil); err == nil {
			fmt.Println("------result of TopologicalSort")
			for _, item := range result {
				fmt.Printf("%s ", item.Label)
			}
			fmt.Println("\n------inVertexes")
			for _, inv := range inVertexes {
				fmt.Printf("%s ", inv)
			}
			fmt.Println()
		} else {
			fmt.Println(err)
		}
	}

}
