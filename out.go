package main

import "fmt"

type Node struct {
	Value int   // значение
	Next  *Node // указатель на следующую Node
}

func NewNode(value int) *Node { // конструктор
	return &Node{Value: value}
}

type SingleLinkedList struct {
	Head  *Node // указатель на начало списка
	count int   // кол-во узлов в списке
}

// func NewSingleLinkedList(head *Node) *SingleLinkedList { // конструктор
// 	return &SingleLinkedList{Head: head}
// }

func NewSingleLinkedList(nodes ...*Node) *SingleLinkedList { // конструктор
	list := &SingleLinkedList{}
	if len(nodes) == 0 {
		return list
	}

	list.Head = nodes[0]
	list.count = len(nodes)

	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}

	// последний указывает на nil
	nodes[len(nodes)-1].Next = nil
	return list
}

func (s *SingleLinkedList) GetCount() int {
	return s.count
}

func (s *SingleLinkedList) Print() string {
	var res string
	curr := s.Head

	for range curr.Value {
		if curr != nil {
			res += fmt.Sprintf("%d ", curr.Value)
			curr = curr.Next
		}
	}

	return res
}

// поиск по ключу
func (s *SingleLinkedList) Find(key int) *Node {
	if s == nil || s.count == 0 {
		return nil
	}

	curr := s.Head
	for curr.Value != key { // пока value != key
		curr = curr.Next // двигаемся вперед
		if curr == nil { // если увидели nil -> конец
			return nil
		}
	}
	return curr
}

func (l *SingleLinkedList) FindByIndex(index int) *Node {
	if l == nil || l.count == 0 || index < 0 {
		return nil
	}

	curr := l.Head

	for i := 0; i < l.count; i++ {
		if i == index {
			return curr
		}
		curr = curr.Next
	}

	return nil
}

// FindLast - ищет последний узел равный переданному значению
func (l *SingleLinkedList) FindLast(key int) *Node {

	if l == nil || l.count == 0 {
		return nil
	}

	var lastFound *Node = nil
	curr := l.Head

	for i := 0; i < l.count; i++ {
		if curr.Value == key {
			lastFound = curr
		}

		curr = curr.Next
	}

	return lastFound
}

func (l *SingleLinkedList) AddAfter(node *Node, item int) {
	if node == nil {
		return
	}
	newNode := &Node{Value: item}

	newNode.Next = node.Next
	node.Next = newNode
	l.count++
}

// AddBefore - добавляет узел перед заданным
func (l *SingleLinkedList) AddBefore(node *Node, item int) {
	if node == nil {
		return
	}

	curr := l.Head
	newNode := &Node{Value: item}

	if curr == node {
		newNode.Next = curr
		l.Head = newNode
		l.count++
		return
	}

	for curr != nil {
		if curr.Next == node {
			newNode.Next = node
			curr.Next = newNode
			l.count++
			return
		}
		curr = curr.Next
	}
}

func out() {
	node := NewNode(10)
	node2 := NewNode(55)
	node3 := NewNode(60)
	list := NewSingleLinkedList(node, node2, node3)

	list.AddBefore(node, 88)

	fmt.Println(list.Print())
}
