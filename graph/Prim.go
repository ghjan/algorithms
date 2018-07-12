package graph

import (
	"github.com/kataras/iris/core/errors"
)

//Prim 邻接矩阵存储 - Prim最小生成树算法
//返回:
//mst, sum, parent, err 分别表示最小生成树，加权和，边的起始点数组，错误信息
func (mg *MGraph) Prim(start int) ([]VertexType, int, []int, error) {
	index := 0
	sum := 0
	mst := make([]VertexType, mg.vexNum, mg.vexNum)
	parent, dist := initDist(mg, start)

	//将初始点收录进MST
	mst[index] = mg.vexes[start]
	index++

	dist[start] = 0 //delete vex start（权重为0表示不再考虑）

	for i := 0; i < mg.vexNum; i++ {
		if start == i {
			continue
		}
		//next vex：MinV

		minV := mg.findMinDist(dist)
		if minV < 0 {
			break
		}
		sum += dist[minV]
		mst[index] = mg.vexes[minV]
		index++
		dist[minV] = 0 //delete vex minV

		//update parent和dist
		for W := 0; W < mg.vexNum; W++ {
			if dist[W] != 0 && mg.matrix[minV][W] < dist[W] { //若W未被收录，若收录MinV使得dist[W]变小
				dist[W] = mg.matrix[minV][W]
				parent[W] = minV
			}
		}
	}
	if index < mg.vexNum {
		return mst, sum, parent, errors.New("MST中收的顶点不到|V|个")
	}

	return mst, sum, parent, nil
}

func initArray(n, value int) []int {
	arr := make([]int, n, n)
	for i := 0; i < n; i++ {
		arr[i] = value
	}
	return arr
}

//initDist 返回parent和dist
//（parent放入每个边的起始点先假定为start,dist放入该到i的边的权重
func initDist(mg *MGraph, start int) ([]int, []int) {
	parent := initArray(mg.vexNum, start) //每个边的起始点先假定为start
	dist := initArray(mg.vexNum, MaxIntValue)

	//初始化dist, 放入start到i的边的权重
	for i := 0; i < mg.vexNum; i++ {
		dist[i] = mg.matrix[start][i] //v
	}
	return parent, dist
}

//findMinDist 返回未被收录顶点中dist最小者
//返回 minV
//    最小dist对应的顶点minV（dist[minV]最小）
//     -1：表示没有找到最小dist
func (mg *MGraph) findMinDist(dist []int) int {
	min := MaxIntValue
	minV := -1
	for v := 0; v < mg.vexNum; v++ {
		if dist[v] != 0 && dist[v] < min { //若V未被收录，且dist[V]更小
			min = dist[v] //更新最小距离
			minV = v      //更新对应顶点
		}
	}
	return minV
}
