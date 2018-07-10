package graph

import (
	"fmt"
	"strconv"

	stack2 "github.com/ghjan/algorithms/stack"
)

//Dijkstra 邻接矩阵存储 - 有权图的单源最短路算法 Dijkstra算法
func (mg *MGraph) Dijkstra(start int) ([]int, []int) {

	var dist = make([]int, mg.vexNum) //路径长度数组
	var ss = make([]bool, mg.vexNum)  //最短路径节点集合
	var path = make([]int, mg.vexNum) //路径数组

	//init
	dist = mg.matrix[start]
	ss[start] = true //find start to start as true
	dist[start] = 0  //start to start length

	for i := 0; i < mg.vexNum; i++ {
		k := 0
		min := MaxIntValue
		// fmt.Println("-----------")
		// fmt.Println(dist, ss)
		//find next 贪心
		for j := 0; j < len(dist); j++ {
			if ss[j] == false && dist[j] != MaxIntValue && dist[j] < min {
				min = dist[j]
				k = j
			}
		}

		//set find
		ss[k] = true

		//update dist length
		for u := 0; u < mg.vexNum; u++ {
			if mg.matrix[k][u] != MaxIntValue && ss[u] == false {
				weight := min + mg.matrix[k][u]
				if weight < dist[u] {
					dist[u] = weight
					path[u] = k
				}
			}
		}

	}
	return dist, path
}

//GetPathDijkstra 获得到start到target的路径（基于Dijkstra算法）
func (mg *MGraph) GetPathDijkstra(path []int, start, target int) string {
	if path == nil || len(path) == 0 {
		_, path := mg.Dijkstra(start)
		return mg.GetPathDijkstra(path, start, target)
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
