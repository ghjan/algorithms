package cbt

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initData() []int {
	A := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	return A
}

func TestGetLeftLength(t *testing.T) {
	assert.Equal(t, 3, GetLeftLength(6))
	assert.Equal(t, 6, GetLeftLength(10))
}

func TestCompleteBST(t *testing.T) {
	A := initData()
	sort.Ints(A)
	N := len(A)
	T := make([]int, N)
	SolveCBT(A, T, 0, N-1, 0)
	result := ""
	for i := 0; i < N; i++ {
		result += fmt.Sprintf("%d ", T[i])
	}
	fmt.Println(strings.TrimRight(result, " "))
	assert.Equal(t, "6 3 8 1 5 7 9 0 2 4", strings.TrimRight(result, " "))
}
