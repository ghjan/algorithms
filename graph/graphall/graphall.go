package graphall

import (
	"fmt"

	"github.com/ghjan/algorithms/queue"
	"github.com/ghjan/algorithms/stack"
)

//MaxInt 大整数
const MaxInt = 999999999

//Vertex 顶点类型
type Vertex struct {
	Label     string
	Edges     []*Edge
	isVisited bool
}

//Edge 边类型
type Edge struct {
	FromVertex *Vertex
	ToVertex   *Vertex
	Weight     int
	isUsed     bool
}

//Graph 图类型
type Graph struct {
	Vertices    []*Vertex
	inDegreeMap map[string]int
}

//AddEdge 增加边(无向图）
func (graph *Graph) AddEdge(indexFrom, indexTo, weight int, isUsed bool) error {
	// 首次建立图
	//mg.matrix[u][v] = weight // 建立 u->v 的边
	//mg.matrix[v][u] = weight // 由于是无向图，同时存在 v->u 的边

	fromVertex := graph.Vertices[indexFrom]
	toVertex := graph.Vertices[indexTo]
	fromVertex.Edges = append(fromVertex.Edges, &Edge{fromVertex, toVertex, weight, isUsed})

	return nil
}


//BreadthFirstSearch 广度优先遍历
func (graph *Graph) BreadthFirstSearch(startVertex *Vertex) []*Vertex {
	if graph.Vertices == nil || len(graph.Vertices) == 0 {
		panic("Graph has no vertex.")
	}
	vertexs := []*Vertex{}
	//fmt.Printf("%s ", startVertex.Label)
	vertexs = append(vertexs, startVertex)
	startVertex.isVisited = true
	queue := &queue.ItemQueue{}
	queue.Enqueue(*startVertex)
	for queue.Size() > 0 { // Visit the nearest vertices that haven't been visited.
		vertex := convertToVertex(*queue.Peek())
		for _, edge := range vertex.Edges {
			if !edge.ToVertex.isVisited {
				//fmt.Printf("%s ", edge.ToVertex.Label)
				vertexs = append(vertexs, edge.ToVertex)
				edge.ToVertex.isVisited = true
				queue.Enqueue(*edge.ToVertex)
			}
		}
		queue.Remove()
	}
	return vertexs

}

//DepthFirstSearch 深度优先遍历
func (graph *Graph) DepthFirstSearch(startVertex *Vertex) []*Vertex {
	if graph.Vertices == nil || len(graph.Vertices) == 0 {
		panic("Graph has no vertex.")
	}
	vertexs := []*Vertex{}
	//fmt.Printf("%s ", startVertex.Label)
	vertexs = append(vertexs, startVertex)
	startVertex.isVisited = true
	stack := &stack.ItemStack{}
	stack.Push(*startVertex)
	for stack.Size() > 0 { // Visit the the vertices by edges that hasn't been visited, until the path ends.
		vertex := convertToVertex(*stack.Peek())
		edge := graph.findEdgeWithUnvistedToVertex(vertex)
		if edge != nil && !edge.ToVertex.isVisited {
			//fmt.Printf("%s ", edge.ToVertex.Label)
			vertexs = append(vertexs, edge.ToVertex)
			edge.ToVertex.isVisited = true
			stack.Push(*edge.ToVertex)
		} else {
			stack.Pop()
		}
	}
	return vertexs
}

//PrimMinimumSpanningTree Prim最小生成树算法
func (graph *Graph) PrimMinimumSpanningTree(startVertex *Vertex) {
	treeEdges := []*Edge{}
	startVertex.isVisited = true
	for len(graph.getVisitedVertices()) < len(graph.Vertices) {
		minWeightEdge := getMinWeightEdgeInUnvisitedVertices(graph.getVisitedVertices())
		if minWeightEdge != nil {
			treeEdges = append(treeEdges, minWeightEdge)
		}
		minWeightEdge.ToVertex.isVisited = true
	}
	graph.clearVerticesVisitHistory()
	for _, edge := range treeEdges {
		fmt.Printf("%s->%s(%d)\n", edge.FromVertex.Label, edge.ToVertex.Label, edge.Weight)
	}
}

//getMinWeightEdgeInUnvisitedVertices 获取最小权重的边（未访问过的到达顶点）
func getMinWeightEdgeInUnvisitedVertices(vertices []*Vertex) *Edge {
	var minWeightEdge *Edge
	for _, vertex := range vertices {
		for _, edge := range vertex.Edges {
			if !edge.ToVertex.isVisited {
				if minWeightEdge == nil || minWeightEdge.Weight > edge.Weight {
					minWeightEdge = edge
				}
			}
		}
	}
	return minWeightEdge
}

