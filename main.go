package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// problem → https://neetcode.io/problems/range-sum-of-bst/
// задача: посчитать сумму узлов, которые входят в заданных диапазон

// in: [5 3 8 1 4 7 9 nil 2], low=3, hight=8
// out: 27
// explanation: узлы 5 3 8 4 и 7 входят в диапазон, а их сумма = 27.

func rangeSumBST(root *TreeNode, low int, high int) int {

	if root == nil {
		return 0
	}

	var sum int

	// проверяю сначал сам корень
	if root.Val >= low && root.Val <= high {
		sum += root.Val
	}

	// проверяю левое поддерево
	if root.Val > low {
		sum += rangeSumBST(root.Left, low, high)
	}

	// проверяю правое поддерево
	if root.Val < high {
		sum += rangeSumBST(root.Right, low, high)
	}

	return sum
}
