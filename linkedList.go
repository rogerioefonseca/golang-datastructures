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

func newLinkedList(node *Node) LinkedList {
	l := LinkedList{head: node}
	l.tail = node
	l.lenght = 1

	return l
}

func main() {
	n1 := &Node{data: "111"}
	n2 := &Node{data: "222"}
	n3 := &Node{data: "333"}
	n4 := &Node{data: "444"}
	n5 := &Node{data: "555"}
	n6 := &Node{data: "666"}
	l := newLinkedList(n1)
	l.prepend(n2)
	l.prepend(n3)
	l.prepend(n4)
	l.add(n5)
	l.add(n6)

	actualNode := l.head
	fmt.Println(l.head.next)
	for actualNode.next != nil {
		fmt.Printf("%s\n", actualNode.data)
		actualNode = actualNode.next
	}

	lastNode := l.tail

	for lastNode != nil {
		fmt.Println(lastNode)
		lastNode = lastNode.previous
	}

}
