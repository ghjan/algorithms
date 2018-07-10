package main

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

/*
邻接矩阵存储 - 有权图的单源最短路算法 Dijkstra算法
邻接矩阵存储 - 有权图的单源最短路算法
邻接矩阵存储 - 多源最短路算法 Floyd
邻接矩阵存储 - Prim最小生成树算法
*/
const (
	//最大整数值
	MaxIntValue int = int(^uint(0)>>1) / 10000
)

//VertexType 顶点数值类型
type VertexType string

//MGraph 邻接矩阵存储的有权图
type MGraph struct {
	vexs    []VertexType //顶点集合
	vexNum  int          //顶点数量
	edgeNum int          //边数量
	matrix  [][]int      //邻接矩阵 保存路径长度(权重）
	lock    sync.RWMutex // 保证线程安全
}

//PrintMatrix 显示矩阵
func (mg MGraph) PrintMatrix(l int) {
	for i := 0; i < l; i++ {
		fmt.Println(mg.matrix[i])
	}
}

//InitMGraph 初始化 边的权重全部设置为MaxIntValue
func (mg *MGraph) InitMGraph(vexs []VertexType) {
	mg.vexs = vexs
	mg.vexNum = len(vexs)
	for i := 0; i < mg.vexNum; i++ {
		var sl []int
		for j := 0; j < mg.vexNum; j++ {
			if i != j {
				sl = append(sl, MaxIntValue)
			} else {
				sl = append(sl, 0)
			}
		}
		mg.matrix = append(mg.matrix, sl)
	}
}

//AddEdge 增加边(无向图）
func (mg *MGraph) AddEdge(u, v, weight int) error {
	mg.lock.Lock()
	defer mg.lock.Unlock()
	// 首次建立图
	if len(mg.matrix) == 0 {
		return errors.New("Please use InitMGraph at first.")
	}
	mg.matrix[u][v] = weight // 建立 u->v 的边
	mg.matrix[v][u] = weight // 由于是无向图，同时存在 v->u 的边
	return nil
}

//AddEdgeDirection 增加边(有向图）
func (mg *MGraph) AddEdgeDirection(u, v, weight int) error {
	mg.lock.Lock()
	defer mg.lock.Unlock()
	// 首次建立图
	if len(mg.matrix) == 0 {
		return errors.New("Please use InitMGraph at first.")
	}
	mg.matrix[u][v] = weight // 建立 u->v 的边
	return nil
}

//DFS 深度遍历
func (mg *MGraph) DFS(operationFunc func(i int)) {
	visit := make([]bool, mg.vexNum, mg.vexNum)
	//fmt.Println(visit)
	visit[0] = true
	mg.dfs(&visit, 0, operationFunc)
}

func (mg *MGraph) dfs(visit *[]bool, i int, operationFunc func(i int)) {
	//fmt.Println(mg.vexs[i])
	operationFunc(i)
	for j := 0; j < mg.vexNum; j++ {
		if mg.matrix[i][j] != MaxIntValue && !(*visit)[j] {
			(*visit)[j] = true
			mg.dfs(visit, j, operationFunc)
		}
	}
}

//BFS 广度遍历
func (mg *MGraph) BFS(operationFunc func(v VertexType)) {
	listQ := list.New()
	visit := make([]bool, mg.vexNum, mg.vexNum)

	//first push
	visit[0] = true
	listQ.PushBack(0)

	for listQ.Len() > 0 {
		index := listQ.Front()
		//fmt.Println(mg.vexs[index.Value.(int)])
		operationFunc(mg.vexs[index.Value.(int)])
		for i := 0; i < mg.vexNum; i++ {
			if !visit[i] && mg.matrix[index.Value.(int)][i] != MaxIntValue {
				visit[i] = true
				listQ.PushBack(i)
			}
		}
		listQ.Remove(index)
	}
}

func (mg *MGraph) getPosition(ch VertexType) int {
	for i := 0; i < mg.vexNum; i++ {
		if mg.vexs[i] == ch {
			return i
		}
	}
	return -1
}
