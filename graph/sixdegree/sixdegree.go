package sixdegree

import (
	"github.com/ghjan/algorithms/queue"
	"os"
	"fmt"
	"bufio"
	"io"
	"strings"
	"strconv"
)

//const maxn = 10005

//某个顶点对应的六度空间的结果
//vis: 是否访问过的标记
//mp: 保存边权重的矩阵
//n, m: 顶点数，边数
//返回一个百分化的比例整数
func bfsSixDegree(mp [][]int, N, fromIndex int) int {
	// 最后一个结点，尾巴，每个点的最终结果，6层范围内的数，临时节点
	last := fromIndex
	tail := 0
	cnt := 1
	lvl := 0
	vis := make([]int, N, N)

	vis[fromIndex] = 1
	que := queue.ItemQueue{}
	que.Enqueue(fromIndex)
	for !que.IsEmpty() {
		tmp := (*que.Dequeue()).(int)
		for i := 0; i < N; i++ {
			if mp[tmp][i] > 0 && vis[i] == 0 { // 有边且未被访问过
				cnt++
				vis[i] = 1
				tail = i // 为了每一层最后一个的节点标记
				que.Enqueue(i)
			}
		}

		if tmp == last {
			last = tail
			lvl++
		}

		if lvl == 6 {
			break
		}
	}
	return cnt
}

//BuildGraphForBond 构建graph对象和cords分片
func BuildGraphForSixDegree(filename string) (int, int, [][]int) {
	fi, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return 0, 0, nil
	}
	defer fi.Close()

	var mp [][]int
	br := bufio.NewReader(fi)
	i := 0
	var N, M int
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if i == 0 { //顶点数量， 边数量
			array := strings.Split(string(a), " ")
			N, _ = strconv.Atoi(array[0]) //节点数量
			M, _ = strconv.Atoi(array[1]) //边数量
			for i := 0; i < N; i++ {
				lp := make([]int, N, N)
				mp = append(mp, lp)
			}
		} else if i <= M { //边的数据输入 start, end
			array2 := strings.Split(string(a), " ")
			x, _ := strconv.Atoi(array2[0])
			y, _ := strconv.Atoi(array2[1])
			mp[x-1][y-1] = 1
			mp[y-1][x-1] = 1
		} else {
			break
		}
		i++
	}
	return N, M, mp

}

//SolveSixDegree 解决六度空间问题
func SolveSixDegree(N int, mp [][]int) (string) {
	result := ""
	for fromIndex := 0; fromIndex < N; fromIndex++ {
		count := bfsSixDegree(mp, N, fromIndex)
		result += strconv.Itoa(count) + " "
		fmt.Printf("%d: %.2f %% \n", fromIndex, float64(float64(count)*100.0/float64(N)))
	}
	return strings.TrimRight(result, " ")

}
