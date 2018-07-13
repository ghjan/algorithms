package kruskal

import (
	"errors"
	"fmt"

	"strconv"
	"strings"

	"github.com/ghjan/algorithms/queue"
	"github.com/ghjan/algorithms/set"
	"github.com/ghjan/algorithms/stack"
)

/*
使用并查集和边的最小堆优化过的kruskal算法
*/
//MaxInt 大整数
const MaxInt = 999999999

//Vertex 顶点类型
type Vertex struct {
	Label     string
	Edges     []*Edge
	IsVisited bool
}

//Edge 边类型
type Edge struct {
	FromVertex int
	ToVertex   int
	Weight     int
	isUsed     bool
}

//Graph 图类型
type Graph struct {
	Vertices    []*Vertex //顶点数组
	EdgeNum     int       //边数量
	inDegreeMap map[string]int
}

//AddEdge 增加边
func (graph *Graph) AddEdge(indexFrom, indexTo, weight int, isUsed bool) error {
	fromVertex := graph.Vertices[indexFrom]
	fromVertex.Edges = append(fromVertex.Edges, &Edge{indexFrom, indexTo, weight, isUsed})
	graph.EdgeNum++
	return nil
}

//BreadthFirstSearch 广度优先遍历
func (graph *Graph) BreadthFirstSearch(startVertex int, operationFunc func(vertex int)) []int {
	if graph.Vertices == nil || len(graph.Vertices) == 0 {
		panic("Graph has no vertex.")
	}
	var vertexes []int
	operationFunc(startVertex)
	vertexes = append(vertexes, startVertex)
	graph.Vertices[startVertex].IsVisited = true
	que := &queue.ItemQueue{}
	que.Enqueue(startVertex)
	for que.Size() > 0 { // Visit the nearest vertices that haven't been visited.
		vertexIndex := (*que.Peek()).(int)
		vertexObj := graph.Vertices[vertexIndex]
		for _, edge := range vertexObj.Edges {
			if !graph.Vertices[edge.ToVertex].IsVisited {
				operationFunc(edge.ToVertex)
				vertexes = append(vertexes, edge.ToVertex)
				graph.Vertices[edge.ToVertex].IsVisited = true
				que.Enqueue(edge.ToVertex)
			}
		}
		que.Remove()
	}
	graph.clearVerticesVisitHistory()
	return vertexes

}

//DepthFirstSearch 深度优先遍历
func (graph *Graph) DepthFirstSearch(startVertex int, operationFunc func(vertex int)) []int {
	if graph.Vertices == nil || len(graph.Vertices) == 0 {
		panic("Graph has no vertex.")
	}
	var vertexs []int
	operationFunc(startVertex)
	vertexs = append(vertexs, startVertex)
	graph.Vertices[startVertex].IsVisited = true
	stk := &stack.ItemStack{}
	stk.Push(startVertex)
	for stk.Size() > 0 { // Visit the the vertices by edges that hasn't been visited, until the path ends.
		vertexIndex := (*stk.Peek()).(int)
		vertexObj := graph.Vertices[vertexIndex]
		edge := graph.findEdgeWithUnvistedToVertex(vertexObj)
		if edge != nil && !graph.Vertices[edge.ToVertex].IsVisited {
			operationFunc(edge.ToVertex)
			vertexs = append(vertexs, edge.ToVertex)
			graph.Vertices[edge.ToVertex].IsVisited = true
			stk.Push(edge.ToVertex)
		} else {
			stk.Pop()
		}
	}
	graph.clearVerticesVisitHistory()
	return vertexs
}

//PrimMinimumSpanningTree Prim最小生成树算法
//返回：构成最小生成树的那些边
func (graph *Graph) PrimMinimumSpanningTree(startVertex *Vertex) []*Edge {
	var MST []*Edge
	startVertex.IsVisited = true
	for len(graph.getVisitedVertices()) < len(graph.Vertices) {
		minWeightEdge := graph.getMinWeightEdgeInUnvisitedVertices(graph.getVisitedVertices())
		if minWeightEdge != nil {
			MST = append(MST, minWeightEdge)
		}
		graph.Vertices[minWeightEdge.ToVertex].IsVisited = true
	}
	graph.clearVerticesVisitHistory()
	return MST
}

