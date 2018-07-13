package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ghjan/algorithms/pat/graph/kruskal"
)

func earliestTest(graph *kruskal.Graph) (int, error) {
	if earliest, _, err := graph.Earliest(nil); err == nil {
		return earliest[len(earliest)-1], nil
	} else {
		return -1, errors.New("Impossible")
	}
}

func buildGraph(filename string) *kruskal.Graph {
	graph := kruskal.Graph{}
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return &kruskal.Graph{}
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
			//单向的
			graph.AddEdge(start, end, weight, false)
		} else {
			break
		}
		i++
	}
	return &graph

}

func main() {
	GOPATH := os.Getenv("GOPATH")
	fileList := []string{"howlongtake_case_1.txt", "howlongtake_case_2.txt"}
	for _, f := range fileList {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		howLongTest(filename)
	}
}
func howLongTest(filename string) {
	graph := buildGraph(filename)
	if howLong, err := earliestTest(graph); err == nil {
		fmt.Println(howLong)
	} else {
		fmt.Println(err)
	}
}
