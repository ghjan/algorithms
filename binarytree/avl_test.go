package binarytree

import (
	"fmt"
	"testing"
)

func TestAvl(t *testing.T) {
	data := []int{3, 2, 1, 4, 5, 6, 7, 16, 15, 14, 13, 12, 11, 10, 8}
	fmt.Println(data)
	root := NewAVLTreeNode(9)
	for _, value := range data {
		root = Insert(root, value)
		fmt.Print(" root -> key:", root.key, ", high:", root.high)
		if root.left != nil {
			fmt.Print(", left:", root.left.key)
		}
		if root.right != nil {
			fmt.Print(", right:", root.right.key)
		}
		fmt.Println()
	}

	fmt.Println(DisplayAsc(root))
	fmt.Println(DisplayDesc(root))
}
