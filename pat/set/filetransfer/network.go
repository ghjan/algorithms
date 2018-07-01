package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/ghjan/algorithms/set"
)

func test1(filename string) {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var n int //几个节点（电脑）
	begin := true
	var network set.IntSet
	for i := 0; ; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if begin {
			n, _ = strconv.Atoi(string(a))
			network = set.Initialization(n)
			begin = false
		} else //读取节点数据
		{
			cmds := strings.Split(string(a), " ")
			switch cmds[0] {
			case "S":
				counter := network.CheckNetwork(n)
				if counter == 1 {
					fmt.Print("The network is connected.\n")
				} else {
					fmt.Printf("There are %d components.\n", counter)
				}
				return
			case "C": //CheckConnection
				if len(cmds) >= 3 {
					u, _ := strconv.Atoi(cmds[1])
					v, _ := strconv.Atoi(cmds[2])
					if network.CheckConnection(u, v) {
						fmt.Println("yes")
					} else {
						fmt.Println("no")
					}
				}
				break
			case "I": // InputConnection
				if len(cmds) >= 3 {
					u, _ := strconv.Atoi(cmds[1])
					v, _ := strconv.Atoi(cmds[2])
					network.InputConnection(u, v)
				}
				break
			}
		}
	}
}

func main() {
	filename := strings.Join([]string{"E:/go-work/bin", "filetransfer_case_1.txt"}, "/")
	test1(filename)
}
