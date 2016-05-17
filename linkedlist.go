package main

type linkedList struct {
	length int
	first  *linkedNode
	last   *linkedNode
}

type linkedNode struct {
	value    string
	previous *linkedNode
	next     *linkedNode
}

func NewLinkedList() *linkedList {
	return &linkedList{
		length: 0,
		first:  nil,
		last:   nil,
	}
}

func (list *linkedList) Length() int {
	return list.length
}

func (list *linkedList) Add(value string) {
	node := &linkedNode{value: value}
	if list.length == 0 {
		list.first = node
		list.last = node
	} else {
		list.last.next = node
		node.previous = list.last
		list.last = node
	}
	list.length++
}

func (list *linkedList) Remove(index int) string {
	node := list.GetNode(index)
	list.length--
	if node.previous != nil {
		node.previous.next = node.next
	}
	if node.next != nil {
		node.next.previous = node.previous
	}
	return node.value
}

func (list *linkedList) Get(index int) string {
	return list.GetNode(index).value
}

func (list *linkedList) Set(index int, value string) string {
	node := list.GetNode(index)
	originalValue := node.value
	node.value = value
	return originalValue
}

func (list *linkedList) GetNode(index int) *linkedNode {
	if index < 0 || index >= list.Length() {
		panic("Index out of bounds")
	}
	node := list.first
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node
}
