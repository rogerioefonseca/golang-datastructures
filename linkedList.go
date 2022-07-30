package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	tmpValue := l.head
	l.head = n
	l.head.next = tmpValue
	l.length++
}

func (l linkedList) listAll() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d %d\n", toPrint.data, toPrint.next)
		toPrint = toPrint.next
		l.length--
	}

	fmt.Printf("\n")
}

func main() {
	n1 := &node{data: 100}
	n2 := &node{data: 99}
	n3 := &node{data: 98}
	n4 := &node{data: 97}

	l := linkedList{}
	l.prepend(n1)
	l.prepend(n2)
	l.prepend(n3)
	l.prepend(n4)

	l.listAll()
	fmt.Println(l.length)
}
