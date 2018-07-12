package graphall

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
	//"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/assert"
)

func initGraph() *Graph {
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
					graph.AddEdge(indexTo, indexFrom, weight, false)
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
	graph := initGraph()
	resultEdges := graph.KruskalMinimumSpanningTree()
	for _, edge := range resultEdges {
		if edge.isUsed {
			fmt.Printf("%s->%s(%d)\n", edge.FromVertex.Label, edge.ToVertex.Label, edge.Weight)
		}
	}
	graph.clearEdgesUseHistory()
}

func TestGraph_DepthFirstSearch(t *testing.T) {
	graph := initGraph()
	vertexes := graph.DepthFirstSearch(graph.Vertices[0])
	vertexesLabelString := ""
	for _, vertex := range vertexes {
		vertexesLabelString += " " + vertex.Label
		fmt.Printf("%s ", vertex.Label)
	}
	expectedVertexLabel := "A B D C F E G"
	assert.Equal(t, expectedVertexLabel, strings.TrimLeft(vertexesLabelString, " "))
	graph.clearVerticesVisitHistory()

}

func TestGraph_BreadthFirstSearch(t *testing.T) {
	graph := initGraph()
	vertexes := graph.BreadthFirstSearch(graph.Vertices[0])
	vertexesLabelString := ""
	for _, vertex := range vertexes {
		vertexesLabelString += " " + vertex.Label
		fmt.Printf("%s ", vertex.Label)
	}
	expectedVertexLabel := "A B D E C F G"
	assert.Equal(t, expectedVertexLabel, strings.TrimLeft(vertexesLabelString, " "))
	graph.clearVerticesVisitHistory()
}
