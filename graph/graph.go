package graph

import (
	"fmt"
	"sync"
	"github.com/ghjan/algorithms/queue"
	stack2 "github.com/ghjan/algorithms/stack"
)

/*
邻接表存储 - 无权图的单源最短路算法  Unweighted
邻接表存储 - Kruskal最小生成树算法 Kruskal
邻接表存储 - 拓扑排序算法 TopSort
 */
type GNode struct {
	value int
}

//Equal GNode gnode==target
func (gnode *GNode) Equal(target *GNode) bool {

	if gnode == nil || target == nil {
		return gnode == nil && target == nil
	} else {
		return gnode.value == target.value
	}
}

// 输出节点
func (n *GNode) String() string {
	return fmt.Sprintf("%v", n.value)
}

//Graph graph struct
type Graph struct {
	nodes []*GNode      // 节点集
	edges map[int][]int // 邻接表表示的无向图
	lock  sync.RWMutex  // 保证线程安全
}

//AddNode 增加节点
func (g *Graph) AddNode(n *GNode) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.nodes = append(g.nodes, n)
}

//AddEdge 增加边(无向图）
func (g *Graph) AddEdge(u, v int, bZeroBased bool) {
	g.lock.Lock()
	defer g.lock.Unlock()
	// 首次建立图
	if g.edges == nil {
		g.edges = make(map[int][]int)
	}
	if bZeroBased {
		g.edges[u] = append(g.edges[u], v) // 建立 u->v 的边
		g.edges[v] = append(g.edges[v], u) // 由于是无向图，同时存在 v->u 的边
	} else {
		g.edges[u-1] = append(g.edges[u-1], v-1) // 建立 u->v 的边
		g.edges[v-1] = append(g.edges[v-1], u-1) // 建立 v->u 的边

	}
}

//AddEdgeDirection 增加边(有向图）
func (g *Graph) AddEdgeDirection(u, v int, bZeroBased bool) {
	g.lock.Lock()
	defer g.lock.Unlock()
	// 首次建立图
	if g.edges == nil {
		g.edges = make(map[int][]int)
	}
	if bZeroBased {
		g.edges[u] = append(g.edges[u], v) // 建立 u->v 的边
	} else {
		g.edges[u-1] = append(g.edges[u-1], v-1) // 建立 u->v 的边
	}
}

//String 输出图
func (g *Graph) String() {
	g.lock.RLock()
	defer g.lock.RUnlock()
	str := ""
	for index, iNode := range g.nodes {
		str += iNode.String() + " -> "
		edges := g.edges[index]
		for _, next := range edges {
			str += g.nodes[next].String() + " "
		}
		str += "\n"
	}
	fmt.Println(str)
}

//Unweighted 无权单源最短路径
func (g *Graph) Unweighted(S int) []int {
	g.lock.RLock()
	defer g.lock.RUnlock()
	N := len(g.nodes)
	dist := make([]int, N, N)
	path := make([]int, N, N)

	for index, node := range g.nodes {
		if node == nil {
			continue
		}
		dist[index] = -1
		path[index] = -1
	}
	var Q queue.ItemQueue

	Q.New()
	Q.Enqueue(S)
	dist[S] = 0 // source
	for !Q.IsEmpty() {
		V := (*Q.Dequeue()).(int)
		//traverse connected node of V
		for _, W := range g.edges[V] {
			if dist[W] == -1 || dist[W] > dist[V]+1 {
				dist[W] = dist[V] + 1
				path[W] = V
				Q.Enqueue(W)
			}
		}
	}
	return path

}

//GetPathDijkstra 获得从source到target的路径
func (g *Graph) GetPath(path []int, source, target int) string {
	if path == nil || len(path) == 0 {
		path = g.Unweighted(source)
	}
	var stack stack2.ItemStack
	stack.New()
	stack.Push(target)
	for pathPrev := path[target]; pathPrev != -1; pathPrev = path[pathPrev] {
		if pathPrev == -1 {
			break
		}
		stack.Push(pathPrev)
	}
	result := ""
	for node := (*stack.Pop()).(int); node != -1; node = (*stack.Pop()).(int) {
		result += fmt.Sprintf("%d ", node)
		if stack.IsEmpty() {
			break
		}
	}
	return result
}

//Floyd 邻接表存储 - 多源最短路算法 Floyd算法
//return: 路径加权距离矩阵, 路径上一结点矩阵
//func (g *Graph) Floyd() ([][]int, [][]int, error) {
//	N := len(g.nodes)
//	var D, path [][]int //路径加权距离矩阵, 路径上一结点矩阵
//	for i := 0; i < N; i++ {
//		var sd, sp []int
//		for j := 0; j < N; j++ {
//			sd = append(sd, g.matrix[i][j])
//			sp = append(sp, -1)
//		}
//		D = append(D, sd)
//		path = append(path, sp)
//	}
//	for k := 0; k < N; k++ {
//		for i := 0; i < N; i++ {
//			for j := 0; j < N; j++ {
//				if D[i][k]+D[k][j] < D[i][j] {
//					D[i][j] = D[i][k] + D[k][j]
//					if i == j && D[i][j] < 0 { /* 若发现负值圈 */
//						return D, path, errors.New("发现负值圈") /* 不能正确解决，返回错误标记 */
//					}
//					path[i][j] = k
//				}
//			}
//		}
//
//	}
//	return D, path, nil
//}
