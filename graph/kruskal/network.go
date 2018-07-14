package kruskal

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/ghjan/algorithms/set"
)

func buildNetworkGraph(filename string, isZeroBased bool) (*Graph, set.UnionFindSet) {

	graph := &Graph{}
	var ufs set.UnionFindSet

	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return &Graph{}, ufs
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
			ufs = set.InitializationUFS(N)

		} else if i <= M { //边的数据输入 u, v
			array2 := strings.Split(string(a), " ")
			u, _ := strconv.Atoi(array2[0])
			v, _ := strconv.Atoi(array2[1])
			//网络是双向的
			graph.AddEdge(u, v, 1, false)
			graph.AddEdge(v, u, 1, false)

			if isZeroBased {
				ufs.InputConnection(u+1, v+1)
			} else {
				ufs.InputConnection(u, v)
			}
		} else {
			break
		}
		i++
	}

	return graph, ufs
}
func SolveConnectedComponents(filename string, isZeroBased bool) {
	graph, ufs := buildNetworkGraph(filename, isZeroBased)
	_, roots := ufs.CheckNetwork()

	//每个顶点的边 排序
	for _, vertex := range graph.Vertices {
		if !sort.IsSorted(vertex.Edges) {
			sort.Sort(vertex.Edges)
		}
	}
	for _, root := range roots {
		subNetwork := "{ "
		graph.DepthFirstSearch(root, func(vertex int) bool{
			subNetwork += strconv.Itoa(vertex) + " "
			return false
		})
		subNetwork += "}"
		fmt.Println(subNetwork)
	}
	for _, root := range roots {
		subNetwork := "{ "
		graph.BreadthFirstSearch(root, func(vertex int) {
			subNetwork += strconv.Itoa(vertex) + " "
		})
		subNetwork += "}"
		fmt.Println(subNetwork)
	}
}