func (graph *Graph) findEdgeWithUnvistedToVertex(vertex *Vertex) *Edge {
	for _, edge := range vertex.Edges {
		if !edge.ToVertex.isVisited {
			return edge
		}
	}
	return nil
}

//KruskalMinimumSpanningTree Kruskal最小生成树算法
func (graph *Graph) KruskalMinimumSpanningTree() []*Edge {
	treeEdges := []*Edge{}
	resultEdges := []*Edge{}
	for _, vertex := range graph.Vertices {
		for _, edge := range vertex.Edges {
			treeEdges = append(treeEdges, edge)
		}
	}
	treeCount := len(graph.Vertices)
	for treeCount > 1 {
		minWeightUnUsedEdge := getMinWeightUnUsedEdgeInEdges(treeEdges)
		if minWeightUnUsedEdge != nil {
			oppositeEdge := getOppositeEdgeInEdges(treeEdges, minWeightUnUsedEdge) //反向边
			if !graph.hasUsedEdgeBetweenVertices(minWeightUnUsedEdge.FromVertex, minWeightUnUsedEdge.ToVertex) {
				minWeightUnUsedEdge.isUsed = true
				resultEdges = append(resultEdges, minWeightUnUsedEdge)
				if oppositeEdge != nil {
					oppositeEdge.isUsed = true
				}
				treeCount--
			} else { // There's a ring, remove the edge and its opposite edge.
				fmt.Println("There's a ring, remove the edge and its opposite edge")
				treeEdges = removeEdgeInEdges(treeEdges, minWeightUnUsedEdge)
				treeEdges = removeEdgeInEdges(treeEdges, oppositeEdge)
				resultEdges = removeEdgeInEdges(resultEdges, minWeightUnUsedEdge)
				resultEdges = removeEdgeInEdges(resultEdges, oppositeEdge)
			}
		}
	}
	return resultEdges
}

//getMinWeightUnUsedEdgeInEdges 获取最小权重边（未使用过的边）
func getMinWeightUnUsedEdgeInEdges(edges []*Edge) *Edge {
	var minWeightUnUsedEdge *Edge = nil
	for _, edge := range edges {
		if !edge.isUsed {
			if minWeightUnUsedEdge == nil || minWeightUnUsedEdge.Weight > edge.Weight {
				minWeightUnUsedEdge = edge
			}
		}
	}
	return minWeightUnUsedEdge
}

//getOppositeEdgeInEdges 获取反向边
func getOppositeEdgeInEdges(edges []*Edge, edge *Edge) *Edge {
	for _, e := range edges {
		if e.FromVertex == edge.ToVertex && e.ToVertex == edge.FromVertex && e.Weight == edge.Weight {
			return e
		}
	}
	return nil
}

//hasUsedEdgeBetweenVertices 是否在顶点间已经存在使用过的路径
func (graph *Graph) hasUsedEdgeBetweenVertices(v1 *Vertex, v2 *Vertex) bool {
	queue := &queue.ItemQueue{}
	v1.isVisited = true
	for _, edge := range v1.Edges { //所有v1开始的已经使用过的边的到达节点
		if edge.isUsed {
			queue.Enqueue(*edge.ToVertex)
		}
	}
	for queue.Size() > 0 {
		vertex := convertToVertex(*queue.Peek())
		if vertex == v2 {
			return true
		} else {
			for _, e := range vertex.Edges {
				if e.isUsed && !e.ToVertex.isVisited {
					queue.Enqueue(*e.ToVertex)
				}
			}
		}
		vertex.isVisited = true
		queue.Remove()
	}
	graph.clearVerticesVisitHistory()
	return false
}

func removeEdgeInEdges(edges []*Edge, e *Edge) []*Edge {
	var es []*Edge
	for _, x := range edges {
		if x != e {
			es = append(es, x)
		}
	}
	return es
}

