package main

import "fmt"

var arr [5]string

type linkedList struct {
	head   int
	lenght int
}

type node struct {
	data string
	next *node
}

func (l *linkedList) prepend(d *node) {
	tmpValue := l.head
	l.head = d
	l.head.next = tmpValue
}

func hashFunction(value string) int {
	sum := 0
	for i, _ := range value {
		sum += int(value[i])
	}
	return sum % len(arr)
}

func main() {
	listStrings := []string{"test", "1amigo", "2amigos", "3amigos", "asdasd"}
	for _, v := range listStrings {
		arr[hashFunction(v)] = v
	}

	fmt.Println(arr)
}
