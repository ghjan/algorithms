package binarytree

import (
	"fmt"
	"testing"
)

func TestAvl(t *testing.T) {
	data := []int{3, 2, 1, 4, 5, 6, 7, 16, 15, 14, 13, 12, 11, 10, 8}
	fmt.Println(data)
	root := NewAVLTreeNode(9)
	var err error
	for _, value := range data {
		if root, err = root.Insert(value); err != nil {
			fmt.Println(err)
		} else {
			fmt.Print(" root -> key:", root.key, ", height:", root.height)
			if root.left != nil {
				fmt.Print(", left:", root.left.key)
			}
			if root.right != nil {
				fmt.Print(", right:", root.right.key)
			}
			fmt.Println()
		}
	}

	// fmt.Println(root.DisplayAsc())
	// fmt.Println(root.DisplayDesc())

	root.ProcessDepth(0)
	depth := -1
	root.LevelOrderTraverse(func(avlNode *AvlTreeNode) {
		if avlNode.depth == depth+1 {
			fmt.Println()
			depth++
		}
		fmt.Printf("%d \t", avlNode.key)
	})
}
