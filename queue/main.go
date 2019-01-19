package main

import "fmt"

type Node struct {
	item int
	next *Node
}

type List struct {
	size int
	head *Node
	tail *Node
}

func NewList() *List {
	return &List{}
}

func (l *List) append(node *Node) {
	if node == nil {
		return
	}
	if l.size == 0 {
		l.head = node
		l.tail = node
		l.size++
		return
	}

	l.tail.next = node
	l.tail = node
	l.size++
}

func (l *List) insert(i int, node *Node) bool {
	if i > l.size-1 || l.size == 0 || node == nil {
		return false
	}

	if i == 0 {
		node.next = l.head
		l.head = node
		l.size++
		return true
	}

	preItem := l.head
	for j := 1; j < i; j++ {
		preItem = preItem.next
	}
	node.next = preItem.next
	preItem.next = node
	l.size++

	return true
}

func (l *List) remove(i int) bool {
	if i > l.size-1 {
		return false
	}

	if i == 0 {
		item := l.head.next
		l.head = item
		l.size--
		return true
	}

	preItem := l.head
	for j := 1; j < i; j++ {
		preItem = preItem.next
	}

	preItem.next = preItem.next.next

	if i == l.size-1 {
		l.tail = preItem
	}

	l.size--
	return true
}

func (l *List) get(i int) *Node {
	if i > l.size-1 {
		return nil
	}

	if i == 0 {
		return l.head
	}
	if i == l.size-1 {
		return l.tail
	}

	item := l.head
	for j := 1; j <= i; j++ {
		item = item.next
	}

	return item
}

func main() {
	l := NewList()
	l.append(&Node{1, nil})
	l.append(&Node{2, nil})
	l.append(&Node{3, nil})
	fmt.Println(l.get(0))
	fmt.Println(l.get(1))
	fmt.Println(l.get(2))
	fmt.Println("===")
	l.insert(0, &Node{4, nil})
	fmt.Println(l.get(0))
	fmt.Println(l.get(1))
	fmt.Println(l.get(2))
	fmt.Println(l.get(3))
	fmt.Println("===")
	l.insert(3, &Node{4, nil})
	fmt.Println(l.get(0))
	fmt.Println(l.get(1))
	fmt.Println(l.get(2))
	fmt.Println(l.get(3))
	fmt.Println(l.get(4))
	fmt.Println("===")
	l.remove(3)
	fmt.Println(l.get(0))
	fmt.Println(l.get(1))
	fmt.Println(l.get(2))
	fmt.Println(l.get(3))
	fmt.Println(l.get(4))
	fmt.Println("===")
	l.remove(3)
	fmt.Println(l.get(0))
	fmt.Println(l.get(1))
	fmt.Println(l.get(2))
	fmt.Println(l.get(3))
	fmt.Println(l.get(4))
}
