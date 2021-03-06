package cbt

import "math"
/*
CBT :Complete Binary Search Tree
 */

//SolveCBT 产生完全二叉搜索树
//初始化调用binarysearchtree.SolveCBT(A, T, 0, N-1, 0)
//产生的结果保存在T分片里面
func SolveCBT(A, T []int, ALeft, ARight, TRoot int) { /* 初始调用为Solve(A, T, 0, N-1, 0) */
	n := ARight - ALeft + 1
	if n == 0 {
		return
	}
	L := GetLeftLength(n) /* 计算出n个结点的树其左子树有多少个结点*/
	T[TRoot] = A[ALeft+L]
	LeftTRoot := TRoot*2 + 1
	RightTRoot := LeftTRoot + 1
	SolveCBT(A, T, ALeft, ALeft+L-1, LeftTRoot)
	SolveCBT(A, T, ALeft+L+1, ARight, RightTRoot)
}

func GetLeftLength(N int) int {
	H := (int)(math.Log2(float64(N + 1)))
	X := N + 1 - int(math.Pow(2.0, (float64)(H)))
	Hminus := int(math.Pow(2.0, (float64)(H-1)))
	X = int(math.Min(float64(X), float64(Hminus)))
	return Hminus - 1 + X
	//return n - int(math.Pow(2.0, (float64)(H-1)))
}
