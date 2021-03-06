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
	Edges     EdgeSlice
	IsVisited bool
}

//Edge 边类型
type Edge struct {
	FromVertex int
	ToVertex   int
	Weight     int
	isUsed     bool
}
type EdgeSlice []*Edge

func (c EdgeSlice) Len() int {
	return len(c)
}
func (c EdgeSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c EdgeSlice) Less(i, j int) bool {
	return c[i].ToVertex < c[j].ToVertex || c[i].Weight < c[j].Weight
}

//Maneuver 机动时间（AOE网络中的关键路径问题用得到）
type Maneuver struct {
	FromVertex int
	ToVertex   int
	time       int
}

//Graph 图类型
type Graph struct {
	Vertices    []*Vertex //顶点数组
	EdgeNum     int       //边数量
	inDegreeMap map[string]int
}

//AddEdge 增加边
func (graph *Graph) AddEdge(indexFrom, indexTo, weight int, isUsed bool) error {
	if indexFrom > len(graph.Vertices)-1 || indexFrom < 0 {
		return errors.New(fmt.Sprintf("indexFrom(%d) is out of range", indexFrom))
	}
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
func (graph *Graph) DepthFirstSearch(startVertex int, operationFunc func(vertex int) bool) []int {
	if graph.Vertices == nil || len(graph.Vertices) == 0 {
		panic("Graph has no vertex.")
	}
	var vertexes []int
	if operationFunc != nil {
		operationFunc(startVertex)
	}
	vertexes = append(vertexes, startVertex)
	graph.Vertices[startVertex].IsVisited = true
	stk := &stack.ItemStack{}
	stk.Push(startVertex)
	for stk.Size() > 0 { // Visit the the vertices by edges that hasn't been visited, until the path ends.
		vertexIndex := (*stk.Peek()).(int)
		vertexObj := graph.Vertices[vertexIndex]
		edge := graph.findEdgeWithUnvistedToVertex(vertexObj)
		if edge != nil && !graph.Vertices[edge.ToVertex].IsVisited {
			if operationFunc != nil {
				stop := operationFunc(edge.ToVertex)
				if stop {
					break
				}
			}
			vertexes = append(vertexes, edge.ToVertex)
			graph.Vertices[edge.ToVertex].IsVisited = true
			stk.Push(edge.ToVertex)
		} else {
			stk.Pop()
		}
	}
	graph.clearVerticesVisitHistory()
	return vertexes
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
func (graph *Graph) DijkstraShortestPath(startVertexIndex int, endVertexIndex int) map[string]int {
	distanceMap := make(map[string]int)
	prevVertexMap := make(map[string]int)
	for _, v := range graph.Vertices {
		distanceMap[v.Label] = MaxInt
		prevVertexMap[v.Label] = -1
	}
	startVertex := graph.Vertices[startVertexIndex]
	distanceMap[startVertex.Label] = 0
	for len(graph.getVisitedVertices()) < len(graph.Vertices) {
		nearestVertexIndex := graph.getNearestVertex(startVertexIndex, distanceMap)
		nearestVertex := graph.Vertices[nearestVertexIndex]
		if nearestVertexIndex == endVertexIndex { //Reached EndVertex.
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
				prevVertexMap[toVertex.Label] = nearestVertexIndex
			}
		}
		nearestVertex.IsVisited = true
	}
	graph.clearVerticesVisitHistory()
	var labelsShouldDelete []string
	for label, vertexIndex := range prevVertexMap {
		if vertexIndex == -1 { // Filter StartVertex.
			labelsShouldDelete = append(labelsShouldDelete, label)
		} else { // Filter the vertices that can't reach StartVertex and EndVertex.
			if !graph.canGoToStart(vertexIndex, startVertexIndex, prevVertexMap) {
				labelsShouldDelete = append(labelsShouldDelete, label)
			}
			if !graph.canGoToEnd(graph.getVertexByLabel(label), endVertexIndex, prevVertexMap) {
				labelsShouldDelete = append(labelsShouldDelete, label)
			}
		}
	}
	for _, label := range labelsShouldDelete {
		delete(prevVertexMap, label)
	}

	return prevVertexMap
}

//DijkstraShortestPath2 Dijkstra最短路径算法
func (graph *Graph) DijkstraShortestPath2(startVertexIndex int, endVertexIndex int) []int {
	N := len(graph.Vertices)
	distanceSlice := make([]int, N, N)
	prevVertexSlice := make([]int, N, N)
	var pathSlice []int
	for index := 0; index < N; index++ {
		distanceSlice[index] = MaxInt
		prevVertexSlice[index] = -1
	}
	distanceSlice[startVertexIndex] = 0

	nearestVertex := graph.Vertices[startVertexIndex]
	for _, edge := range nearestVertex.Edges { // Update distance slice.
		distance := distanceSlice[startVertexIndex] + edge.Weight
		if distance < distanceSlice[edge.ToVertex] {
			distanceSlice[edge.ToVertex] = distance
			prevVertexSlice[edge.ToVertex] = startVertexIndex
		}
	}
	nearestVertex.IsVisited = true
	for len(graph.getVisitedVertices()) < len(graph.Vertices) {
		nearestVertexIndex := graph.getNearestVertex2(startVertexIndex, distanceSlice)
		if nearestVertexIndex == endVertexIndex { //Reached EndVertex.
			break
		}
		if distanceSlice[nearestVertexIndex] == MaxInt || startVertexIndex == nearestVertexIndex { //There's no path between two vertices.
			break
		}
		nearestVertex := graph.Vertices[nearestVertexIndex]
		for _, edge := range nearestVertex.Edges { // Update distance slice.
			distance := distanceSlice[nearestVertexIndex] + edge.Weight
			if distance < distanceSlice[edge.ToVertex] {
				distanceSlice[edge.ToVertex] = distance
				prevVertexSlice[edge.ToVertex] = nearestVertexIndex
			}
		}
		nearestVertex.IsVisited = true
	}
	graph.clearVerticesVisitHistory()
	for now := endVertexIndex; now >= 0; now = prevVertexSlice[now] {
		pathSlice = append(pathSlice, now)
	}
	return pathSlice
}
func (graph *Graph) getNearestVertex(startVertex int, distanceMap map[string]int) int {
	distance := MaxInt
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
		return index
	}
}

