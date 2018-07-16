package subsequence

func MaxSumSubsequenceSum1(A []int) int {
	MaxSum := 0
	NN := len(A)
	for i := 0; i < NN; i++ {
		for j := i; j < NN; j++ {
			ThisSum := 0
			for k := i; k <= j; k++ {
				ThisSum += A[k]
				if ThisSum > MaxSum {
					MaxSum = ThisSum
				}

			}
		}
	}
	return MaxSum
}

func MaxSumSubsequenceSum2(A []int) int {
	MaxSum := 0
	NN := len(A)
	for i := 0; i < NN; i++ {
		ThisSum := 0
		for j := i; j < NN; j++ {
			ThisSum += A[j]
			if ThisSum > MaxSum {
				MaxSum = ThisSum
			}
		}
	}
	return MaxSum
}

/* 返回3个整数中的最大值 */
func Max3(A, B, C int) int {
	if A > B {
		if A > C {
			return A
		} else {
			return C
		}
	} else {
		if B > C {
			return B
		} else {
			return C
		}
	}
}

/* 分治法求List[left]到List[right]的最大子列和 */
func DivideAndConquer(sequence []int, left, right int) int {

	if (left == right) { /* 递归的终止条件，子列只有1个数字 */
		if (sequence[left] > 0) {
			return sequence[left]
		} else {
			return 0
		}
	}

	/* 下面是"分"的过程 */
	center := (left + right) / 2 /* 找到中分点 */
	/* 递归求得两边子列的最大和 */
	MaxLeftSum := DivideAndConquer(sequence, left, center)     //存放左子问题的解
	MaxRightSum := DivideAndConquer(sequence, center+1, right) //存放右子问题的解

	/* 下面求跨分界线的最大子列和 */
	MaxLeftBorderSum := 0 /*存放左边跨分界线的结果*/
	LeftBorderSum := 0
	for i := center; i >= left; i-- {
		/* 从中线向左扫描 */
		LeftBorderSum += sequence[i]
		if LeftBorderSum > MaxLeftBorderSum {
			MaxLeftBorderSum = LeftBorderSum
		}
	} /* 左边扫描结束 */

	MaxRightBorderSum := 0 /*存放右边跨分界线的结果*/
	RightBorderSum := 0
	for i := center + 1; i <= right; i++ { /* 从中线向右扫描 */
		RightBorderSum += sequence[i]
		if RightBorderSum > MaxRightBorderSum {
			MaxRightBorderSum = RightBorderSum
		}
	} /* 右边扫描结束 */

	/* 下面返回"治"的结果 */
	return Max3(MaxLeftSum, MaxRightSum, MaxLeftBorderSum+MaxRightBorderSum)
}

//算法3：分而治之
/* 保持与前2种算法相同的函数接口 */
func MaxSumSubsequenceSum3(A []int) int {
	return DivideAndConquer(A, 0, len(A)-1)
}

//算法4：在线处理
func MaxSumSubsequenceSum4(A []int) int {
	MaxSum := 0
	NN := len(A)
	ThisSum := 0
	for i := 0; i < NN; i++ {
		ThisSum += A[i] //向右累加
		if ThisSum > MaxSum {
			MaxSum = ThisSum // 发现更大和则更新当前结果
		} else if ThisSum < 0 {
			// 如果当前子列和为负数
			ThisSum = 0 // 则不可能使后面部分和增大，抛弃之
		}
	}
	return MaxSum
}
