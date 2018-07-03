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

func createNetwork() IntSet {
	network := Initialization(5)
	network.InputConnection(3, 2)
	network.InputConnection(4, 5)
	network.InputConnection(2, 4)
	return network
}
func TestIntSet_Initialization(t *testing.T) {
	network := Initialization(5)
	counter := network.CheckNetwork(5)
	assert.Equal(t, 5, counter)
	if counter == 1 {
		fmt.Print("The network is connected.\n")
	} else {
		fmt.Printf("There are %d components.\n", counter)
	}

}

func TestIntSet_FindRoot2(t *testing.T) {
	network := createNetwork()
	assert.Equal(t, 2, network.FindRoot2(3))
	assert.Equal(t, 2, network.FindRoot2(4))
	assert.Equal(t, 2, network.FindRoot2(2))
	assert.Equal(t, 2, network.FindRoot2(1))
	assert.Equal(t, 0, network.FindRoot2(0))
}

func TestIntSet_CheckConnection(t *testing.T) {
	network := createNetwork()
	assert.Equal(t, true, network.CheckConnection(2, 3))
	assert.Equal(t, true, network.CheckConnection(4, 3))
	assert.Equal(t, false, network.CheckConnection(1, 3))
}

func TestIntSet_CheckNetwork(t *testing.T) {
	network := createNetwork()
	assert.Equal(t, 2, network.CheckNetwork(5))
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
	var network IntSet
	for i := 0; ; i++ {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if begin {
			n, _ = strconv.Atoi(string(a))
			network = Initialization(n)
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
					}else{
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
