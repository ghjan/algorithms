package sixdegree

import (
	"testing"
	"fmt"
	"os"
	"strings"
	"github.com/stretchr/testify/assert"
)

func TestSolveSixDegree(t *testing.T) {
	fmt.Println("----------TestSolveSixDegree----")
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"sixdegree_case_1.txt"}

	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		N, _, mp := BuildGraphForSixDegree(filename)
		result := SolveSixDegree(N, mp)
		assert.Equal(t, "7 8 9 10 10 10 10 9 8 7", result)
	}
}
