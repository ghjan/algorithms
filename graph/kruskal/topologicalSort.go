package kruskal

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"github.com/ghjan/algorithms/hashtable/inthashtable"
)

func solveEarliest(graph *Graph) (int, error) {
	if earliest, _, err := graph.Earliest(nil); err == nil {
		return earliest[len(earliest)-1], nil
	} else {
		return 0, errors.New("Impossible")
	}
}

func BuildGraphForTopologicalSort(filename string, isZeroIndex bool) *Graph {
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

		} else if i <= M { //边的数据输入 start, end, weight 含义：从start到end的边表示start是end的前驱节点
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

func BuildGraphFromHashTable(table inthashtable.IntHashTable) *Graph {
	graph := Graph{}

	i := 0
	N := table.TableSize
	graph.Vertices = make([]*Vertex, N, N)
	for i := 0; i < N; i++ { //顶点数量 N
		if table.Cells[i].Info != inthashtable.Legitimate || table.Cells[i].Data < 0 {
			graph.Vertices[i] = nil
			continue
		}
		graph.Vertices[i] = &Vertex{strconv.Itoa(table.Cells[i].Data), nil, false}
	}

	var firstGroup, secondGroup []int //第一集团， hashValue正好和i相同的
	var edgesMap map[int][]int = make(map[int][]int, N)
	for i = 0; i < N; i++ { //边的数据输入 start, end, weight 含义：从start到end的边表示start是end的前驱节点
		if table.Cells[i].Info != inthashtable.Legitimate || table.Cells[i].Data < 0 {
			continue
		}
		hashValue := table.Hash(table.Cells[i].Data) //原始的hash值
		if hashValue != i {
			secondGroup = append(secondGroup, i)
			if hashValue < i {
				for start := hashValue; start%table.TableSize < i; start++ {
					end := i
					if graph.Vertices[start] != nil {
						graph.AddEdge(start, end, 1, false)
						edgesMap[start] = append(edgesMap[start], end)
					}
				}
			} else {
				for start := hashValue; start != i; start = (start + 1) % table.TableSize {
					end := i
					if graph.Vertices[start] != nil {
						graph.AddEdge(start, end, 1, false)
						edgesMap[start] = append(edgesMap[start], end)
					}
				}
			}
		} else {
			firstGroup = append(firstGroup, i)
		}
	}
	//for _, first := range firstGroup {
	//	intersectGroup := set.IntSet(secondGroup).Minus(set.IntSet(edgesMap[first]))
	//	for _, second := range intersectGroup.Items() {
	//		if err := graph.AddEdge(first, second.(int), 1, false); err != nil {
	//			fmt.Println(err)
	//		}
	//	}
	//}

	return &graph

}

func SolveHowLong(filename string, isZeroIndex bool) {
	graph := BuildGraphForTopologicalSort(filename, isZeroIndex)
	if howLong, err := solveEarliest(graph); err == nil {
		fmt.Println(howLong)
	} else {
		fmt.Println(err)
	}
}

func SolveCrucialPath(filename string, isZeroIndex, isDebug bool) {
	graph := BuildGraphForTopologicalSort(filename, isZeroIndex)
	if howLong, crucialPath, err := graph.CrucialPath(isDebug); err == nil {
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