//DijkstraShortestPath Dijkstra最短路径算法
func (graph *Graph) DijkstraShortestPath(startVertex *Vertex, endVertex *Vertex) {
	distanceMap := make(map[string]int)
	prevVertexMap := make(map[string]*Vertex)
	for _, v := range graph.Vertices {
		distanceMap[v.Label] = MaxInt
		prevVertexMap[v.Label] = nil
	}
	distanceMap[startVertex.Label] = 0
	for len(graph.getVisitedVertices()) < len(graph.Vertices) {
		nearestVertex := graph.getNearestVertex(startVertex, distanceMap)
		if nearestVertex == endVertex { //Reached EndVertex.
			break
		}
		if distanceMap[nearestVertex.Label] == MaxInt { //There's no path between two vertices.
			break
		}
		for _, edge := range nearestVertex.Edges { // Update distance map.
			toVertex := edge.ToVertex
			distance := distanceMap[nearestVertex.Label] + edge.Weight
			if distance < distanceMap[toVertex.Label] {
				distanceMap[toVertex.Label] = distance
				prevVertexMap[toVertex.Label] = nearestVertex
			}
		}
		nearestVertex.isVisited = true
	}
	graph.clearVerticesVisitHistory()
	for label, vertex := range prevVertexMap {
		if vertex == nil { // Filter StartVertex.
			delete(prevVertexMap, label)
		} else { // Filter the vertices that can't reach StartVertex and EndVertex.
			if !canGoToStart(vertex, startVertex, prevVertexMap) {
				delete(prevVertexMap, label)
			}
			if !canGoToEnd(graph.getVertexByLabel(label), endVertex, prevVertexMap) {
				delete(prevVertexMap, label)
			}
		}
	}
	for label, vertex := range prevVertexMap {
		fmt.Printf("%s->%s(%d)\n", vertex.Label, label, getWeightByLabelAndPrevVertex(label, vertex))
	}
}

func (graph *Graph) getNearestVertex(startVertex *Vertex, distanceMap map[string]int) *Vertex {
	distance := -1
	index := -1
	for i, v := range graph.Vertices {
		if !v.isVisited {
			if distance == -1 || distance > distanceMap[v.Label] {
				distance = distanceMap[v.Label]
				index = i
			}
		}
	}
	if index == -1 { // First scanning, return StartVertex.
		return startVertex
	} else {
		return graph.Vertices[index]
	}
}

func canGoToStart(v *Vertex, startV *Vertex, prevVertexMap map[string]*Vertex) bool {
	if v == startV {
		return true
	}
	prevV := prevVertexMap[v.Label]
	for prevV != nil {
		if prevV == startV {
			return true
		} else {
			prevV = prevVertexMap[prevV.Label]
		}
	}
	return false
}

func canGoToEnd(v *Vertex, endV *Vertex, prevVertexMap map[string]*Vertex) bool {
	if v == endV {
		return true
	}
	prevV := prevVertexMap[endV.Label]
	for prevV != nil {
		if prevV == v {
			return true
		} else {
			prevV = prevVertexMap[prevV.Label]
		}
	}
	return false
}

//getVertexByLabel 通过label找到相应的顶点
func (graph *Graph) getVertexByLabel(label string) *Vertex {
	for _, v := range graph.Vertices {
		if v.Label == label {
			return v
		}
	}
	return nil
}

func getWeightByLabelAndPrevVertex(label string, prevVertex *Vertex) int {
	for _, edge := range prevVertex.Edges {
		if edge.ToVertex.Label == label {
			return edge.Weight
		}
	}
	return -1
}

//TopologicalSort 拓扑排序
func (graph *Graph) TopologicalSort() {
	graph.inDegreeMap = make(map[string]int)
	for _, v := range graph.Vertices {
		graph.inDegreeMap[v.Label] = 0
	}
	for _, v := range graph.Vertices {
		for _, e := range v.Edges {
			graph.inDegreeMap[e.ToVertex.Label]++
		}
	}
	for len(graph.getVisitedVertices()) < len(graph.Vertices) {
		topVertices := graph.getZeroInDegreeVertices()
		for _, v := range topVertices { // Visit the zero-in-degree-vertex, and decrease the next vertices' in-degree.
			fmt.Printf("%s ", v.Label)
			v.isVisited = true
			for _, edge := range v.Edges {
				graph.inDegreeMap[edge.ToVertex.Label]--
			}
		}
		fmt.Println()
	}
	graph.clearVerticesVisitHistory()
}

func (graph *Graph) getZeroInDegreeVertices() []*Vertex {
	vertices := []*Vertex{}
	for _, v := range graph.Vertices {
		if graph.inDegreeMap[v.Label] == 0 && !v.isVisited {
			vertices = append(vertices, v)
		}
	}
	return vertices
}

func (graph *Graph) getVisitedVertices() []*Vertex {
	vertices := []*Vertex{}
	for _, vertex := range graph.Vertices {
		if vertex.isVisited {
			vertices = append(vertices, vertex)
		}
	}
	return vertices
}

func (graph *Graph) clearVerticesVisitHistory() {
	for _, v := range graph.Vertices {
		v.isVisited = false
	}
}

//clearEdgesUseHistory 清除边的使用记录
func (graph *Graph) clearEdgesUseHistory() {
	for _, v := range graph.Vertices {
		for _, e := range v.Edges {
			e.isUsed = false
		}
	}
}

func convertToVertex(x interface{}) *Vertex {
	if v, ok := x.(Vertex); ok {
		return &v
	} else {
		panic("Type convertion exception.")
	}
}
