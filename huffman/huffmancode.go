package huffman

import (
	"sort"
	"strconv"
	"strings"
)

// CodeNode is a unit of the huffman tree
type CodeNode struct {
	Value  rune
	Weight int
	Left   *CodeNode
	Right  *CodeNode
}

func (n CodeNode) WPL(Depth int) int {
	if n.Left == nil && n.Right == nil { //leaf
		return Depth * n.Weight
	} else { //否则T一定有2个孩子
		return n.Left.WPL(Depth+1) + n.Right.WPL(Depth+1)
	}
}

// InOrderTraverse visit the whole substree from this node
func (n CodeNode) InOrderTraverse(code string, visit func(rune, string)) {
	if leftNode := n.Left; leftNode != nil {
		leftNode.InOrderTraverse(code+"0", visit) // left 0
	} else {
		visit(n.Value, code)
		return
	}
	n.Right.InOrderTraverse(code+"1", visit) // right 1
}

//NodeHeap, implements heap.interface, 存储Node的最小堆（按照Node.Weight排序）
type NodeHeap []CodeNode

//Tree: huffman tree
type Tree struct {
	Root *CodeNode
}

func (tree Tree) WPL() int {
	return tree.Root.WPL(0)
}

// encode InOrderTraverse from the root of the tree and put the encoding result into a map
func (tree Tree) encode() map[rune]string {
	var initialCode string
	encodeMap := make(map[rune]string)
	tree.Root.InOrderTraverse(initialCode, func(value rune, code string) {
		encodeMap[value] = code
	})
	return encodeMap
}

//IsBestCode 是否最佳编码
func (tree *Tree) IsBestCode(encodeMap map[rune]string, freqMap map[rune]int) (bool, error) {
	resultTree := make([]Node, len(encodeMap))
	weight := 0
	for k, v := range encodeMap {
		if _, err := resultTree[0].InsertCode(k, v); err != nil {
			return false, err
		} else {
			weight += freqMap[k] * len(v)
		}
	}
	return tree.WPL() == weight, nil
}

// Len implements Len() int in sort.Interface
func (h NodeHeap) Len() int {
	return len(h)
}

// Less implements Less(i, j int) bool in sort.Interface
func (h NodeHeap) Less(i, j int) bool {
	return h[i].Weight > h[j].Weight
}

// Swap implements Swap(i, j int) int in sort.Interface
func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Encode encode a str into a map[rune]string
// Example
//   result := huffmancoding.Encode("111223") // result: map[3:00 2:01 1:1]
func Encode(str string) map[rune]string {
	priorityMap := MakeFrequencyMapFromArticle(str)
	stortedNodes := MakeSortedNodes(priorityMap)
	hfmTree := generateHuffmanTreeFromSortedNodes(stortedNodes)
	return hfmTree.encode()
}

// makePriorityMap make a map[string]int
// key is the distinct character in string, value is the key's times of appration
//输入字符串就是文章全文
func MakeFrequencyMapFromArticle(str string) map[rune]int {
	matchMap := make(map[rune]int)
	for _, chr := range str {

		matchMap[chr] += 1
	}
	return matchMap
}

//MakeFrequencyMapFromFreqency
//输入字符串格式 a 2 b 8 c 5 d 6
func MakeFrequencyMapFromFreqency(str string) map[rune]int {
	strs := strings.Split(str, " ")
	priorityMap := make(map[rune]int)
	for i := 0; i+1 < len(strs); i += 2 {
		priorityMap[rune(strs[i][0])], _ = strconv.Atoi(strs[i+1])
	}
	return priorityMap
}

// MakeSortedNodes make a []CodeNode ordered by ascending Weight(递增排序)
func MakeSortedNodes(priorityMap map[rune]int) []CodeNode {
	hfmNodes := make(NodeHeap, len(priorityMap))
	i := 0
	for value, weight := range priorityMap {
		hfmNodes[i] = CodeNode{Value: value, Weight: weight}
		i++
	}
	sort.Sort(sort.Reverse(hfmNodes))
	return hfmNodes
}

//GenerateHuffmanTreeFromFrequencyMap 从一个map（value是词频）产生一个哈夫曼树
func GenerateHuffmanTreeFromFrequencyMap(priorityMap map[rune]int) *Tree {
	sortedNodes := MakeSortedNodes(priorityMap)
	return generateHuffmanTreeFromSortedNodes(sortedNodes)

}

// generateHuffmanTreeFromSortedNodes make a huffman tree using the sorted node array
func generateHuffmanTreeFromSortedNodes(nodes NodeHeap) *Tree {
	if len(nodes) < 2 {
		panic("Must contain 2 or more elements")
	}
	hfmTree := &Tree{&CodeNode{Weight: nodes[0].Weight + nodes[1].Weight, Left: &nodes[0], Right: &nodes[1]}}
	for i := 2; i < len(nodes); {
		if nodes[i].Weight == 0 {
			i++
			continue
		}
		oldRoot := hfmTree.Root
		if i+1 < len(nodes) && hfmTree.Root.Weight > nodes[i+1].Weight {
			newNode := CodeNode{Weight: nodes[i].Weight + nodes[i+1].Weight, Left: &nodes[i], Right: &nodes[i+1]}
			hfmTree.Root = &CodeNode{Weight: newNode.Weight + oldRoot.Weight, Left: oldRoot, Right: &newNode}
			i += 2
		} else {
			hfmTree.Root = &CodeNode{Weight: nodes[i].Weight + oldRoot.Weight, Left: oldRoot, Right: &nodes[i]}
			i++
		}
	}
	return hfmTree
}
