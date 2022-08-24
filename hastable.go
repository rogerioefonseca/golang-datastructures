// Having in count 2 lists, write a code to find the entry that matches in both
// [1,3,4,5] and [5,9,0,70] = true
// [1,3,4,5] and [10,9,0,4] = false

package main

import "fmt"

// That is actually a O(n) time Complexity, and the time will raise a lot depending of the size of the list
// O(a * b) = O(n^2)
func checkIfThereAreMatchItemsInArraysiWithOn2(array1 []int, array2 []int) bool {

	for i := range array1 {
		for j := range array2 {
			if array1[i] == array2[j] {
				return true
			}
		}
	}

	return false
}

type hashNumber struct {
	number int
	exits  bool
}

func checkIfThereAreMatchItemsInArraysiWithOn(array1 []int, array2 []int) bool {
	for i, _ := range array1 {

	}

}

func main() {
	// is this list going to increase?
	// The list contains ony integer values?
	// Are the numbers going to repeat in the list?

	array1 := []int{1, 3, 4, 5}
	array2 := []int{8, 9, 0, 70}
	fmt.Println(checkIfThereAreMatchItemsInArraysiWithOn2(array1, array2))
	fmt.Println(checkIfThereAreMatchItemsInArraysiWithOn(array1, array2))
}
