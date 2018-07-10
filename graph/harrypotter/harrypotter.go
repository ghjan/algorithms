package harrypotter

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
)

/*
哈利·波特的考试
https://pintia.cn/problem-sets/951072707007700992/problems/985411779057623040
邻接矩阵存储 - 多源最短路算法 Floyd
*/
const (
	//最大整数值
	MaxWeightValue WeightType = WeightType(^uint(0)>>1) / 10000
)

type WeightType int

//HarryGraph 邻接矩阵存储的有权图
type HarryGraph struct {
	vexNum  int            //顶点数量
	edgeNum int            //边数量
	matrix  [][]WeightType //邻接矩阵 保存路径长度(权重）
	lock    sync.RWMutex   // 保证线程安全
}

//CreateMGraph 初始化 边的权重全部设置为MaxIntValue
func (mg *HarryGraph) CreateMGraph(vexNum int) {
	mg.lock.Lock()
	defer mg.lock.Unlock()
	mg.vexNum = vexNum
	for i := 0; i < mg.vexNum; i++ {
		var sl []WeightType
		for j := 0; j < mg.vexNum; j++ {
			if i != j {
				sl = append(sl, MaxWeightValue)
			} else {
				sl = append(sl, 0)
			}
		}
		mg.matrix = append(mg.matrix, sl)
	}
}

//AddEdge 增加边(无向图）
func (mg *HarryGraph) AddEdge(u, v int, weight WeightType) error {
	mg.lock.Lock()
	defer mg.lock.Unlock()
	// 首次建立图
	if len(mg.matrix) == 0 {
		return errors.New("Please use CreateMGraph at first.")
	}
	mg.matrix[u][v] = weight // 建立 u->v 的边
	mg.matrix[v][u] = weight // 由于是无向图，同时存在 v->u 的边
	return nil
}

//Floyd 邻接矩阵存储 - 多源最短路算法 Floyd算法
//return: 路径加权距离矩阵
func (mg *HarryGraph) Floyd() ([][]WeightType, error) {
	mg.lock.RLock()
	defer mg.lock.RUnlock()
	var D [][]WeightType //路径加权距离矩阵, 路径上一结点矩阵
	for i := 0; i < mg.vexNum; i++ {
		var sd []WeightType
		for j := 0; j < mg.vexNum; j++ {
			sd = append(sd, mg.matrix[i][j])
		}
		D = append(D, sd)
	}
	for k := 0; k < mg.vexNum; k++ {
		for i := 0; i < mg.vexNum; i++ {
			for j := 0; j < mg.vexNum; j++ {
				if D[i][k]+D[k][j] < D[i][j] {
					D[i][j] = D[i][k] + D[k][j]
					if i == j && D[i][j] < 0 { /* 若发现负值圈 */
						return D, errors.New("发现负值圈") /* 不能正确解决，返回错误标记 */
					}
					//path[i][j] = k
				}
			}
		}

	}
	return D, nil
}

func buildMGraph(filename string) HarryGraph {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return HarryGraph{}
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	i := 0
	var n, m int
	var mg HarryGraph
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { //顶点数量， 边数量
			array := strings.Split(string(a), " ")
			n, _ = strconv.Atoi(array[0])
			m, _ = strconv.Atoi(array[1])
			mg.CreateMGraph(n)
		} else if i <= m { //边的数据输入 start, end, weight
			array2 := strings.Split(string(a), " ")
			start, _ := strconv.Atoi(array2[0])
			end, _ := strconv.Atoi(array2[1])
			weight, _ := strconv.Atoi(array2[2])
			mg.AddEdge(start-1, end-1, WeightType(weight))
		} else {
			break
		}
		i++
	}
	return mg
}

//findMaxDist 找到第i个顶点到其他顶点的最长距离
func findMaxDist(D [][]WeightType, i, N int) WeightType {
	var MaxDist WeightType
	var j int
	MaxDist = 0
	for j = 0; j < N; j++ { // 找出i到其他动物j的最长距离
		if i != j && D[i][j] > MaxDist {
			MaxDist = D[i][j]
		}
	}
	return MaxDist
}

func findAnimal(mg HarryGraph) (int, WeightType, error) {
	var MaxDist, MinDist WeightType
	var Animal, i int
	if D, err := mg.Floyd(); err == nil {
		MinDist = MaxWeightValue
		for i = 0; i < mg.vexNum; i++ {
			MaxDist = findMaxDist(D, i, mg.vexNum)
			if MaxDist >= MaxWeightValue { /* 说明有从i无法变出的动物*/
				return -1, -1, errors.New(fmt.Sprintf("有从i无法变出的动物,i:%d", i))
			}
			if MinDist > MaxDist { /* 找到最长距离更小的动物*/
				MinDist = MaxDist
				Animal = i + 1 /* 更新距离，记录编号*/
			}
		}
	}
	return Animal, MinDist, nil

}

func harrypotterTest() {
	GOPATH := os.Getenv("GOPATH")
	f := "harrypotter_case_1.txt"
	filename := strings.Join([]string{GOPATH, "bin", f}, "/")
	mg := buildMGraph(filename)
	if Animal, MinDist, err := findAnimal(mg); err == nil {
		fmt.Printf("%d %d\n", Animal, MinDist)
	} else {
		fmt.Println(err)
	}
}

func main() {
	harrypotterTest()

}
