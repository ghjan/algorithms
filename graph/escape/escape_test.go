package escape

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestSolveEscape(t *testing.T) {
	fmt.Println("----------TestSolveEscape----")
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"007_case_1.txt", "007_case_2.txt"}
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		graph, cords := BuildGraphForBond(filename, 100, 100, 15)
		canEscape := SolveEscape(graph, cords)
		if canEscape {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}