//getMinWeightEdgeInUnvisitedVertices 获取最小权重的边（未访问过的到达顶点）
func (graph Graph) getMinWeightEdgeInUnvisitedVertices(vertices []*Vertex) *Edge {
	var minWeightEdge *Edge
	for _, vertex := range vertices {
		for _, edge := range vertex.Edges {
			if !graph.Vertices[edge.ToVertex].IsVisited {
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
		if !graph.Vertices[edge.ToVertex].IsVisited {
			return edge
		}
	}
	return nil
}

//KruskalMinimumSpanningTree Kruskal最小生成树算法
//返回：构成最小生成树的那些边
func (graph *Graph) KruskalMinimumSpanningTree() []*Edge {
	var treeEdges, MST []*Edge
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
				MST = append(MST, minWeightUnUsedEdge)
				if oppositeEdge != nil {
					oppositeEdge.isUsed = true
				}
				treeCount--
			} else { // There's a ring, remove the edge and its opposite edge.
				fmt.Printf("There's a ring, remove the edge and its opposite edge,%s-%s\n", graph.Vertices[minWeightUnUsedEdge.FromVertex].Label, graph.Vertices[minWeightUnUsedEdge.ToVertex].Label)
				treeEdges = removeEdgeInEdges(treeEdges, minWeightUnUsedEdge)
				treeEdges = removeEdgeInEdges(treeEdges, oppositeEdge)
				MST = removeEdgeInEdges(MST, minWeightUnUsedEdge)
				MST = removeEdgeInEdges(MST, oppositeEdge)
			}
		}
	}
	graph.clearEdgesUseHistory()
	return MST
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
func (graph *Graph) hasUsedEdgeBetweenVertices(start, end int) bool {
	que := &queue.ItemQueue{}
	v1 := graph.Vertices[start]
	v1.IsVisited = true
	for _, edge := range v1.Edges { //所有v1开始的已经使用过的边的到达节点
		if edge.isUsed {
			que.Enqueue(edge.ToVertex) //仅仅对字符串进行队列操作
		}
	}
	for que.Size() > 0 {
		vertexIndex := (*que.Peek()).(int)
		realVertex := graph.Vertices[vertexIndex]
		if vertexIndex == end {
			return true
		} else {
			for _, e := range realVertex.Edges {
				if e.isUsed && !graph.Vertices[e.ToVertex].IsVisited {
					que.Enqueue(e.ToVertex)
				}
			}
		}
		realVertex.IsVisited = true
		que.Remove()
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
			toVertex := graph.Vertices[edge.ToVertex]
			distance := distanceMap[nearestVertex.Label] + edge.Weight
			if distance < distanceMap[toVertex.Label] {
				distanceMap[toVertex.Label] = distance
				prevVertexMap[toVertex.Label] = nearestVertex
			}
		}
		nearestVertex.IsVisited = true
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
		fmt.Printf("%s->%s(%d)\n", vertex.Label, label, graph.getWeightByLabelAndPrevVertex(label, vertex))
	}
}

