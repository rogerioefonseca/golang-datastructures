package main

import "fmt"

type LinkedList struct {
	head   *Node
	tail   *Node
	lenght int
}

type Node struct {
	data     string
	next     *Node
	previous *Node
}

func (l *LinkedList) prepend(n *Node) {
	n.next = l.head
	l.head.previous = n
	l.head = n
}

//Append into the end of the linked list
func (l *LinkedList) add(n *Node) {
	l.tail.next = n
	n.previous = l.tail
	l.tail = n
}

func (l *LinkedList) insert(n *Node, index int) {
	actualNode := l.head
	actualIndex := 0
	for actualNode.next != nil {
		if actualIndex == index {
			n.previous = actualNode.previous
			n.next = actualNode
			n.previous.next = n
			return
		}
		actualNode = actualNode.next
		actualIndex++
	}
}

func newLinkedList(node *Node) LinkedList {
	l := LinkedList{head: node}
	l.tail = node
	l.lenght = 1

	return l
}

func (l *LinkedList) remove(nIndex int) {
	actualNode := l.head
	index := 0
	for actualNode != nil {
		if (nIndex - 1) == index {
			actualNode.next = actualNode.next.next
			return
		}

		actualNode = actualNode.next
		index++
	}
}

func printAll(l LinkedList) {
	actualNode := l.head
	var arr []string
	for actualNode != nil {
		arr = append(arr, actualNode.data)
		actualNode = actualNode.next
	}

	fmt.Println(arr)

}

func (l *LinkedList) reverse() {
	actualNode := l.tail
	l.head = actualNode
	actualNode.previous = nil
	for actualNode != nil {
		actualNode.next = actualNode.previous
		actualNode = actualNode.previous
	}
	fmt.Println(l.head)
}

func main() {
	n1 := &Node{data: "111"}
	n2 := &Node{data: "222"}
	n3 := &Node{data: "333"}
	n4 := &Node{data: "444"}
	n5 := &Node{data: "555"}
	n6 := &Node{data: "666"}
	n7 := &Node{data: "AAAA"}
	l := newLinkedList(n1)
	l.prepend(n2)
	l.prepend(n3)
	l.prepend(n4)
	l.prepend(n5)
	l.prepend(n6)
	l.insert(n7, 2)
	printAll(l)
	l.remove(3)
	printAll(l)
	l.reverse()
	printAll(l)
	//	lastNode := l.tail
	//
	//	for lastNode != nil {
	//		//fmt.Println(lastNode)
	//		lastNode = lastNode.previous
	//	}
}
