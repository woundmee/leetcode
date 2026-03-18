package main

func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// problem → https://neetcode.io/problems/binary-tree-diameter
// задача: определить диаметр дерева
// диаметр дерева - это определение самой длинное пути между ребер в левом и правом поддереве
// какой путь самый длинный от значения к значению — это и есть диаметр

// Пример:
//         1
//        / \
//       2   3
//      / \    \
//     4   5    6
//    /         \
//   7           8

// Корень не в счёт. Самый длинный путь это от 7 до 5 (или наоборот), там 3 ребра.

var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {

	if root == nil {
		return 0
	}

	// обнуляю, тесты валятся, т.к. состояние переменной сохраняется
	maxDiameter = 0

	depth(root)
	return maxDiameter
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := depth(node.Left)
	right := depth(node.Right)

	maxDiameter = max(maxDiameter, left+right)

	return 1 + max(left, right)
}