func (graph *Graph) getNearestVertex2(startVertex int, distanceSlice []int) int {
	distance := MaxInt
	index := -1
	for i, v := range graph.Vertices {
		if startVertex != i && !v.IsVisited {
			if distance == -1 || distance > distanceSlice[i] {
				distance = distanceSlice[i]
				index = i
			}
		}
	}
	if index == -1 { // First scanning, return StartVertex.
		return startVertex
	} else {
		return index
	}
}
func (graph *Graph) canGoToStart(vIndex int, startVIndex int, prevVertexMap map[string]int) bool {
	if vIndex == startVIndex {
		return true
	}
	v := graph.Vertices[vIndex]
	prevVIndex := prevVertexMap[v.Label]
	for prevVIndex != -1 {
		if prevVIndex == startVIndex {
			return true
		} else {
			prevV := graph.Vertices[prevVIndex]
			prevVIndex = prevVertexMap[prevV.Label]
		}
	}
	return false
}

func (graph *Graph) canGoToStart2(vIndex int, startVIndex int, prevVertexSlice []int) bool {
	if vIndex == startVIndex {
		return true
	}
	prevVIndex := prevVertexSlice[vIndex]
	for prevVIndex != -1 {
		if prevVIndex == startVIndex {
			return true
		} else {
			prevVIndex = prevVertexSlice[prevVIndex]
		}
	}
	return false
}

func (graph *Graph) canGoToEnd(vIndex int, endVIndex int, prevVertexMap map[string]int) bool {
	if vIndex == endVIndex {
		return true
	}
	endV := graph.Vertices[endVIndex]
	prevVIndex := prevVertexMap[endV.Label]
	for prevVIndex != -1 {
		if prevVIndex == vIndex {
			return true
		} else {
			prevV := graph.Vertices[prevVIndex]
			prevVIndex = prevVertexMap[prevV.Label]
		}
	}
	return false
}

func (graph *Graph) canGoToEnd2(vIndex int, endVIndex int, prevVertexSlice []int) bool {
	if vIndex == endVIndex {
		return true
	}
	prevVIndex := prevVertexSlice[endVIndex]
	for prevVIndex != -1 {
		if prevVIndex == vIndex {
			return true
		} else {
			prevVIndex = prevVertexSlice[prevVIndex]
		}
	}
	return false
}

//getVertexByLabel 通过label找到相应的顶点
func (graph *Graph) getVertexByLabel(label string) int {
	for index, v := range graph.Vertices {
		if v.Label == label {
			return index
		}
	}
	return -1
}

