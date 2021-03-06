package inthashtable

import (
	"os"
	"strings"
	"testing"
)

func TestSolveHashing(t *testing.T) {
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"hashing_case_1.txt"}
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		positionSlice := CreateTableForHashing(filename, false)
		PrintPosition(positionSlice)
	}
}
