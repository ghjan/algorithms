package graph

import (
	"fmt"
	"sync"
)

type GNode struct {
	value int
}

type Graph struct {
	nodes []*GNode           // 节点集
	edges map[GNode][]*GNode // 邻接表表示的无向图
	lock  sync.RWMutex       // 保证线程安全
}

// 增加节点
func (g *Graph) AddNode(n *GNode) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.nodes = append(g.nodes, n)
}

// 增加边
func (g *Graph) AddEdge(u, v *GNode) {
	g.lock.Lock()
	defer g.lock.Unlock()
	// 首次建立图
	if g.edges == nil {
		g.edges = make(map[GNode][]*GNode)
	}
	g.edges[*u] = append(g.edges[*u], v) // 建立 u->v 的边
	g.edges[*v] = append(g.edges[*v], u) // 由于是无向图，同时存在 v->u 的边
}

// 输出图
func (g *Graph) String() {
	g.lock.RLock()
	defer g.lock.RUnlock()
	str := ""
	for _, iNode := range g.nodes {
		str += iNode.String() + " -> "
		nexts := g.edges[*iNode]
		for _, next := range nexts {
			str += next.String() + " "
		}
		str += "\n"
	}
	fmt.Println(str)
}

// 输出节点
func (n *GNode) String() string {
	return fmt.Sprintf("%v", n.value)
}

//Unweighted 无权单源最短路径
func (g *Graph) Unweighted(S GNode, dist map[GNode]int, path map[GNode]*GNode, Q GNodeQueue) {
	Q.Enqueue(S)
	for !Q.IsEmpty() {
		V := Q.Dequeue()
		for _, W := range g.edges[*V] {
			if dist[*W] == -1 {
				dist[*W] = dist[*V] + 1
				path[*W] = V
				Q.Enqueue(*W)
			}
		}
	}
}
