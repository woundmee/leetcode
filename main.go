package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// problem → https://neetcode.io/problems/invert-a-binary-tree
// задача: инвертировать дерево
// in: [4,2,7,1,3,6,9]    out: [4,7,2,9,6,3,1]

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// root.Left, root.Right = root.Right, root.Left
	// invertTree(root.Left)
	// invertTree(root.Right)

	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}
