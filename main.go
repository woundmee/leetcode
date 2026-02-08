package main

func main() {
}

// problem --> https://leetcode.com/problems/remove-duplicates-from-sorted-list

type ListNode struct {
	Val  int
	Next *ListNode
}

// [1,1,2]
func deleteDuplicates(head *ListNode) *ListNode {

	// current/cursor нужен, чтобы не трогать head
	curr := head

	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			// если текущее значение == следующему
			curr.Next = curr.Next.Next // присваиваем текущей следующий за ним
		} else {
			// иначе, если значение не совпадют - идем дальше
			curr = curr.Next
		}
	}

	// возвращаем весь список
	return head
}
