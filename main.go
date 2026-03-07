package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// problem --> https://neetcode.io/problems/linked-list-cycle-detection/

// Input: head = [1,2,3,4], index = 1
//          ⮤—————⤶
// Output: true

// голова может начать откуда угодно.
// нужно испольовать 2 указателя, 1-ый идет на 1 шаг, 2-ой - на 2шага вперед.
// Проверки:
//    1. проверить 2-ой на next==nil, тогда вернуть false (цикла нет)
//    2. и проверить first==second (проверить 2 указателя равны ли друг другу), если цикл есть - они встретятся.

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}
