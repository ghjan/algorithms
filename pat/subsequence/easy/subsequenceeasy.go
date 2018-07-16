package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ghjan/algorithms/pat/subsequence/subsequence"
)

var algoSlice []func([]int) (int, int, int)

func main() {
	fileNames := [...]string{"maxsumsubsequence_case_1.txt"}

	GOPATH := os.Getenv("GOPATH")

	for _, f := range fileNames {
		filename := strings.Join([]string{GOPATH, "bin", f}, "/")
		solveMaxSumSubsequence(filename)
	}
}

func buildSequence(filename string) (int, []int) {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return 0, nil
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var K int //元素个数
	for i := 0; i < 2; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		switch i {
		case 0:
			K, _ = strconv.Atoi(string(a))
			break
		case 1: // 给出
			sequence := strings.Split(string(a), " ")
			var resultSlice []int
			for j := 0; j < len(sequence); j++ {
				num, _ := strconv.Atoi(sequence[j])
				resultSlice = append(resultSlice, num)
			}
			return K, resultSlice
		}
	}
	return 0, nil
}

func solveMaxSumSubsequence(filename string) {
	_, sequence := buildSequence(filename)
	//algoSlice = append(algoSlice, subsequence.MaxSumSubsequenceSum1)
	//algoSlice = append(algoSlice, subsequence.MaxSumSubsequenceSum2)
	//algoSlice = append(algoSlice, subsequence.MaxSumSubsequenceSum3)
	algoSlice = append(algoSlice, subsequence.MaxSumSubsequenceSum4)
	for _, algorithm := range algoSlice {
		sum, _, _ := algorithm(sequence)
		fmt.Printf("%d", sum)
	}
}
