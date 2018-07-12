package travelplan

import (
	"testing"
	"os"
	"strings"
	"fmt"
)

func findBestTravel() (int, int) {
	GOPATH := os.Getenv("GOPATH")
	f := "travelplan_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")
	tg, S, D := buildTravelGraph(filename)
	dist, cost, _ := tg.Dijkstra(S)
	return dist[D], cost[D]
}
func TestTravelGraph_Dijkstra(t *testing.T) {
	totalLength, totalCost := findBestTravel()
	fmt.Printf("%d %d", totalLength, totalCost)
}
