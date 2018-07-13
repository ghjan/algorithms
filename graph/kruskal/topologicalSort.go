package kruskal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func SolveEarliest(graph *Graph) (int, error) {
	if earliest, _, err := graph.Earliest(nil); err == nil {
		return earliest[len(earliest)-1], nil
	} else {
		return 0, errors.New("Impossible")
	}
}

func BuildGraphForToplogicalSort(filename string, isZeroIndex bool) *Graph {
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
			if isZeroIndex { //数据已经从0开始
				graph.AddEdge(start, end, weight, false)
			} else { //数据从1开始，需要先减去1
				graph.AddEdge(start-1, end-1, weight, false)
			}
		} else {
			break
		}
		i++
	}
	return &graph

}

func SolveHowLong(filename string) {
	graph := BuildGraphForToplogicalSort(filename, true)
	if howLong, err := SolveEarliest(graph); err == nil {
		fmt.Println(howLong)
	} else {
		fmt.Println(err)
	}
}

func SolveCrucialPath(filename string, isZeroIndex bool) {
	graph := BuildGraphForToplogicalSort(filename, isZeroIndex)
	if howLong, crucialPath, err := graph.CrucialPath(); err == nil {
		fmt.Println(howLong)
		for _, cp := range crucialPath {
			if !isZeroIndex {
				fmt.Printf("%d->%d\n", cp.FromVertex+1, cp.ToVertex+1)
			} else {
				fmt.Printf("%d->%d\n", cp.FromVertex, cp.ToVertex)
			}
		}
	} else {
		fmt.Println(0)
	}

}
