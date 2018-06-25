package main

import "fmt"

func main() {
	test1()
	print("----------")
	print(getCurrentDirectory())
	fileName := "E:/go-work/src/github.com/ghjan/algorithms/pat/treetranverseagain/test.txt"
	if err := ReadLine(fileName, Print); err != nil {
		fmt.Printf("err: %s/n", err)
	}
}
