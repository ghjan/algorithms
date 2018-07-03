package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"strconv"
	"strings"

	"github.com/ghjan/algorithms/huffman"
)

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
	var freqMap map[rune]int
	var hfmTree *huffman.Tree
	countStudentsSubmited := 0
	studentSubmitEncodeMap := make(map[rune]string)
	resultString := ""
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 {
			n, _ = strconv.Atoi(string(a))
		} else if i == 1 {
			freqMap = huffman.MakeFrequencyMapFromFreqency(string(a))
			hfmTree = huffman.GenerateHuffmanTreeFromFrequencyMap(freqMap)
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
		i++
	}
	fmt.Printf(strings.Replace(resultString, " ", "\n", 4))
}

func main() {
	GOPATH := os.Getenv("GOPATH")
	f := "huffmancode_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")
	testBest(filename)
}
