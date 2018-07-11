package graph

import (
	"fmt"
)

//Prim 邻接矩阵存储 - Prim最小生成树算法
func (mg *MGraph) Prim(start int) {
	index := 0
	sum := 0
	prims := make([]VertexType, mg.vexNum, mg.vexNum)
	var weights [][]int //[[0 0] [0 5] [0 3] [0 9] [0 9]]
	for i := 0; i < mg.vexNum; i++ {
		sl := make([]int, 0, mg.vexNum)
		for j := 0; j < 2; j++ {
			sl = append(sl, MaxIntValue)
		}
		weights = append(weights, sl)
	}

	prims[index] = mg.vexs[start]
	index++

	//next vex
	for i := 0; i < mg.vexNum; i++ {
		weights[i][0] = start               //k
		weights[i][1] = mg.matrix[start][i] //v
	}

	//delete vex
	weights[start][1] = 0

	for i := 0; i < mg.vexNum; i++ {
		//fmt.Println(weights)
		if start == i {
			continue
		}

		min := MaxIntValue
		next := 0
		for j := 0; j < mg.vexNum; j++ {
			if weights[j][1] != 0 && weights[j][1] < min {
				min = weights[j][1]
				next = j
			}
		}

		fmt.Println(mg.vexs[weights[next][0]], mg.vexs[next], "权重", weights[next][1])
		sum += weights[next][1]
		prims[index] = mg.vexs[next]
		index++

		//delete vex
		weights[next][1] = 0

		//update
		for j := 0; j < mg.vexNum; j++ {
			if weights[j][1] != 0 && mg.matrix[next][j] < weights[j][1] {
				weights[j][1] = mg.matrix[next][j]
				weights[j][0] = next
			}
		}
	}

	fmt.Println("sum:", sum)
	fmt.Println(prims)
}

//func main() {
//	fmt.Println("Prim")
//	var mg MGraph
//	var vexs = []string{"B", "A", "C", "D", "E"}
//	mg.vexNum = 5
//	mg.vexs = vexs
//
//	for i := 0; i < len(vexs); i++ {
//		for j := 0; j < len(vexs); j++ {
//			mg.matrix[i][j] = MAX_VALUE
//		}
//	}
//	CreateMGraph(&mg)
//	fmt.Println(mg.vexs)
//	BFS(&mg)
//	DFS(&mg)
//
//	//listmg := list.New()
//	prim(&mg, 0)
//	PrintWeight(mg, len(vexs))
//}
