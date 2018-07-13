package kruskal

import (
	"errors"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
	"testing"

	"bufio"
	"os"

	"github.com/stretchr/testify/assert"
)

func initGraph(directed bool) *Graph { //有环路
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

func initGraph2(directed bool) *Graph { //去掉A的环路
	vertexLablels := []string{"A", "B", "C", "D", "E", "F", "G"}
	graph := &Graph{}
	for _, vertexLabel := range vertexLablels {
		graph.Vertices = append(graph.Vertices, &Vertex{vertexLabel, nil, false})
	}

	edgesString := "0 1 2,0 3 1,1 3 3,1 4 10,2 5 5,3 2 2,3 4 2,3 5 8,3 6 4,4 6 6,6 5 1"
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

func buildGraph(filename string) *Graph {
	graph := Graph{}
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return &Graph{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	i := 0
	var N, M int
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { //顶点数量， 边数量
			array := strings.Split(string(a), " ")
			N, _ = strconv.Atoi(array[0])
			M, _ = strconv.Atoi(array[1])
			for i := 0; i < N; i++ {
				graph.Vertices = append(graph.Vertices, &Vertex{strconv.Itoa(i), nil, false})
			}

		} else if i <= M { //边的数据输入 start, end, weight
			array2 := strings.Split(string(a), " ")
			start, _ := strconv.Atoi(array2[0])
			end, _ := strconv.Atoi(array2[1])
			weight, _ := strconv.Atoi(array2[2])
			//单向的
			graph.AddEdge(start, end, weight, false)
		} else {
			break
		}
		i++
	}
	return &graph

}
func buildVilGraph(filename string) Graph {

	graph := Graph{}
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return Graph{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	i := 0
	var N, M int
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { //顶点数量， 边数量
			array := strings.Split(string(a), " ")
			N, _ = strconv.Atoi(array[0])
			M, _ = strconv.Atoi(array[1])
			for i := 0; i < N; i++ {
				graph.Vertices = append(graph.Vertices, &Vertex{strconv.Itoa(i), nil, false})
			}

		} else if i <= M { //边的数据输入 start, end, weight
			array2 := strings.Split(string(a), " ")
			start, _ := strconv.Atoi(array2[0])
			end, _ := strconv.Atoi(array2[1])
			weight, _ := strconv.Atoi(array2[2])
			//马路是双向的
			graph.AddEdge(start-1, end-1, weight, false)
			graph.AddEdge(end-1, start-1, weight, false)
		} else {
			break
		}
		i++
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

func testVilvilPath() {
	GOPATH := os.Getenv("GOPATH")
	f := "vilvilpath_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")

	graph := buildVilGraph(filename)
	isSparse := float64(len(graph.Vertices)) > 2*math.Sqrt(float64(graph.EdgeNum))
	TotalWeight := 0
	if isSparse {
		TotalWeight, _ = graph.Kruskal()
	} else {
		MST := graph.PrimMinimumSpanningTree(graph.Vertices[0])
		for _, edge := range MST {
			TotalWeight += edge.Weight
		}
	}
	fmt.Println(TotalWeight)
}

func TestVilPath(t *testing.T) {
	testVilvilPath()
}

func TestGraph_TopologicalSort(t *testing.T) {
	fmt.Println("----------TestGraph_TopologicalSort-------------")

	graph := initGraph(true) //环图
	if result, _, err := graph.TopologicalSort(nil); err == nil {
		assert.Equal(t, 0, len(result))
	} else {
		assert.Error(t, err)
	}

	graph = initGraph2(true)

	sortedString := ""
	if result, inVertexes, err := graph.TopologicalSort(func(vertexIndex int, isSectionEnd bool) {
		if isSectionEnd {
			sortedString += ";"
			fmt.Println()
		} else {
			vertex := graph.Vertices[vertexIndex]
			sortedString += vertex.Label + " "
			fmt.Printf("%s ", vertex.Label)
		}
	}); err == nil {
		expectedIn := []string{"", "0 2", "3 2", "0 1,1 3", "1 10,3 2", "2 5,3 8,6 1", "3 4,4 6"}
		for i := 0; i < len(graph.Vertices); i++ {
			assert.Equal(t, expectedIn[i], strings.TrimRight(inVertexes[i], " ,"))
		}
		assert.Equal(t, "A B D C E G F", strings.TrimRight(sortedString, " "))
		assert.Equal(t, len(graph.Vertices), len(result))
	} else {
		assert.Equal(t, "", err)
	}
}

func TestGraph_Earliest(t *testing.T) {
	fmt.Println("----------TestGraph_Earliest-------------")

	graph := initGraph2(true)

	expectedEarliest := strings.Split("0 2 5 2 12 18 11", " ") //[]int{0, 2, 5, 2, 12, 18, 11}
	earliestTest(t, graph, expectedEarliest)

	fmt.Println("----howlongtake_case_1----")
	GOPATH := os.Getenv("GOPATH")
	f := "howlongtake_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")
	graph = buildGraph(filename)
	expectedEarliest = strings.Split("0 6 4 5 7 7 16 14 18", " ")
	earliestTest(t, graph, expectedEarliest)

	fmt.Println("----howlongtake_case_2----")
	f = "howlongtake_case_2.txt"
	filename = strings.Join([]string{GOPATH, "bin", f}, "/")
	graph = buildGraph(filename)
	expectedEarliest = strings.Split("0 6 4 5 7 7 16 14 18", " ")
	err := earliestTest(t, graph, expectedEarliest)
	assert.Equal(t, true, err != nil)
	fmt.Println(err)

}
func earliestTest(t *testing.T, graph *Graph, expectedEarliest []string) error {
	sortedString := ""
	if earliest, topSort, err := graph.Earliest(func(vertexIndex int, isSectionEnd bool) {
		if isSectionEnd {
			sortedString += ";"
			fmt.Println()
		} else {
			vertex := graph.Vertices[vertexIndex]
			sortedString += vertex.Label + " "
			//fmt.Printf("%s ", vertex.Label)
		}
	}); err == nil {
		fmt.Println("-----earliest------")
		fmt.Println(earliest)
		for i := 0; i < len(earliest); i++ {
			assert.Equal(t, expectedEarliest[i], strconv.Itoa(earliest[i]))
		}
		assert.Equal(t, len(graph.Vertices), len(earliest))
		assert.Equal(t, len(graph.Vertices), len(topSort))
		return nil
	} else {
		return errors.New("Impossible")
	}
}

func TestGraph_Earliest2(t *testing.T) {
	fmt.Println("----------TestGraph_Earliest2-------------")
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"howlongtake_case_1.txt", "howlongtake_case_2.txt"}
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		SolveHowLong(filename)
	}
}

func TestGraph_CrucialPath(t *testing.T) {
	fmt.Println("----------TestGraph_CrucialPath-------------")
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"crucialpath_case_1.txt",}
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		SolveCrucialPath(filename, false)
	}

}
