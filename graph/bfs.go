package graph

import (
	"github.com/ghjan/algorithms/queue"
)

// 实现 BFS 遍历
func (g *Graph) BFS(f func(nodeIndex int)) {
	g.lock.RLock()
	defer g.lock.RUnlock()

	N := len(g.nodes)
	// 初始化队列
	q := queue.ItemQueue{}
	// 取图的第一个节点入队列
	head := 0
	q.Enqueue(head)
	// 标识节点是否已经被访问过-已入队列
	visited := make([]bool, N, N)
	visited[head] = true
	// 遍历所有节点直到队列为空
	for {
		if q.IsEmpty() {
			break
		}
		node := (*q.Dequeue()).(int)
		visited[node] = true
		edges := g.edges[node]
		// 将所有未访问过的邻接节点入队列
		for _, next := range edges {
			// 如果节点已被访问过-已入队列
			if visited[next] {
				continue
			}
			q.Enqueue(next)
			visited[next] = true
		}
		// 对每个正在遍历的（出队列）节点执行回调
		if f != nil {
			f(node)
		}
	}
}
