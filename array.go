package main

import "fmt"

func addToBeggining(v int, arr []int) []int {
	previous := arr[0]
	fmt.Println(arr[0])
	arr[0] = v
	for i := 1; i <= len(arr)+1; i++ {
		arr[i] = previous
		previous = arr[i]
	}

	return arr
}

func main() {
	arr := []int{1, 2, 3, 4}
	arr[4] = 300
	fmt.Println(addToBeggining(0, arr))
}
