package huffman

import (
	"sort"
)

// Node is a unit of the huffman tree
type Node struct {
	Value  rune
	Weight int
	Left   *Node
	Right  *Node
}

// inOrderTraverse visit the whole substree from this node
func (n Node) inOrderTraverse(code string, visit func(rune, string)) {
	if leftNode := n.Left; leftNode != nil {
		leftNode.inOrderTraverse(code+"0", visit) // left 0
	} else {
		visit(n.Value, code)
		return
	}
	n.Right.inOrderTraverse(code+"1", visit) // right 1
}

//NodeHeap, implements heap.interface, 存储Node的最小堆（按照Node.Weight排序）
type NodeHeap []Node

//Tree: huffman tree
type Tree struct {
	Root *Node
}

// encode inOrderTraverse from the root of the tree and put the encoding result into a map
func (tree Tree) encode() map[rune]string {
	var initialCode string
	encodeMap := make(map[rune]string)
	tree.Root.inOrderTraverse(initialCode, func(value rune, code string) {
		encodeMap[value] = code
	})
	return encodeMap
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
	priorityMap := makePriorityMap(str)
	stortedNodes := makeSortedNodes(priorityMap)
	hfmTree := makeHuffmanTree(stortedNodes)
	return hfmTree.encode()
}

// makePriorityMap make a map[string]int
// key is the distinct character in string, value is the key's times of appration
func makePriorityMap(str string) map[rune]int {
	matchMap := make(map[rune]int)
	for _, chr := range str {
		matchMap[chr] += 1
	}
	return matchMap
}

// makeSortedNodes make a []Node ordered by ascending Weight(递增排序)
func makeSortedNodes(priorityMap map[rune]int) []Node {
	hfmNodes := make(NodeHeap, len(priorityMap))
	i := 0
	for value, weight := range priorityMap {
		hfmNodes[i] = Node{Value: value, Weight: weight}
		i++
	}
	sort.Sort(sort.Reverse(hfmNodes))
	return hfmNodes
}

// makeHuffmanTree make a huffman tree using the sorted node array
func makeHuffmanTree(nodes NodeHeap) *Tree {
	if len(nodes) < 2 {
		panic("Must contain 2 or more elements")
	}
	hfmTree := &Tree{&Node{Weight: nodes[0].Weight + nodes[1].Weight, Left: &nodes[0], Right: &nodes[1]}}
	for i := 2; i < len(nodes); {
		if nodes[i].Weight == 0 {
			i++
			continue
		}
		oldRoot := hfmTree.Root
		if i+1 < len(nodes) && hfmTree.Root.Weight > nodes[i+1].Weight {
			newNode := Node{Weight: nodes[i].Weight + nodes[i+1].Weight, Left: &nodes[i], Right: &nodes[i+1]}
			hfmTree.Root = &Node{Weight: newNode.Weight + oldRoot.Weight, Left: oldRoot, Right: &newNode}
			i += 2
		} else {
			hfmTree.Root = &Node{Weight: nodes[i].Weight + oldRoot.Weight, Left: oldRoot, Right: &nodes[i]}
			i++
		}
	}
	return hfmTree
}
