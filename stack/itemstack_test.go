package stack

import (
	"os"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)


// 初始化栈
func initStack() *ItemStack {
	var stack ItemStack
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
	stack := initStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	assert.Equal(t, 3, stack.Size())
	stack.Pop()
	assert.Equal(t, 2, stack.Size())

	stack.Pop()
	stack.Pop()
	assert.Equal(t, 0, stack.Size())
}

func TestStack_SolveTreeTranverseAgain(t *testing.T) {
	GOPATH := os.Getenv("GOPATH")
	f := "treeranverseagain_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")

	SolveTreeTranverseAgain(filename)
}