func (graph *Graph) getWeightByLabelAndPrevVertex(label string, prevIndex int) int {
	prevVertex := graph.Vertices[prevIndex]
	for _, edge := range prevVertex.Edges {
		if graph.Vertices[edge.ToVertex].Label == label {
			return edge.Weight
		}
	}
	return -1
}

func (graph *Graph) GetWeightByIndexAndPrevIndex(index, prevIndex int) int {
	prevVertex := graph.Vertices[prevIndex]
	for _, edge := range prevVertex.Edges {
		if edge.ToVertex == index {
			return edge.Weight
		}
	}
	return -1
}

//TopologicalSort 拓扑排序 (使用队列优化过的）
//operationFunc 可以对拓扑排序输出的结果进行操作
//result 拓扑排序的结果
//inVertexes 每个结点前驱结点
func (graph *Graph) TopologicalSort(operationFunc func(vertex int, isSectionEnd bool)) ([]int, []string, error) {
	var result []int
	inVertexes := make([]string, len(graph.Vertices), len(graph.Vertices)) //存放每个节点的前驱节点
	graph.inDegreeMap = make(map[string]int)
	countNilVertix := 0
	for _, v := range graph.Vertices {
		if v == nil {
			countNilVertix++
			continue
		}
		graph.inDegreeMap[v.Label] = 0
	}
	for _, v := range graph.Vertices {
		if v == nil {
			continue
		}
		for _, e := range v.Edges {
			graph.inDegreeMap[graph.Vertices[e.ToVertex].Label]++
			inVertexes[e.ToVertex] += strconv.Itoa(e.FromVertex) + " " + strconv.Itoa(e.Weight) + ","
		}
	}
	que := queue.ItemQueue{}
	for i := 0; i < len(graph.Vertices); i++ {
		v := graph.Vertices[i]
		if v != nil && graph.inDegreeMap[v.Label] == 0 {
			que.Enqueue(i)
		}
	}
	count := 0
	for !que.IsEmpty() {
		vertexIndex := (*que.Dequeue()).(int)
		result = append(result, vertexIndex)
		v := graph.Vertices[vertexIndex]
		count++
		if operationFunc != nil {
			operationFunc(vertexIndex, false)
		}
		for _, edge := range v.Edges {
			graph.inDegreeMap[graph.Vertices[edge.ToVertex].Label]--
			if graph.inDegreeMap[graph.Vertices[edge.ToVertex].Label] == 0 {
				que.Enqueue(edge.ToVertex)
			}
		}
	}

	if count != len(graph.Vertices)-countNilVertix {
		return nil, nil, errors.New("图中有回路")
	}
	return result, inVertexes, nil
}

