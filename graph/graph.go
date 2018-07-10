package graph

import (
	"fmt"
	"sync"
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
	nodes []*GNode           // 节点集
	edges map[GNode][]*GNode // 邻接表表示的无向图
	lock  sync.RWMutex       // 保证线程安全
}

//AddNode 增加节点
func (g *Graph) AddNode(n *GNode) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.nodes = append(g.nodes, n)
}

//AddEdge 增加边(无向图）
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

//AddEdgeDirection 增加边(有向图）
func (g *Graph) AddEdgeDirection(u, v *GNode) {
	g.lock.Lock()
	defer g.lock.Unlock()
	// 首次建立图
	if g.edges == nil {
		g.edges = make(map[GNode][]*GNode)
	}
	g.edges[*u] = append(g.edges[*u], v) // 建立 u->v 的边
}

//String 输出图
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

//Unweighted 无权单源最短路径
func (g *Graph) Unweighted(S GNode) map[GNode]*GNode {
	g.lock.RLock()
	defer g.lock.RUnlock()

	dist := make(map[GNode]int)
	path := make(map[GNode]*GNode)

	for _, node := range g.nodes {
		if node == nil {
			continue
		}
		dist[*node] = -1
		path[*node] = nil
	}
	var Q GNodeQueue

	Q.New()
	Q.Enqueue(S)
	dist[S] = 0 // source
	for !Q.IsEmpty() {
		V := Q.Dequeue()
		//tranverse connected node of V
		for _, W := range g.edges[*V] {
			if dist[*W] == -1 {
				dist[*W] = dist[*V] + 1
				path[*W] = V
				Q.Enqueue(*W)
			}
		}
	}
	return path

}

//GetPathDijkstra 获得从source到target的路径
func (g *Graph) GetPath(path map[GNode]*GNode, source *GNode, target *GNode) string {
	if path == nil || len(path) == 0 {
		path = g.Unweighted(*source)
	}
	var stack GNodeStack
	stack.New()
	stack.Push(*target)
	for pathPrev := path[*target]; pathPrev != nil; pathPrev = path[*pathPrev] {
		if pathPrev == nil {
			break
		}
		stack.Push(*pathPrev)
	}
	result := ""
	for node := stack.Pop(); node != nil; node = stack.Pop() {
		result += fmt.Sprintf("%d ", node.value)
		if stack.IsEmpty() {
			break
		}
	}
	return result
}
