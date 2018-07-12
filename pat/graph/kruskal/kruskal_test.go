package kruskal

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initGraph(directed bool) *Graph {
	vertexLablels := []string{"A", "B", "C", "D", "E", "F", "G"}
	graph := &Graph{}
	for _, vertexLabel := range vertexLablels {
		graph.Vertices = append(graph.Vertices, &Vertex{vertexLabel, nil, false})
	}

	edgesString := "0 1 2,0 3 1,1 3 3,1 4 10,2 0 4,2 5 5,3 2 2,3 4 2,3 5 8,3 6 4,4 6 6,6 5 1"
	edgesInfoSlice := strings.Split(edgesString, ",")
	for _, edgeInfo := range edgesInfoSlice {
		edgeInfoSlice := strings.Split(edgeInfo, " ")
		var err error
		if indexFrom, err := strconv.Atoi(edgeInfoSlice[0]); err == nil {
			if indexTo, err := strconv.Atoi(edgeInfoSlice[1]); err == nil {
				if weight, err := strconv.Atoi(edgeInfoSlice[2]); err == nil {
					graph.AddEdge(indexFrom, indexTo, weight, false)
					if !directed {
						graph.AddEdge(indexTo, indexFrom, weight, false)
					}
				}
			}
		}
		if err != nil {
			fmt.Printf("create edge fail:%s\n", edgeInfo)
			fmt.Println(err)
		}

	}

	return graph
}

func TestGraph_KruskalMinimumSpanningTree(t *testing.T) {
	fmt.Println("----------TestGraph_KruskalMinimumSpanningTree-------------")
	graph := initGraph(false)
	resultEdges := graph.KruskalMinimumSpanningTree()
	weightString := ""
	totalWeight2 := 0
	for _, edge := range resultEdges {
		totalWeight2 += edge.Weight
		weightString += strconv.Itoa(edge.Weight) + " "
		fmt.Printf("%s->%s(%d)\n", graph.Vertices[edge.FromVertex].Label, graph.Vertices[edge.ToVertex].Label, edge.Weight)
	}
	assert.Equal(t, "1 1 2 2 2 4", strings.TrimRight(weightString, " "))
	assert.Equal(t, 12, totalWeight2)

}

func TestGraph_DepthFirstSearch(t *testing.T) {
	fmt.Println("----------TestGraph_DepthFirstSearch-------------")
	graph := initGraph(true)
	vertexesLabelString := ""
	vertexes := graph.DepthFirstSearch(0, func(vertexIndex int) {
		vertexesLabelString += " " + graph.Vertices[vertexIndex].Label
	})
	expectedVertexLabel := "A B D C F E G"
	assert.Equal(t, expectedVertexLabel, strings.TrimLeft(vertexesLabelString, " "))
	for _, vertex := range vertexes {
		fmt.Printf("%s ", graph.Vertices[vertex].Label)
	}
}

func TestGraph_BreadthFirstSearch(t *testing.T) {
	fmt.Println("----------TestGraph_BreadthFirstSearch-------------")
	graph := initGraph(true)
	vertexesLabelString := ""
	vertexes := graph.BreadthFirstSearch(0, func(vertexIndex int) {
		vertexesLabelString += " " + graph.Vertices[vertexIndex].Label
	})
	expectedVertexLabel := "A B D E C F G"
	assert.Equal(t, expectedVertexLabel, strings.TrimLeft(vertexesLabelString, " "))
	for _, vertex := range vertexes {
		fmt.Printf("%s ", graph.Vertices[vertex].Label)
	}
}

func TestGraph_Kruskal(t *testing.T) {
	graph := initGraph(false)
	totalWeight, resultEdges := graph.Kruskal()
	weightString := ""
	totalWeight2 := 0
	for _, edge := range resultEdges {
		totalWeight2 += edge.Weight
		weightString += strconv.Itoa(edge.Weight) + " "
		fmt.Printf("%s->%s(%d)\n", graph.Vertices[edge.FromVertex].Label, graph.Vertices[edge.ToVertex].Label, edge.Weight)
	}
	assert.Equal(t, "1 1 2 2 2 4", strings.TrimRight(weightString, " "))
	assert.Equal(t, totalWeight2, totalWeight)

}
