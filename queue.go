package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type Queue struct {
	first  *Node
	length int
}

func (q *Queue) add(value interface{}) {
	n := Node{data: value}

	if q.length == 0 {
		q.first = &n
		q.length++
		return
	}
	tmp := q.first
	q.first = &n
	q.first.next = tmp
	q.length++
}

func (q *Queue) remove() {
	q.first = q.first.next
	q.length--
}

func main() {
	q := Queue{}
	q.add("12333")
	q.add(123444)
	q.add(true)
	q.add(1.50)

	node := q.first
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}

	fmt.Println("--------------------------------------")
	q.remove()
	q.remove()
	node = q.first
	for node != nil {
		fmt.Println(node.data)
		node = node.next
	}

}
