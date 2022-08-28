package main

import (
	"errors"
	"fmt"
)

type Node struct {
	data  int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func (t *Tree) insert(d int) {
	n := Node{data: d}

	if t.root == nil {
		t.root = &n
		return
	}

	currentNode := t.root
	for true {
		if n.data < currentNode.data {
			//left
			if currentNode.left == nil {
				currentNode.left = &n
				return
			}
			currentNode = currentNode.left
		} else {
			//right
			if currentNode.right == nil {
				currentNode.right = &n
				return
			}
			currentNode = currentNode.right
		}
	}
}

func (t *Tree) lookup(data int) (int, error) {
	if t.root == nil {
		return 0, errors.New("Empty BinarySearch Tree")
	}

	currentNode := t.root
	for true {
		if currentNode.data == data {
			return currentNode.data, nil
		}

		if currentNode.left.data < data {
			currentNode = currentNode.left
		} else {
			currentNode = currentNode.right
		}
	}

	return 0, errors.New("Not found")
}

func main() {
	t := Tree{}
	t.insert(1)
	t.insert(2)
	t.insert(3)
	t.insert(4)
	fmt.Println(t.root.right.right)
}
