package main

func main() {}

type ListNode struct {
	Val  int
	Next *ListNode
}

// problem —> https://neetcode.io/problems/remove-linked-list-elements/
// удалить все вхождения val

func removeElements(head *ListNode, val int) *ListNode {

	// удаляю все совпадающие головы
	for head != nil && head.Val == val {
		head = head.Next
	}

	// удаляю совпадающие узлы
	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Next.Val == val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next // move
		}
	}
	return head
}
