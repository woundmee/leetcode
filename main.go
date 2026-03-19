package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// problem → https://neetcode.io/problems/subtree-of-a-binary-tree/
// задача: проверить, является ли поддерево subTree поддеревом дерева root
// пример:
// Input: root = [1,2,3,4,5,null,null,6], subRoot = [2,4,5]
// Output: false

// root:
//         1
//        / \
//       2   3
//      / \
//     4   5
//    /
//   6

// subRoot:
//     2
//    / \
//   4   5

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	} else if root == nil {
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}

	return isSubtree(root, subRoot)
}

func isSameTree(root, sub *TreeNode) bool {

	if root == nil && sub == nil {
		return true
	} else if root == nil || sub == nil {
		return false
	} else if root.Val != sub.Val {
		return false
	}

	left := isSameTree(root.Left, sub.Left)
	right := isSameTree(root.Right, sub.Right)
	return left && right
}
