package main

import "fmt"

//O(n^2) Not recomended iff the list grows than the time complexity incresases as well
func mergeAndSortArrayOn2(arr1 []int, arr2 []int) []int {
	//check if the arr is empty - 1
	//check if the firstArray entry
	arrSorted := append(arr1, arr2...)
	tmpValue := 0
	for i := 0; i <= len(arrSorted)-1; i++ {
		for j := 0; j <= len(arrSorted)-1; j++ {
			if arrSorted[i] < arrSorted[j] {
				tmpValue = arrSorted[i]
				arrSorted[i] = arrSorted[j]
				arrSorted[j] = tmpValue
			}
		}
	}

	return arrSorted
}

func mergeAndSortArrayOn(arr1 []int, arr2 []int) []int {
	mergedArray := []int{}
	arr1Value := arr1[0]
	arr2Value := arr2[0]
	i := 1
	j := 1

	for {
		if arr1Value < arr2Value {
			mergedArray = append(mergedArray, arr1Value)
			arr1Value = arr1[i]
			i++
		}
		if arr2Value < arr1Value {
			mergedArray = append(mergedArray, arr2Value)
			arr2Value = arr2[j]
			j++
		}
		// I still need to figure out how to use this
		if _, ok := arr1Value[i] {
			break	
		}
	}

	return mergedArray

}
func main() {
	fmt.Println("vim-go")
	//	fmt.Println("%t\n", mergeAndSortArrayOn2([]int{1, 2, 34, 5}, []int{6, 7, 8, 900}))
	fmt.Println("%t\n", mergeAndSortArrayOn([]int{1, 2, 34, 5}, []int{6, 7, 8, 900}))

}
