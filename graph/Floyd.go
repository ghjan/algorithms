package graph

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/ghjan/algorithms/stack"
)

//Floyd 邻接矩阵存储 - 多源最短路算法 Floyd算法
//return: 路径加权距离矩阵, 路径上一结点矩阵
func (mg *MGraph) Floyd() ([][]int, [][]int, error) {
	var D, path [][]int //路径加权距离矩阵, 路径上一结点矩阵
	for i := 0; i < mg.vexNum; i++ {
		var sd, sp []int
		for j := 0; j < mg.vexNum; j++ {
			sd = append(sd, mg.matrix[i][j])
			sp = append(sp, -1)
		}
		D = append(D, sd)
		path = append(path, sp)
	}
	for k := 0; k < mg.vexNum; k++ {
		for i := 0; i < mg.vexNum; i++ {
			for j := 0; j < mg.vexNum; j++ {
				if D[i][k]+D[k][j] < D[i][j] {
					D[i][j] = D[i][k] + D[k][j]
					if i == j && D[i][j] < 0 { /* 若发现负值圈 */
						return D, path, errors.New("发现负值圈") /* 不能正确解决，返回错误标记 */
					}
					path[i][j] = k
				}
			}
		}

	}
	return D, path, nil
}

//GetPathFloyd 获得start到target的路径（基于Floyd算法）
func (mg *MGraph) GetPathFloyd(path [][]int, start, target int) (string, error) {
	if path == nil || len(path) == 0 {
		if _, path, err := mg.Floyd(); err == nil {
			return mg.GetPathFloyd(path, start, target)
		} else {
			return "", err
		}
	}
	var stack stack.ItemStack
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
	return result, nil
}
