package stack

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
)

var stack ItemStack

// 初始化栈
func initStack() *ItemStack {
	if stack.items == nil {
		stack = ItemStack{}
		stack.New()
	}
	return &stack
}

func TestPush(t *testing.T) {
	stack := initStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if size := len(stack.items); size != 3 {
		t.Errorf("Wrong stack size, expected 3 and got %d", size)
	}
}

func TestPop(t *testing.T) {
	stack.Pop()
	if size := len(stack.items); size != 2 {
		t.Errorf("Wrong stack size, expected 2 and got %d", size)
	}

	stack.Pop()
	stack.Pop()
	if size := len(stack.items); size != 0 {
		t.Errorf("Wrong stack size, expected 0 and got %d", size)
	}
}

var pre, in []int
var post []int

func TestStack_Operation(t *testing.T) {
	f := "treeranverseagain_case_1.txt"
	fileName := strings.Join([]string{"E:/go-work/bin", f}, "/")
	result := ""
	if N, err := readData(fileName); err == nil {
		solve(0, 0, 0, N)
		for i := 0; i < N; i++ {
			result += fmt.Sprintf("%d ", post[i])
		}
		fmt.Println(strings.TrimRight(result, " "))
	}
}

func readData(filename string) (int, error) {
	fi, err := os.Open(filename)
	if err != nil {
		return -1, err
	}
	defer fi.Close()

	var itemStack ItemStack
	itemStack.New()
	br := bufio.NewReader(fi)
	var N int
	for i := 0; ; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { // n is the total number of keys to be inserted.
			N, _ = strconv.Atoi(string(a))
			// pre = make([]int, N)
			// in = make([]int, N)
			post = make([]int, N)
		} else {
			if stackOperation := strings.Split(string(a), " "); len(stackOperation) >= 1 {
				if stackOperation[0] == "Push" {
					data, _ := strconv.Atoi(stackOperation[1])
					itemStack.Push(data)
					pre = append(pre, data)
				} else if stackOperation[0] == "Pop" {
					dataPop := *itemStack.Pop()
					in = append(in, dataPop.(int))
				}
			}

		}
	}
	return N, nil
}

//solve
/*
第一次调用 solve(0, 0 , 0, n)
 preL  pre的左边第一个index
 inL   in的左边第一个index
 postL post的左边第一个index
 n    数组元素的个数
*/
func solve(preL, inL, postL, n int) {
	if n == 0 {
		return
	}
	if n == 1 {
		post[postL] = pre[preL]
		return
	}
	root := pre[preL]
	post[postL+n-1] = root
	i := 0
	for ; i < n && in[inL+i] != root; i++ {
	}
	//一分为二， 两个数组的大小分别是nL和nR
	nL := i
	nR := n - nL - 1
	solve(preL+1, inL, postL, nL)
	solve(preL+nL+1, inL+nL+1, postL+nL, nR)
}
