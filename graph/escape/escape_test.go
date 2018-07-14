package escape

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSolveEscape(t *testing.T) {
	fmt.Println("----------TestSolveEscape----")
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"007_case_1.txt", "007_case_2.txt", "007hard_case_1.txt", "007hard_case_2.txt"}
	radius := float64(15.0 / 2.0)
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
	fileList := []string{"007hard_case_1.txt", "007hard_case_2.txt"} //
	radius := float64(15.0 / 2.0)
	expectedWeight := []int{3, 0}
	for indexFile, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		graph, cords := BuildGraphForBond(filename, 100, 100, radius)
		shortestTotalWeight, shortestPathSlice := SolveEscapeShortest(graph, cords)
		if shortestPathSlice != nil {
			assert.Equal(t, expectedWeight[indexFile], shortestTotalWeight)
			fmt.Println(shortestTotalWeight + 1)
			for u := len(shortestPathSlice) - 1; u >= 1; u-- {
				v := u - 1
				if v >= 0 {
					toIndex := shortestPathSlice[v]
					fmt.Printf("%d %d\n", cords[toIndex].X, cords[toIndex].Y)
				} else {
					break
				}
			}
		} else {
			fmt.Println(0)
		}
	}
}
