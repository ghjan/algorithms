package main

import (
	"github.com/ghjan/algorithms/graph/escape"
	"os"
	"strings"
	"fmt"
)

func main() {
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"007_case_1.txt", "007_case_2.txt"}
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		graph, cords := escape.BuildGraphForBond(filename, 100, 100, 15)
		canEscape := escape.SolveEscape(graph, cords)
		if canEscape {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}
