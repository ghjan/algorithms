package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ghjan/algorithms/pat/graph/kruskal"
)

/*
公路村村通
https://pintia.cn/problem-sets/951072707007700992/problems/987660874267004928
 */

func buildVilGraph(filename string) kruskal.Graph {

	graph := kruskal.Graph{}
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return kruskal.Graph{}
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
				graph.Vertices = append(graph.Vertices, &kruskal.Vertex{strconv.Itoa(i), nil, false})
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

func main() {
	start := time.Now()
	testVilvilPath()
	dis := time.Now().Sub(start).Nanoseconds()
	fmt.Printf("%s, detailed:%s, TimeoutWarning using %d ns", "testVilvilPath", "Total", dis)
}
