package escape

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/ghjan/algorithms/graph/kruskal"
)

/*
利用图论知识解决逃脱问题

 */
//Cordinate2d 二维坐标
type Cordinate2d struct {
	x      int
	y      int
	radius int
	escape bool
}

//Distance 两个坐标之间的距离
func (cord Cordinate2d) Distance(target Cordinate2d) float64 {
	return math.Sqrt(float64((cord.x-target.x)*(cord.x-target.x)+(cord.y-target.y)*(cord.y-target.y))) - float64(cord.radius+target.radius)
}

//CanEscape 是否能够逃脱
//width, length分别表示湖的宽/长
//D是007能够跳跃的最大距离
func (cord Cordinate2d) CanEscape(width, length, D int) bool {
	rightBank := width / 2
	leftBank := -1 * rightBank
	topBank := length / 2
	bottomeBank := -1 * topBank
	if cord.radius == 0 {
		return rightBank-cord.x <= D || cord.x-leftBank <= D || topBank-cord.y <= D || cord.y-bottomeBank <= D
	} else {
		return rightBank-cord.x-cord.radius <= D || cord.x-cord.radius-leftBank <= D || topBank-cord.radius-cord.y <= D || cord.y-cord.radius-bottomeBank <= D
	}
}

//BuildGraphForBond 构建graph对象和cords分片
func BuildGraphForBond(filename string, width, lendth, radius int) (*kruskal.Graph, []Cordinate2d) {
	graph := kruskal.Graph{}
	var cords []Cordinate2d
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return &kruskal.Graph{}, cords
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	i := 0
	var N, D int
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { //顶点数量， 边数量
			array := strings.Split(string(a), " ")
			N, _ = strconv.Atoi(array[0]) //节点数量
			D, _ = strconv.Atoi(array[1]) //007最大跳跃距离
			for i := 0; i < N+1; i++ {
				graph.Vertices = append(graph.Vertices, &kruskal.Vertex{strconv.Itoa(i), nil, false})
			}
			cords = append(cords, Cordinate2d{0, 0, radius, false})

		} else if i <= N { //边的数据输入 start, end
			array2 := strings.Split(string(a), " ")
			x, _ := strconv.Atoi(array2[0])
			y, _ := strconv.Atoi(array2[1])
			cords = append(cords, Cordinate2d{x, y, 0, false})
			//处理新输入的坐标点，和原先的每个坐标点是否能够成为有连线的边（距离<=D)
			processCordAndGraph(graph, cords, D)
		} else {
			break
		}
		i++
	}
	for i = 0; i < len(cords); i++ {
		cords[i].escape = cords[i].CanEscape(width, lendth, D)
	}
	return &graph, cords

}

//processCordAndGraph 处理新输入的坐标点，和原先的每个坐标点是否能够成为有连线的边（距离<=D)
func processCordAndGraph(graph kruskal.Graph, cords []Cordinate2d, D int) {
	lastCord := cords[len(cords)-1]
	for i := 0; i < len(cords)-1; i++ {
		distance := lastCord.Distance(cords[i])
		if distance <= float64(D) {
			//双向的
			//顶点index已经从0开始
			graph.AddEdge(i, len(cords)-1, int(distance), false)
			graph.AddEdge(len(cords)-1, i, int(distance), false)
		}
	}
}

//SolveEscape 解决是否能够逃脱
func SolveEscape(graph *kruskal.Graph, cords []Cordinate2d) bool {
	result := false
	graph.DepthFirstSearch(0, func(vertexIndex int) bool {
		if cords[vertexIndex].escape {
			result = true
			return true // stop
		}
		return false
	})

	return result
}
