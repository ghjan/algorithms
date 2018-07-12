package travelplan

import (
	"testing"
	"os"
	"strings"
	"fmt"
	"github.com/stretchr/testify/assert"
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
	assert.Equal(t, 3, totalLength)
	assert.Equal(t, 40, totalCost)
	fmt.Printf("%d %d", totalLength, totalCost)
}
