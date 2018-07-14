package kruskal

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestSolveEscape(t *testing.T) {
	fmt.Println("----------TestSolveEscape----")
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"007_case_1.txt", "007_case_2.txt", "007hard_case_1.txt", "007hard_case_2.txt"}
	radius := float64(15.0/2.0)
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		graph, cords := BuildGraphForBond(filename, 100, 100, radius)
		canEscape := SolveCanEscape(graph, cords)
		if canEscape {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}

func TestSolveEscapeShortest(t *testing.T) {
	fmt.Println("----------TestSolveEscapeShortest----")
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"007hard_case_1.txt", } //"007hard_case_2.txt"
	radius := float64(15.0/2.0)
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		graph, cords := BuildGraphForBond(filename, 100, 100, radius)
		shortestTotalWeight, shortestPathSlice := SolveEscapeShortest(graph, cords)
		if shortestPathSlice != nil {
			fmt.Println(shortestTotalWeight)
			_, pathString2, _ := GetPathString(shortestPathSlice, cords, graph)
			fmt.Println(pathString2)
		} else {
			fmt.Println(0)
		}
	}
}
