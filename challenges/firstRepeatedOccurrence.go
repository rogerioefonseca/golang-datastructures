package main

import "fmt"

//O(n) Time complexity
//O(n) Space Complexity
func findFirstOccurrence(arr []int) bool {
	mapInt := map[int]int{}
	for i := 0; i < len(arr)-1; i++ {
		fmt.Println(mapInt[arr[i]])

		if _, ok := mapInt[arr[i]]; ok {
			return false
		}

		mapInt[arr[i]] = arr[i]
	}

	return true
}

func main() {
	arr := []int{1, 2, 1, 4, 5, 2, 5}
	fmt.Printf("%T -  %t", findFirstOccurrence(arr), findFirstOccurrence(arr))
}
