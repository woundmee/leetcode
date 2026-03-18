package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// problem → https://neetcode.io/problems/same-binary-tree/
// проверка эквивалентности левого поддерева правому.

// Пример:
//  a)        b)
//     1         1
//    / \       / \
//   2   3     2   3

// деревья a и b эквивалентны.

func isSameTree(p *TreeNode, q *TreeNode) bool {

	if p == nil && q == nil {
		return true
	} else if (p == nil || q == nil) || (p.Val != q.Val) {
		return false
	}

	left := isSameTree(p.Left, q.Left)
	right := isSameTree(p.Right, q.Right)

	return left && right
}