//TopologicalSortConditional 拓扑排序 (出队列的时候取最小值）
//operationFunc 可以对拓扑排序输出的结果进行操作
//result 拓扑排序的结果
//inVertexes 每个结点前驱结点
func (graph *Graph) TopologicalSortConditional(operationFunc func(vertex Vertex, isSectionEnd bool)) ([]Vertex, []string, error) {
	var result []Vertex
	inVertexes := make([]string, len(graph.Vertices), len(graph.Vertices)) //存放每个节点的前驱节点
	graph.inDegreeMap = make(map[string]int)
	countNilVertix := 0
	for _, v := range graph.Vertices {
		if v == nil {
			countNilVertix++
			continue
		}
		graph.inDegreeMap[v.Label] = 0
	}
	for _, v := range graph.Vertices {
		if v == nil {
			continue
		}
		for _, e := range v.Edges {
			graph.inDegreeMap[graph.Vertices[e.ToVertex].Label]++
			inVertexes[e.ToVertex] += strconv.Itoa(e.FromVertex) + " " + strconv.Itoa(e.Weight) + ","
		}
	}
	que := VertexQueue{}
	for i := 0; i < len(graph.Vertices); i++ {
		v := graph.Vertices[i]
		if v != nil && graph.inDegreeMap[v.Label] == 0 {
			que.Enqueue(*v)
		}
	}
	count := 0
	for !que.IsEmpty() {
		que.Sort()
		vertex := *que.Dequeue()
		result = append(result, vertex)
		//v := graph.Vertices[vertex]
		count++
		if operationFunc != nil {
			operationFunc(vertex, false)
		}
		for _, edge := range vertex.Edges {
			toVertex := graph.Vertices[edge.ToVertex]
			graph.inDegreeMap[toVertex.Label]--
			if graph.inDegreeMap[toVertex.Label] == 0 {
				que.Enqueue(*toVertex)
			}
		}
	}

	if count != len(graph.Vertices)-countNilVertix {
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
func (graph *Graph) InitializeVSet() set.UnionFindSet {
	/* 初始化并查集 */
	return set.InitializationUFS(len(graph.Vertices))
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
func (graph Graph) Earliest(operationFunc func(earliest, result, topSort []int, inVertexes []string, sortedString string)) ([]int, []int, error) {
	earliest := make([]int, len(graph.Vertices), len(graph.Vertices))
	var topSort []int
	sortedString := ""
	var operationFuncForTopo func(vertex int, isSectionEnd bool)
	if operationFunc != nil {
		operationFuncForTopo = func(vertexIndex int, isSectionEnd bool) {
			if isSectionEnd {
				sortedString += ";"
				fmt.Println()
			} else {
				vertex := graph.Vertices[vertexIndex]
				sortedString += vertex.Label + " "
				topSort = append(topSort, vertexIndex)
				// fmt.Printf("%s ", vertex.Label)
			}
		}
	}
	if result, inVertexes, err := graph.TopologicalSort(operationFuncForTopo); err == nil {
		for i := 0; i < len(result); i++ {
			vertexIndex := result[i]
			//fmt.Println(strings.TrimRight(inVertexes[vertexIndex], " ,"))
			if preInfoSlice := strings.Split(strings.TrimRight(inVertexes[vertexIndex], " ,"), ","); preInfoSlice != nil {
				for _, temp := range preInfoSlice {
					prevInfo := strings.Split(temp, " ")
					if prevInfo == nil || len(prevInfo) < 2 {
						continue
					}
					prevIndex, _ := strconv.Atoi(prevInfo[0])
					weight, _ := strconv.Atoi(prevInfo[1])
					if earliest[prevIndex]+weight > earliest[result[i]] {
						earliest[vertexIndex] = earliest[prevIndex] + weight
					}
				}
			}
		}
		if operationFunc != nil {
			operationFunc(earliest, result, topSort, inVertexes, sortedString)
		}
		return earliest, topSort, nil
	} else {
		return nil, nil, err
	}
}

func (graph Graph) Latest(earliest, topSort []int) []int {
	latest := make([]int, len(earliest), len(earliest))
	for i := 0; i < len(latest); i++ {
		latest[i] = MaxInt //初始化为无穷大
	}
	latest[len(earliest)-1] = earliest[len(earliest)-1]

	for i := len(earliest) - 2; i >= 0; i-- { //从最后一个开始
		vertexIndex := topSort[i]
		for _, e := range graph.Vertices[vertexIndex].Edges {
			if latest[e.ToVertex]-e.Weight < latest[i] { //取最小值
				latest[i] = latest[e.ToVertex] - e.Weight
			}
		}
	}

	return latest
}

func (graph Graph) ManeuverTime(earliest, latest, topSort []int) []Maneuver {
	var D []Maneuver
	for i := 0; i < len(topSort); i++ {
		vertexIndex := topSort[i]
		for _, e := range graph.Vertices[vertexIndex].Edges {
			D = append(D, Maneuver{vertexIndex, e.ToVertex, latest[e.ToVertex] - earliest[vertexIndex] - e.Weight})
		}
	}
	return D
}

func (graph Graph) CrucialPath(isDebug bool) (int, []Maneuver, error) {
	var crucial []Maneuver
	if earliest, topSort, err := graph.Earliest(func(earliest, result, topSort []int, inVertexes []string, sortedString string) {
		if isDebug {
			fmt.Println("\n-------result of Earliest--")
			fmt.Println("result:", result)
			fmt.Println("topSort:", topSort)
			fmt.Println("sortedString:", sortedString)
			fmt.Println("-------inVertexes of TopologicalSort--", inVertexes)
		}
	}); err == nil {
		latest := graph.Latest(earliest, topSort)
		if isDebug {
			fmt.Println("---earliest----")
			fmt.Println(earliest)
			fmt.Println("---latest----")
			fmt.Println(latest)
		}
		D := graph.ManeuverTime(earliest, latest, topSort)
		for i := 0; i < len(D); i++ {
			maneuver := D[i]
			if maneuver.time == 0 {
				crucial = append(crucial, maneuver)
			}
		}
		return earliest[len(earliest)-1], crucial, nil

	} else {
		return 0, nil, errors.New("Impossible")
	}

}
