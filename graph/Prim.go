package graph

import (
	"fmt"
)

func (gg *MGraph) Prim(start int) {
	index := 0
	sum := 0
	prims := make([]VertexType, gg.vexNum, gg.vexNum)
	var weights [][]int //[[0 0] [0 5] [0 3] [0 9] [0 9]]
	for i := 0; i < gg.vexNum; i++ {
		sl := make([]int, 0, gg.vexNum)
		for j := 0; j <2 ; j++ {
			sl = append(sl, MAX_VALUE)
		}
		weights = append(weights, sl)
	}

	prims[index] = gg.vexs[start]
	index++

	//next vex
	for i := 0; i < gg.vexNum; i++ {
		weights[i][0] = start               //k
		weights[i][1] = gg.matrix[start][i] //v
	}

	//delete vex
	weights[start][1] = 0

	for i := 0; i < gg.vexNum; i++ {
		//fmt.Println(weights)
		if start == i {
			continue
		}

		min := MAX_VALUE
		next := 0
		for j := 0; j < gg.vexNum; j++ {
			if weights[j][1] != 0 && weights[j][1] < min {
				min = weights[j][1]
				next = j
			}
		}

		fmt.Println(gg.vexs[weights[next][0]], gg.vexs[next], "权重", weights[next][1])
		sum += weights[next][1]
		prims[index] = gg.vexs[next]
		index++

		//delete vex
		weights[next][1] = 0

		//update
		for j := 0; j < gg.vexNum; j++ {
			if weights[j][1] != 0 && gg.matrix[next][j] < weights[j][1] {
				weights[j][1] = gg.matrix[next][j]
				weights[j][0] = next
			}
		}
	}

	fmt.Println("sum:", sum)
	fmt.Println(prims)
}

//func main() {
//	fmt.Println("Prim")
//	var gg MGraph
//	var vexs = []string{"B", "A", "C", "D", "E"}
//	gg.vexNum = 5
//	gg.vexs = vexs
//
//	for i := 0; i < len(vexs); i++ {
//		for j := 0; j < len(vexs); j++ {
//			gg.matrix[i][j] = MAX_VALUE
//		}
//	}
//	initMGraph(&gg)
//	fmt.Println(gg.vexs)
//	BFS(&gg)
//	DFS(&gg)
//
//	//listgg := list.New()
//	prim(&gg, 0)
//	PrintMatrix(gg, len(vexs))
//}
