package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// problem —> https://leetcode.com/problems/remove-duplicates-from-sorted-list/
// удалить все дубликаты в списке
// in: [1 1 2 2 3 3 4]
// out: [1 2 3 4]

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}
