package main

import (
	"fmt"
)

// define Node Type :

type Node struct {
	data string
	next *Node
}

// define LinkedList type

type LinkedList struct {
	head *Node
	size int
}

// define LinkedList methods :

func (l *LinkedList) inserHeadNode(target *Node) {
	fmt.Println("Insert New Node :  " + target.data)
	// create new node temp :
	tempNode := l.head
	l.head = target
	l.head.next = tempNode
	l.size++
}

func (l LinkedList) showNodes() {
	// check if the LinkedList is not empty:
	if l.size != 0 && l.head != nil {
		iterator := l.head
		for iterator != nil {
			fmt.Printf(" --> %s ", iterator.data)
			iterator = iterator.next
		}
		fmt.Printf(" --> nill ")
	}
}

func (l LinkedList) isEmpty() bool {
	if l.size == 0 || l.head == nil {
		return true
	}
	return false
}

func (l LinkedList) checkNode(n *Node) bool {
	response := false
	// check if the list is empty
	if !l.isEmpty() {
		iterator := l.head
		for iterator != nil {
			if iterator.data == n.data {
				response = true
			}
			iterator = iterator.next
		}
	}
	return response
}

func (l LinkedList) getNode(index int) (*Node, string) {
	response := l.head
	if l.isEmpty() || index > l.size {
		return nil, "index out of range"
	}
	iterator := l.head
	int_itr := 0
	for iterator != nil {
		if int_itr == index {
			response = iterator
		}
		int_itr++
		iterator = iterator.next
	}
	return response, ""
}

func (l *LinkedList) replaceNode(index int, newNode *Node) {
	if l.isEmpty() || index > l.size || index < 0 {
		return
	}
	iterator := l.head
	itr := 0
	for iterator != nil {
		if itr == index {
			iterator.data = newNode.data
			return
		}
		itr++
		iterator = iterator.next
	}
}

func (l *LinkedList) insertAfter(index int, newNode *Node) {
	if l.isEmpty() || index > l.size || index < 0 {
		return
	}
	iterator := l.head
	itr := 0
	for iterator != nil {
		if itr == index {
			// insert After the current node :
			tempNode := iterator.next
			iterator.next = newNode
			newNode.next = tempNode
			l.size++
			return
		}
		itr++
		iterator = iterator.next
	}
}

func (l *LinkedList) insertBefor(index int, newNode *Node) {
	if l.isEmpty() || index > l.size || index < 0 {
		return
	}
	iterator := l.head
	itr := 0
	for itr != index-1 {
		itr++
		iterator = iterator.next
	}
	// current iterator is the target
	tempNode := iterator.next
	iterator.next = newNode
	newNode.next = tempNode
	l.size++
	return
}

func (l *LinkedList) deleteHader() {
	tempNode := l.head
	l.head = tempNode.next
	l.size--
	return
}
