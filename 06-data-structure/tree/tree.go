package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func CreateEmptyTree() *Tree {
	return &Tree{nil}
}

func Add(tree *Tree, data int) *Tree {
	node := &TreeNode{data, nil, nil}
	if tree.Root == nil {
		tree.Root = node
		return tree
	}

	current := tree.Root
	for {
		if data < current.Val {
			if current.Left != nil {
				current = current.Left
			} else {
				current.Left = node
				return tree
			}
		} else if data > current.Val {
			if current.Right != nil {
				current = current.Right
			} else {
				current.Right = node
				return tree
			}
		}
	}
	return tree
}

func InOrdem(tree *Tree) {
	root := tree.Root

	var inordem func(node *TreeNode)
	inordem = func(node *TreeNode) {
		if node == nil {
			return
		}

		inordem(node.Left)
		fmt.Printf("data: %d\n", node.Val)
		inordem(node.Right)
	}

	inordem(root)
}
