package main

import (
	"os"
	"strings"
	"github.com/ghjan/algorithms/hashtable/inthashtable"
	"fmt"
	"github.com/ghjan/algorithms/graph/kruskal"
)

func main() {
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"hashinghard_case_1.txt"}
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		table := inthashtable.CreateTableForHashingHard(filename, true)
		graph := kruskal.BuildGraphFromHashTable(table)
		if result, _, err := graph.TopologicalSortConditional(nil); err == nil {
			for index, item := range result {
				if index < len(result)-1 {
					fmt.Printf("%s ", item.Label)
				} else {
					fmt.Printf("%s\n", item.Label)
				}
			}
		} else {
			fmt.Println(err)
		}
	}
}
