package stack

import (
	"os"
	"bufio"
	"io"
	"strconv"
	"strings"
	"fmt"
)

//CanPopSequence 是否能够pop出target
// M:stack的最大容量；
func (s *ItemStack) CanPopSequence(maxCapacity int, target []int) bool {
	can := false
	//source := make([]int, N, N) //原序列
	//for i := 0; i < N; i++ {
	//	source[i] = i
	//}
	//
	N := len(target)
	indexSource := 0 //原序列的index
	sourceData := -1 //原序列数据 应该是indexSource+1
	for indexTarget := 0; indexTarget < N; indexTarget++ {
		topData := -1 //stack 顶元素 负数表示空栈
		if !s.IsEmpty() {
			topData = (*s.Peek()).(int)
		}
		if indexSource >= 0 && indexSource < N {
			sourceData = indexSource + 1
		}
		targetData := target[indexTarget]
		if topData >= 0 && topData == targetData { //栈顶数据正好满足目标
			s.Pop()
		} else if indexSource >= 0 && indexSource < N && sourceData == targetData { //原数据正好满足目标
			canPush := pushToLimitedStack(s, maxCapacity, sourceData)
			if !canPush {
				return false
			}
			indexSource ++
			s.Pop()
		} else if indexSource >= 0 && indexSource < N && targetData > sourceData {
			for targetData > sourceData {
				canPush := pushToLimitedStack(s, maxCapacity, sourceData)
				if !canPush {
					return false
				}
				indexSource ++
				sourceData = indexSource + 1
			}
			canPush := pushToLimitedStack(s, maxCapacity, sourceData)
			if !canPush {
				return false
			}
			indexSource ++
			s.Pop()
		} else {
			return false
		}
	}
	can = s.IsEmpty() && indexSource == N

	return can
}

func pushToLimitedStack(s *ItemStack, maxCapacity int, item Item) bool {
	if s.Size() >= maxCapacity {
		return false
	}
	s.Push(item)
	return true
}

func SolvePopSequence(filename string) {
	maxCap, N, K, targets := readDataPopSequence(filename)
	for i := 0; i < K; i++ {
		if len(targets[i]) != N {
			fmt.Println("NO")
			continue
		}
		stk := ItemStack{}
		can := stk.CanPopSequence(maxCap, targets[i])
		if can {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
func readDataPopSequence(filename string) (int, int, int, [][]int) {
	fi, err := os.Open(filename)
	if err != nil {
		return -1, -1, -1, nil
	}
	defer fi.Close()

	var targets [][]int
	br := bufio.NewReader(fi)
	var maxCap, N, K int
	for i := 0; ; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { // maxCap, N, K
			if headInfo := strings.Split(string(a), " "); len(headInfo) >= 3 {
				maxCap, _ = strconv.Atoi(headInfo[0])
				N, _ = strconv.Atoi(headInfo[1])
				K, _ = strconv.Atoi(headInfo[2])
				for j := 0; j < K; j++ { // K*N matrix
					row := make([]int, N)
					targets = append(targets, row)
				}
			}
		} else if i <= K { //最多K个target
			if targetRow := strings.Split(string(a), " "); len(targetRow) >= N {
				for jj := 0; jj < len(targetRow); jj++ {
					targets[i-1][jj], _ = strconv.Atoi(targetRow[jj])
				}
			} else {
				fmt.Printf("fail len(targetRow):%d\n", len(targetRow))
			}
		}
	}
	return maxCap, N, K, targets
}
