package main

import (
	"os"
	"strings"
	"github.com/ghjan/algorithms/hashtable/inthashtable"
)

func main() {
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"hashing_case_1.txt"}
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		positionSlice := inthashtable.CreateTableForHashing(filename, false)
		inthashtable.PrintPosition(positionSlice)
	}
}
