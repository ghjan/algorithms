package set

import (
	"fmt"
	"io"
	"testing"

	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/stretchr/testify/assert"
)

func createNetwork() UnionFindSet {
	network := InitializationUFS(5)
	network.InputConnection(3, 2)
	network.InputConnection(4, 5)
	network.InputConnection(2, 4)
	return network
}
func TestUnionFindSet_Initialization(t *testing.T) {
	network := InitializationUFS(5)
	counter, _ := network.CheckNetwork()
	assert.Equal(t, 5, counter)
	if counter == 1 {
		fmt.Print("The network is connected.\n")
	} else {
		fmt.Printf("There are %d components.\n", counter)
	}

}

func TestUnionFindSet_FindRoot(t *testing.T) {
	network := createNetwork()
	assert.Equal(t, 2, network.FindRoot(3))
	assert.Equal(t, 2, network.FindRoot(4))
	assert.Equal(t, 2, network.FindRoot(2))
	assert.Equal(t, 2, network.FindRoot(1))
	assert.Equal(t, 0, network.FindRoot(0))
}

func TestUnionFindSet_CheckConnection(t *testing.T) {
	network := createNetwork()
	assert.Equal(t, true, network.CheckConnection(2, 3))
	assert.Equal(t, true, network.CheckConnection(4, 3))
	assert.Equal(t, false, network.CheckConnection(1, 3))
}

func TestUnionFindSet_CheckNetwork(t *testing.T) {
	network := createNetwork()
	count, _ := network.CheckNetwork()
	assert.Equal(t, 2, count)
}

func TestNetworkComponent(t *testing.T) {
	GOPATH := os.Getenv("GOPATH")
	f := "filetransfer_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	var n int //几个节点（电脑）
	begin := true
	var network UnionFindSet
	for i := 0; ; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if begin {
			n, _ = strconv.Atoi(string(a))
			network = InitializationUFS(n)
			begin = false
		} else //读取节点数据
		{
			cmds := strings.Split(string(a), " ")
			switch cmds[0] {
			case "S":
				counter, _ := network.CheckNetwork()
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

func TestUnionFindSet_CheckCycle(t *testing.T) {
	network := createNetwork()
	assert.Equal(t, false, network.CheckCycle(3, 2))
	assert.Equal(t, false, network.CheckCycle(3, 4))
	assert.Equal(t, false, network.CheckCycle(3, 1))
	assert.Equal(t, false, network.FindRoot(3-1) == network.FindRoot(1-1))
	assert.Equal(t, true, network.CheckCycle(2, 0))
	assert.Equal(t, true, network.FindRoot(3-1) == network.FindRoot(1-1))
}
