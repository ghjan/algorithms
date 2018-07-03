package huffman

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	Node := CodeNode{Weight: 1}
	if Node.Weight != 1 {
		t.Error("Can not instructed")
	}
}

func TestTree(t *testing.T) {
	RootRode := CodeNode{Weight: 1}
	Tree := Tree{&RootRode}
	if Tree.Root == nil {
		t.Error("Can not instrcuted")
	}
}

func TestMakePriorityMap(t *testing.T) {
	str := "1123145512"
	priorityMap := MakeFrequencyMapFromArticle(str)
	if priorityMap['1'] != 4 {
		t.Error("Can not make a right map, priorityMap[1] =", priorityMap['1'])
	}
}

func TestMakeSortedNodes(t *testing.T) {
	str := "112"
	priorityMap := MakeFrequencyMapFromArticle(str)
	sortedNodes := MakeSortedNodes(priorityMap)
	if sortedNodes[0].Value != '2' || sortedNodes[1].Value != '1' {
		t.Error("Can not sort Map, sortedNodes[0] is:", sortedNodes[0])
	}

	str = "555112333444455"
	priorityMap = MakeFrequencyMapFromArticle(str)
	sortedNodes = MakeSortedNodes(priorityMap)
	if sortedNodes[0].Value != '2' || sortedNodes[1].Value != '1' || sortedNodes[4].Value != '5' {
		t.Error("Can not sort Map, sortedNodes[0] is:", sortedNodes[0])
	}

}

func TestMakeHuffManTree(t *testing.T) {
	str := "111223"
	freqMap := MakeFrequencyMapFromArticle(str)
	hfmTree := GenerateHuffmanTreeFromFrequencyMap(freqMap)
	if hfmTree.Root.Weight != 6 {
		t.Error("Can not make a hfmTree, root is:", hfmTree.Root.Left, hfmTree.Root.Right)
	}
}

func TestTraverse(t *testing.T) {
	str := "111223"
	encoding := Encode(str)
	if encoding['1'] != "1" || encoding['2'] != "01" || encoding['3'] != "00" {
		t.Error("Can not InOrderTraverse in pre order, first element is:", encoding)
	}
}

func TestBest(t *testing.T) {
	GOPATH := os.Getenv("GOPATH")
	f := "huffmancode_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var n int
	var freqMap map[rune]int
	var hfmTree *Tree
	countStudentsSubmited := 0
	studentSubmitEncodeMap := make(map[rune]string)
	resultString := ""
	for i := 0; ; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 {
			n, _ = strconv.Atoi(string(a))
		} else if i == 1 {
			freqMap = MakeFrequencyMapFromFreqency(string(a))
			hfmTree = GenerateHuffmanTreeFromFrequencyMap(freqMap)
		} else if i == 2 {
			_, _ = strconv.Atoi(string(a))
		} else {
			words := strings.Split(string(a), " ")
			if len(words) > 1 {
				studentSubmitEncodeMap[rune(words[0][0])] = words[1]
			}

			if (i-2)%n == 0 {
				countStudentsSubmited++
				if result, err := hfmTree.IsBestCode(studentSubmitEncodeMap, freqMap); err != nil {
					//fmt.Println("NO") //fmt.Println(err)
					resultString += "NO "
				} else {
					if result {
						resultString += "YES "
					} else {
						resultString += "NO "
					}
				}
				countStudentsSubmited = 0
				studentSubmitEncodeMap = make(map[rune]string)
			}
		}
		//fmt.Println(string(a))
	}
	fmt.Printf(strings.Replace(resultString, " ", "\n", 4))
	assert.Equal(t, "YES YES NO NO ", resultString)
}
