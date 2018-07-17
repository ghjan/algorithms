package stack

import (
"testing"
"os"
"strings"
)

func TestItemStack_CanPopSequence(t *testing.T) {
	GOPATH := os.Getenv("GOPATH")
	f := "popsequence_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")
	SolvePopSequence(filename)
}
