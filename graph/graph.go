package graph

import "fmt"

// 有向边
type Edge struct {
	From   int // 入节点
	To     int // 出节点
	Weight int // 权重
}

type Graph struct {
	Vertexes int
	Edges    [][]Edge
}

// 创建具有 n 个节点的有权有向图
func NewGraph(n int) *Graph {
	return &Graph{
		Vertexes: n,
		Edges:    make([][]Edge, n),
	}
}

// 向有向图中添加边
func (g *Graph) AddEdge(u, v, w int) {
	g.Edges[u] = append(g.Edges[u], Edge{
		From:   u,
		To:     v,
		Weight: w,
	})
}

// 将邻接好的图输出
func (g *Graph) AdjacentEdges() {
	for _, vertexes := range g.Edges {
		for _, v := range vertexes {
			// 输出从 u 到 v 的边
			fmt.Printf("[Eage]:\t%v -> %v(%d)\n", v.From, v.To, v.Weight)
		}
	}
}
