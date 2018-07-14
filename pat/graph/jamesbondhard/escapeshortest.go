package main

import (
	"os"
	"strings"
	"github.com/ghjan/algorithms/graph/escape"
	"fmt"
)

/*
07-图5 Saving James Bond - Hard Version
https://pintia.cn/problem-sets/951072707007700992/problems/985411848322359296
 */
func solveEscapeShortest() {
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"007hard_case_1.txt", "007hard_case_2.txt"} //
	radius := float64(15.0 / 2.0)
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		graph, cords := escape.BuildGraphForBond(filename, 100, 100, radius)
		shortestTotalWeight, shortestPathSlice := escape.SolveEscapeShortest(graph, cords)
		if shortestPathSlice != nil {
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

func main() {
	solveEscapeShortest()

}
