package graph

import (
	"container/list"
	"fmt"
	stack2 "github.com/ghjan/algorithms/stack"
	"strconv"
)

/*
邻接矩阵存储 - 有权图的单源最短路算法 Dijkstra算法
邻接矩阵存储 - 有权图的单源最短路算法
邻接矩阵存储 - 多源最短路算法 Floyd
邻接矩阵存储 - Prim最小生成树算法
 */
const (
	MAX_VALUE int = int(^uint(0) >> 1)
)

type VertexType string // 顶点数值类型

type MGraph struct {
	vexs    []VertexType //顶点集合
	vexNum  int          //顶点数量
	edgeNum int          //边数量
	matrix  [][]int      //邻接矩阵 保存最小路径长度
}

func (gg MGraph) PrintMatrix(l int) {
	for i := 0; i < l; i++ {
		fmt.Println(gg.matrix[i])
	}
}

func (gg *MGraph) initMGraph(vexs []VertexType) {
	gg.vexs = vexs
	gg.vexNum = len(vexs)
	for i := 0; i < gg.vexNum; i++ {
		var sl []int
		for j := 0; j < gg.vexNum; j++ {
			sl = append(sl, MAX_VALUE)
		}
		gg.matrix = append(gg.matrix, sl)
	}
}

//深度遍历
func (gg *MGraph) DFS(operationFunc func(i int)) {
	visit := make([]bool, gg.vexNum, gg.vexNum)
	//fmt.Println(visit)
	visit[0] = true
	gg.dfs(&visit, 0, operationFunc)
}

func (gg *MGraph) dfs(visit *[]bool, i int, operationFunc func(i int)) {
	//fmt.Println(gg.vexs[i])
	operationFunc(i)
	for j := 0; j < gg.vexNum; j++ {
		if gg.matrix[i][j] != MAX_VALUE && !(*visit)[j] {
			(*visit)[j] = true
			gg.dfs(visit, j, operationFunc)
		}
	}
}

//广度遍历
func (gg *MGraph) BFS(operationFunc func(v VertexType)) {
	listQ := list.New()
	visit := make([]bool, gg.vexNum, gg.vexNum)

	//first push
	visit[0] = true
	listQ.PushBack(0)

	for listQ.Len() > 0 {
		index := listQ.Front()
		//fmt.Println(gg.vexs[index.Value.(int)])
		operationFunc(gg.vexs[index.Value.(int)])
		for i := 0; i < gg.vexNum; i++ {
			if !visit[i] && gg.matrix[index.Value.(int)][i] != MAX_VALUE {
				visit[i] = true
				listQ.PushBack(i)
			}
		}
		listQ.Remove(index)
	}
}

func (gg *MGraph) getPosition(ch VertexType) int {
	for i := 0; i < gg.vexNum; i++ {
		if gg.vexs[i] == ch {
			return i
		}
	}
	return -1
}

//GetPathDijkstra 获得到start到target的路径
func (g *MGraph) GetPathDijkstra(path []int, start, target int) string {
	if path == nil || len(path) == 0 {
		_, path := g.Dijkstra(start)
		return g.GetPathDijkstra(path, start, target)
	}
	var stack stack2.ItemStack
	stack.New()
	stack.Push(target)
	for pathPrev := path[target]; pathPrev > 0; pathPrev = path[pathPrev] {
		if pathPrev < 0 {
			break
		}
		stack.Push(pathPrev)
	}
	result := strconv.Itoa(start) + " "
	for node := stack.Pop(); node != nil; node = stack.Pop() {
		result += fmt.Sprintf("%d ", (*node).(int))
		if stack.IsEmpty() {
			break
		}
	}
	return result
}

//GetPathDijkstra 获得start到target的路径
func (g *MGraph) GetPathFloyd(path [][]int, start, target int) string {
	if path == nil || len(path) == 0 {
		_, path := g.Dijkstra(start)
		return g.GetPathDijkstra(path, start, target)
	}
	var stack stack2.ItemStack
	stack.New()
	stack.Push(target)
	for pathPrev := path[start][target]; pathPrev > 0; pathPrev = path[start][pathPrev] {
		if pathPrev < 0 {
			break
		}
		stack.Push(pathPrev)
	}
	result := strconv.Itoa(start) + " "
	for node := stack.Pop(); node != nil; node = stack.Pop() {
		result += fmt.Sprintf("%d ", (*node).(int))
		if stack.IsEmpty() {
			break
		}
	}
	return result
}
