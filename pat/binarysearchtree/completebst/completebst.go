package main

import (
	"fmt"
	"io"
	"sort"
	"strings"

	"bufio"
	"os"
	"strconv"

	"github.com/ghjan/algorithms/binarysearchtree"
)

func initData(fileName string) []int {
	fi, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var N int
	var result []int
	for i := 0; i < 2; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { // n is the total number
			N, _ = strconv.Atoi(string(a))
			result = make([]int, N)
		} else {
			array1 := strings.Split(string(a), " ")
			for i := 0; i < len(array1); i++ {
				result[i], _ = strconv.Atoi(array1[i])
			}
		}
	}
	return result
}

func test1(fileName string) {
	A := initData(fileName)
	sort.Ints(A)
	N := len(A)
	T := make([]int, N)
	binarysearchtree.Solve(A, T, 0, N-1, 0)
	result := ""
	for i := 0; i < N; i++ {
		result += fmt.Sprintf("%d ", T[i])
	}
	fmt.Println(strings.TrimRight(result, " "))
}

func main() {
	f := "completebst_case_1.txt"
	fileName := strings.Join([]string{"E:/go-work/bin", f}, "/")
	test1(fileName)

}
