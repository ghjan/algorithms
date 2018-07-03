package graph

import "fmt"

func (gg *MGraph) Dijkstra(start int) []int {

	var dist = make([]int, gg.vexNum) //路径长度数组
	var ss = make([]bool, gg.vexNum)  //最短路径节点集合

	//init
	dist = gg.matrix[start]
	ss[start] = true //find start to start
	dist[start] = 0  //start to start length

	for i := 0; i < gg.vexNum; i++ {
		k := 0
		min := MAX_VALUE
		fmt.Println("-----------")
		fmt.Println(dist, ss)
		//find next 贪心
		for j := 0; j < len(dist); j++ {
			if ss[j] == false && dist[j] != MAX_VALUE && dist[j] < min {
				min = dist[j]
				k = j
			}
		}

		//set find
		ss[k] = true

		//update dist length
		for u := 0; u < gg.vexNum; u++ {
			if gg.matrix[k][u] != MAX_VALUE && ss[u] == false {
				weight := min + gg.matrix[k][u]
				if weight < dist[u] {
					dist[u] = weight
				}
			}
		}

	}
	return dist
}
