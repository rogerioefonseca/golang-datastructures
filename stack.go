package main

import "fmt"

type Node struct {
	data string
	next *Node
}

type Stack struct {
	top    *Node
	bottom *Node
	length int
}

func (s *Stack) peek() *Node {
	return s.top
}

func (s *Stack) add(d string) {
	node := Node{data: d}
	s.length++

	if s.length == 0 {
		s.top = &node
		s.bottom = &node
		return
	}

	tmpNode := s.top
	s.top = &node
	s.top.next = tmpNode
}

func (s *Stack) pop() {
	s.top = s.top.next
	s.length--
}

func main() {
	s := Stack{}
	s.add("Rogerio")
	s.add("Fonseca")
	s.add("Leticia")
	s.add("Maria Clara")
	s.add("Maria Luisa")
	s.add("Matthias")

	node := s.top
	length := s.length
	println("ASASASASASASASS")
	for length > 0 {
		fmt.Println(node.data)
		node = node.next
		length--
	}

	println("---------------------------------")
	s.pop()
	s.pop()
	node = s.top
	length = s.length
	for length > 0 {
		fmt.Println(node.data)
		node = node.next
		length--
	}

	println("---------------------------------")
	println(s.peek().data)

}