func (graph *Graph) getNearestVertex(startVertex *Vertex, distanceMap map[string]int) *Vertex {
	distance := -1
	index := -1
	for i, v := range graph.Vertices {
		if !v.IsVisited {
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

func (graph *Graph) getWeightByLabelAndPrevVertex(label string, prevVertex *Vertex) int {
	for _, edge := range prevVertex.Edges {
		if graph.Vertices[edge.ToVertex].Label == label {
			return edge.Weight
		}
	}
	return -1
}

//TopologicalSort 拓扑排序 (使用队列优化过的）
func (graph *Graph) TopologicalSort(operationFunc func(vertex int, isSectionEnd bool)) ([]int, []string, error) {
	var result []int
	inVertexes := make([]string, len(graph.Vertices), len(graph.Vertices)) //存放每个节点的前驱节点
	graph.inDegreeMap = make(map[string]int)
	for _, v := range graph.Vertices {
		graph.inDegreeMap[v.Label] = 0
	}
	for _, v := range graph.Vertices {
		for _, e := range v.Edges {
			graph.inDegreeMap[graph.Vertices[e.ToVertex].Label]++
			inVertexes[e.ToVertex] += strconv.Itoa(e.FromVertex) + " " + strconv.Itoa(e.Weight) + ","
		}
	}
	que := queue.ItemQueue{}
	for i := 0; i < len(graph.Vertices); i++ {
		v := graph.Vertices[i]
		if graph.inDegreeMap[v.Label] == 0 {
			que.Enqueue(i)
		}
	}
	count := 0
	for !que.IsEmpty() {
		vertexIndex := (*que.Dequeue()).(int)
		result = append(result, vertexIndex)
		v := graph.Vertices[vertexIndex]
		count++
		operationFunc(vertexIndex, false)
		for _, edge := range v.Edges {
			graph.inDegreeMap[graph.Vertices[edge.ToVertex].Label]--
			if graph.inDegreeMap[graph.Vertices[edge.ToVertex].Label] == 0 {
				que.Enqueue(edge.ToVertex)
			}
		}
	}

	if count != len(graph.Vertices) {
		return nil, nil, errors.New("图中有回路")
	}
	return result, inVertexes, nil
}

func (graph *Graph) getZeroInDegreeVertices() []*Vertex {
	var vertices []*Vertex
	for _, v := range graph.Vertices {
		if graph.inDegreeMap[v.Label] == 0 && !v.IsVisited {
			vertices = append(vertices, v)
		}
	}
	return vertices
}

func (graph *Graph) getVisitedVertices() []*Vertex {
	var vertices []*Vertex
	for _, vertex := range graph.Vertices {
		if vertex.IsVisited {
			vertices = append(vertices, vertex)
		}
	}
	return vertices
}

func (graph *Graph) clearVerticesVisitHistory() {
	for _, v := range graph.Vertices {
		v.IsVisited = false
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

/*-------------------- 边的最小堆定义 --------------------*/
//PercDown 向下构造最小堆
//参数: ESet 物理存储
//		p ESet[p]为根的子堆
//		N N个顶点
func PercDown(ESet []*Edge, p, N int) {
	/* 改编代码4.24的PercDown( MaxHeap H, int p )    */
	/* 将N个元素的边数组中以ESet[p]为根的子堆调整为关于Weight的最小堆 */
	var Parent, Child int

	X := ESet[p] /* 取出根结点存放的值 */
	for Parent = p; (Parent*2 + 1) < N; Parent = Child {
		Child = Parent*2 + 1
		if (Child != N-1) && (ESet[Child].Weight > ESet[Child+1].Weight) {
			Child++ /* Child指向左右子结点的较小者 */
		}
		if X.Weight <= ESet[Child].Weight {
			/* 找到了合适位置 */
			break
		} else {
			/* 下滤X */
			ESet[Parent] = ESet[Child]
		}
	}
	ESet[Parent] = X
}

//InitializeESet 初始化边的集合
//返回  最小堆（Edge切片）
func (graph *Graph) InitializeESet() []*Edge {
	/* 将图的边存入数组ESet，并且初始化为最小堆 */
	var V int
	//var W *Edge
	var ESet []*Edge
	/* 将图的边存入数组ESet */
	ECount := 0
	for V = 0; V < len(graph.Vertices); V++ {
		for _, W := range graph.Vertices[V].Edges {
			if V < W.ToVertex {
				// 避免重复录入无向图的边，只收V1<V2的边
				ESet = append(ESet, W)
				ECount++
				//if ESet[ECount] != nil {
				//	ESet[ECount].Weight = W.Weight
				//}
			}
		}
	}
	/* 初始化为最小堆 */
	for ECount = len(ESet) / 2; ECount >= 0; ECount-- {
		PercDown(ESet, ECount, len(ESet))
	}
	return ESet
}

//InitializeVSet 初始化顶点并查集
func (graph *Graph) InitializeVSet() set.IntSet {
	/* 初始化并查集 */
	return set.Initialization(len(graph.Vertices))
}

//GetEdge 获取最小堆的最小元素
//返回最小边所在位置
func GetEdge(ESet []*Edge, CurrentSize int) int {
	/* 给定当前堆的大小CurrentSize，将当前最小边位置弹出并调整堆 */

	/* 将最小边与当前堆的最后一个位置的边交换 */
	Swap(ESet, 0, CurrentSize-1)
	/* 将剩下的边继续调整成最小堆 */
	PercDown(ESet, 0, CurrentSize-1)
	return CurrentSize - 1 /* 返回最小边所在位置 */
}

//Swap 交换最小堆的两个元素
func Swap(ESet []*Edge, u, v int) {
	ESet[u], ESet[v] = ESet[v], ESet[u]
}

/*-------------------- 最小堆定义结束 --------------------*/

//Kruskal 使用并查集和边最小堆优化的Kruskal算法
//返回：
//		TotalWeight  最小生成树的总权重
//		MST			 最小生成树里面的那些边
func (graph Graph) Kruskal() (int, []*Edge) {
	if len(graph.Vertices) == 0 {
		return 0, nil
	}
	var MST []*Edge
	TotalWeight := 0               /* 初始化权重和     */
	ECount := 0                    /* 初始化收录的边数 */
	VSet := graph.InitializeVSet() //初始化顶点并查集
	ESet := graph.InitializeESet() //初始化边的最小堆

	NextEdge := len(ESet) /* 原始边集的规模 */

	for ECount < len(graph.Vertices)-1 { /* 当收集的边不足以构成树时 */
		NextEdge = GetEdge(ESet, NextEdge) /* 从边集中得到最小边的位置 */
		if NextEdge < 0 {
			break /* 边集已空 */
		}
		/* 如果该边的加入不构成回路，即两端结点不属于同一连通集 */
		if VSet.CheckCycle(ESet[NextEdge].FromVertex, ESet[NextEdge].ToVertex) == true {
			/* 将该边插入MST */
			MST = append(MST, ESet[NextEdge])
			TotalWeight += ESet[NextEdge].Weight /* 累计权重 */
			ECount++                             /* 生成树中边数加1 */
		}
	}
	if ECount < len(graph.Vertices)-1 {
		TotalWeight = -1 /* 设置错误标记，表示生成树不存在 */
	}
	return TotalWeight, MST
}

//Earliest 整个工期有多长
func (graph Graph) Earliest(operationFunc func(vertex int, isSectionEnd bool)) ([]int, []int, error) {
	earliest := make([]int, len(graph.Vertices), len(graph.Vertices))
	var topSort []int
	sortedString := ""
	if result, inVertexes, err := graph.TopologicalSort(func(vertexIndex int, isSectionEnd bool) {
		if isSectionEnd {
			sortedString += ";"
			fmt.Println()
		} else {
			vertex := graph.Vertices[vertexIndex]
			sortedString += vertex.Label + " "
			topSort = append(topSort, vertexIndex)
			// fmt.Printf("%s ", vertex.Label)
		}
	}); err == nil {
		//fmt.Println("\n-------result of TopologicalSort--")
		//fmt.Println(result)
		//fmt.Println("-------inVertexes of TopologicalSort--")
		for i := 0; i < len(graph.Vertices); i++ {
			//fmt.Println(strings.TrimRight(inVertexes[i], " ,"))
			if preInfoSlice := strings.Split(strings.TrimRight(inVertexes[i], " ,"), ","); preInfoSlice != nil {
				for _, temp := range preInfoSlice {
					prevInfo := strings.Split(temp, " ")
					if prevInfo == nil || len(prevInfo) < 2 {
						continue
					}
					prevIndex, _ := strconv.Atoi(prevInfo[0])
					weight, _ := strconv.Atoi(prevInfo[1])
					if earliest[prevIndex]+weight > earliest[result[i]] {
						earliest[result[i]] = earliest[prevIndex] + weight
					}
				}
			}

		}
		return earliest, topSort, nil
	} else {
		return nil, nil, err
	}
}
