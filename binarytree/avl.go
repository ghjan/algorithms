package binarytree

import "fmt"

type avlTreeNode struct {
	key   int
	high  int
	left  *avlTreeNode
	right *avlTreeNode
}

func NewAVLTreeNode(value int) *avlTreeNode {
	return &avlTreeNode{key: value}
}

func highTree(p *avlTreeNode) int {
	if p == nil {
		return -1
	} else {
		return p.high
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// look LL
func leftLeftRotation(k *avlTreeNode) *avlTreeNode {
	var kl *avlTreeNode
	kl = k.left
	k.left = kl.right
	kl.right = k
	k.high = max(highTree(k.left), highTree(k.right)) + 1
	kl.high = max(highTree(kl.left), k.high) + 1
	return kl
}

//look RR
func rightRightRotation(k *avlTreeNode) *avlTreeNode {
	var kr *avlTreeNode
	kr = k.right
	k.right = kr.left
	kr.left = k
	k.high = max(highTree(k.left), highTree(k.right)) + 1
	kr.high = max(k.high, highTree(kr.right)) + 1
	return kr
}

func leftRightRotation(k *avlTreeNode) *avlTreeNode {
	k.left = rightRightRotation(k.left)
	return leftLeftRotation(k)
}

func rightLeftRotation(k *avlTreeNode) *avlTreeNode {
	k.right = leftLeftRotation(k.right)
	return rightRightRotation(k)
}

//insert to avl
func Insert(avl *avlTreeNode, key int) *avlTreeNode {
	if avl == nil {
		avl = NewAVLTreeNode(key)
	} else if key < avl.key {
		avl.left = Insert(avl.left, key)
		if highTree(avl.left)-highTree(avl.right) == 2 {
			if key < avl.left.key { //LL
				avl = leftLeftRotation(avl)
			} else { // LR
				avl = leftRightRotation(avl)
			}
		}
	} else if key > avl.key {
		avl.right = Insert(avl.right, key)
		if (highTree(avl.right) - highTree(avl.left)) == 2 {
			if key < avl.right.key { // RL
				avl = rightLeftRotation(avl)
			} else {
				fmt.Println("right right", key)
				avl = rightRightRotation(avl)
			}
		}
	} else if key == avl.key {
		fmt.Println("the key", key, "has existed!")
	}
	//notice: update high(may be this insert no rotation, so you should update high)
	avl.high = max(highTree(avl.left), highTree(avl.right)) + 1
	return avl
}

//display avl tree  key by asc
func DisplayAsc(avl *avlTreeNode) []int {
	return AppendValues([]int{}, avl)
}

func AppendValues(values []int, avl *avlTreeNode) []int {
	if avl != nil {
		values = AppendValues(values, avl.left)
		values = append(values, avl.key)
		values = AppendValues(values, avl.right)
	}
	return values
}

//display avl tree key by desc
func DisplayDesc(avl *avlTreeNode) []int {
	return AppendValues2([]int{}, avl)
}

func AppendValues2(values []int, avl *avlTreeNode) []int {
	if avl != nil {
		values = AppendValues2(values, avl.right)
		values = append(values, avl.key)
		values = AppendValues2(values, avl.left)
	}
	return values
}
