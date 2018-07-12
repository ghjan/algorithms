package travelplan

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"sync"

	"bufio"
	"os"
	"strings"
)

/*
旅游规划
https://pintia.cn/problem-sets/951072707007700992/problems/985411926680264704
*/
const (
	//最大整数值
	MaxIntValue int = int(^uint(0)>>1) / 10000
)

//VertexType 顶点数值类型
type VertexType string

//TravelGraph 旅行图-邻接矩阵存储的有权图（权重有两个：路径长度，路径收费）
type TravelGraph struct {
	vexNum       int          //顶点数量
	edgeNum      int          //边数量
	matrixLength [][]int      //邻接矩阵 保存路径长度（权重1）
	matrixCost   [][]int      //邻接矩阵 保存路径收费（权重2）
	lock         sync.RWMutex // 保证线程安全
}

//InitMGraph 初始化 边的权重全部设置为MaxIntValue
func (mg *TravelGraph) InitMGraph(vexNum int) {
	mg.lock.Lock()
	defer mg.lock.Unlock()
	mg.vexNum = vexNum
	for i := 0; i < mg.vexNum; i++ {
		var s1, s2 []int
		for j := 0; j < mg.vexNum; j++ {
			if i != j {
				s1 = append(s1, MaxIntValue)
				s2 = append(s2, MaxIntValue)
			} else {
				s1 = append(s1, 0)
				s2 = append(s2, 0)
			}
		}
		mg.matrixLength = append(mg.matrixLength, s1)
		mg.matrixCost = append(mg.matrixCost, s2)
	}
}

//AddEdge 增加边(无向图）
func (mg *TravelGraph) AddEdge(u, v, length, cost int) error {
	mg.lock.Lock()
	defer mg.lock.Unlock()
	// 首次建立图
	if len(mg.matrixLength) == 0 {
		return errors.New("Please use InitMGraph at first.")
	}
	mg.matrixLength[u][v] = length // 建立 u->v 的边
	mg.matrixLength[v][u] = length // 由于是无向图，同时存在 v->u 的边

	if len(mg.matrixCost) == 0 {
		return errors.New("Please use InitMGraph at first.")
	}
	mg.matrixCost[u][v] = cost // 建立 u->v 的边
	mg.matrixCost[v][u] = cost // 由于是无向图，同时存在 v->u 的边

	return nil
}

func buildTravelGraph(filename string) (TravelGraph, int, int) {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return TravelGraph{}, -1, -1
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	i := 0
	var N, M, S, D int
	var mg TravelGraph
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { //顶点数量， 边数量
			array := strings.Split(string(a), " ")
			N, _ = strconv.Atoi(array[0])
			M, _ = strconv.Atoi(array[1])
			S, _ = strconv.Atoi(array[2])
			D, _ = strconv.Atoi(array[3])
			mg.InitMGraph(N)
		} else if i <= M { //边的数据输入 start, end, length, cost
			array2 := strings.Split(string(a), " ")
			start, _ := strconv.Atoi(array2[0])
			end, _ := strconv.Atoi(array2[1])
			length, _ := strconv.Atoi(array2[2])
			cost, _ := strconv.Atoi(array2[3])
			mg.AddEdge(start, end, length, cost)
		} else {
			break
		}
		i++
	}
	return mg, S, D
}

//Dijkstra 邻接矩阵存储 - 有权图的单源最短路算法 Dijkstra算法
func (mg *TravelGraph) Dijkstra(start int) ([]int, []int, []int) {

	var collected = make([]bool, mg.vexNum) //最短路径节点集合
	var path = make([]int, mg.vexNum)       //路径数组

	//init
	dist := mg.matrixLength[start] //路径长度数组
	cost := mg.matrixCost[start]   //路径费用数组
	collected[start] = true        //find start to start as true
	dist[start] = 0                //start to start length
	cost[start] = 0

	for i := 0; i < mg.vexNum; i++ {
		v := 0
		min := MaxIntValue
		// fmt.Println("-----------")
		// fmt.Println(dist, collected)
		//find next 贪心
		for j := 0; j < len(dist); j++ {
			if collected[j] == false {
				if dist[j] != MaxIntValue && dist[j] < min {
					min = dist[j]
					v = j
				}
			}
		}

		//set find
		collected[v] = true

		//update dist length
		for w := 0; w < mg.vexNum; w++ {
			if mg.matrixLength[v][w] != MaxIntValue && collected[w] == false {
				weight := min + mg.matrixLength[v][w]
				if weight < dist[w] {
					dist[w] = weight
					path[w] = v
					cost[w] = cost[v] + mg.matrixCost[v][w]
				} else if weight == dist[w] && mg.matrixCost[v][w] < MaxIntValue && cost[v]+mg.matrixCost[v][w] < cost[w] {
					path[w] = v
					cost[w] = cost[v] + mg.matrixCost[v][w]
				}
			}
		}

	}
	return dist, cost, path
}
