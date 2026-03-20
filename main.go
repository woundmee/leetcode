package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// problem → https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
// задача: построить из отсортированного списка бинарное дерево

// in: [-10 -3 0 5 9]
// дерево, которое должно получиться:

//		0
//	   / \
//  -10   5
//    \    \
//    -3    9

// Это BST дерево, которое обладает полезными свойствами:
// * левая часть меньше своего корня
// * правая часть больше своего корня

func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) < 1 {
		return nil
	}

	// находим середину
	mid := len(nums) / 2

	// формируем дерево
	node := &TreeNode{Val: nums[mid]}
	node.Left = sortedArrayToBST(nums[:mid])
	node.Right = sortedArrayToBST(nums[mid+1:])
	return node

}
