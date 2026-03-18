package main

import "math"

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// problem → https://neetcode.io/problems/balanced-binary-tree
// задача: проверить баланс дерева. Если левая и правая часть сбалансирована — вернуть true.
// дерево сбалансировано, если разница вычитания дает 0 или 1, а если > 1, то дисбаланс.

func isBalanced(root *TreeNode) bool {
	return balance(root) != -1
}

func balance(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := balance(node.Left)
	right := balance(node.Right)

	if left == -1 || right == -1 {
		return -1
	}

	if math.Abs(float64(left-right)) > 1 {
		return -1
	}
	return max(left, right) + 1
}
