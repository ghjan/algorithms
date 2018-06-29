package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/ghjan/algorithms/huffman"
	"strconv"
	"strings"
)

//UnknownTree:  需要判断是否已经最优化编码
type UnknownTree struct {
	Root *huffman.Node
}

// encode inOrderTraverse from the root of the tree and put the encoding result into a map
func (tree UnknownTree) encode() map[rune]string {
	var initialCode string
	encodeMap := make(map[rune]string)
	tree.Root.InOrderTraverse(initialCode, func(value rune, code string) {
		encodeMap[value] = code
	})
	return encodeMap
}

func IsBestCode(hfmTree *huffman.Tree, encodeMap map[rune]string) bool {
	for 

	return false
}

func makePriorityMap(str string) map[rune]int {
	strs := strings.Split(str, " ")
	priorityMap := make(map[rune]int)
	for i := 0; i+1 < len(strs); i += 2 {
		priorityMap[rune(strs[i][0])], _ = strconv.Atoi(strs[i+1])
	}
	return priorityMap
}

func testBest(filename string) {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	i := 0
	var n int
	var hfmTree *huffman.Tree
	countStudentsSubmited := 0
	studentSubmitEncodeMap := make(map[rune]string)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 {
			n, _ = strconv.Atoi(string(a))
		} else if i == 1 {
			hfmTree = huffman.GetHuffmanTree(makePriorityMap(string(a)))
		} else if i == 2 {
			_, _ = strconv.Atoi(string(a))
		} else {
			if (i-2)%n == 0 {
				countStudentsSubmited ++
				if IsBestCode(hfmTree, studentSubmitEncodeMap) {
					fmt.Println("YES")
				} else {
					fmt.Println("NO")
				}
				studentSubmitEncodeMap = make(map[rune]string)
			} else {
				words := strings.Split(string(a), " ")
				if len(words) > 1 {
					studentSubmitEncodeMap[rune(words[0][0])] = words[1]
				}
			}
		}
		//fmt.Println(string(a))
		i++
	}
}

func main() {

}
