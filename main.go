package main

// problem --> https://neetcode.io/problems/palindrome-linked-list/
// проверка на палиндрон односвязного списка

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// =============================== //
// ======== RAM = O(n) ========== //
// =============================== //
//
// func isPalindrome(head *ListNode) bool {
// 	arr := make([]int, 0)
// 	for head != nil {
// 		arr = append(arr, head.Val)
// 		head = head.Next
// 	}

// 	left, right := 0, len(arr)-1

// 	// [1 2 3 2 1]
// 	//    ^   ^

// 	for left <= right {
// 		if arr[left] != arr[right] {
// 			return false
// 		}
// 		left++
// 		right--
// 	}
// 	return true
// }

// =============================== //
// ======== RAM = O(1) ========== //
// =============================== //

func isPalindrome(head *ListNode) bool {

	slow, fast := head, head

	// [1 2 3 2 1]

	// slow будет на середине
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

	}

	curr := slow
	var prev *ListNode = nil

	// начиначем с середины и переворачиваем оставшую часть (до конца)
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	left := head
	right := prev

	// 2-мя указателям смотрим уже
	// [1 2 3] [1 2]
	//  ^       ^
	for right != nil {
		if left.Val != right.Val {
			return false
		}
		left = left.Next
		right = right.Next
	}

	return true
}
