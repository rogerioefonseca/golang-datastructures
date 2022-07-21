package main

import "fmt"

type array struct {
	value string
	index int
}

func pop(arr []array) {
	lastIndex := len(arr) - 1
	fmt.Println(lastIndex)
}

func main() {
	arr := []array{}
	arr = append(arr, array{value: "1", index: 0})
}
