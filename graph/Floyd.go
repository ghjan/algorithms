package graph

import "errors"

//Floyd 邻接矩阵存储 - 多源最短路算法 Floyd算法
//return: 路径加权距离矩阵, 路径上一结点矩阵
func (gg *MGraph) Floyd() ([][]int, [][]int, error) {
	var D, path [][]int //路径加权距离矩阵, 路径上一结点矩阵
	for i := 0; i < gg.vexNum; i++ {
		var sd, sp []int
		for j := 0; j < gg.vexNum; j++ {
			sd = append(sd, gg.matrix[i][j])
			sp = append(sp, -1)
		}
		D = append(D, sd)
		path = append(path, sp)
	}
	for k := 0; k < gg.vexNum; k++ {
		for i := 0; i < gg.vexNum; i++ {
			for j := 0; j < gg.vexNum; j++ {
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
