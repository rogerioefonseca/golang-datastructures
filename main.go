package main

import (
	"fmt"
)

var integerArray [4]int

func main() {
	// Array
	integerArray[0] = 1
	integerArray[1] = 2
	integerArray[2] = 3
	integerArray[3] = 4

	fmt.Println(integerArray)

	var stringArray [4]string
	stringArray[0] = "okay"
	stringArray[1] = "okay"
	stringArray[2] = "okay"
	stringArray[3] = "okay"

	fmt.Println(stringArray)

	var boolArray [4]bool
	boolArray[0] = true
	boolArray[1] = true
	boolArray[2] = true
	boolArray[3] = true

	populateStringArray := [4]string{"a", "b", "c"}
	fmt.Println(populateStringArray)

	//Slices

	sliceInt := []int{1, 2, 3, 4, 5}
	sliceInt[0] = 0
	sliceInt[1] = 1
	sliceInt[2] = 2
	sliceInt[3] = 3
	fmt.Println(sliceInt)

	sliceIntWithMake := make([]int, 4)
	sliceIntWithMake[0] = 0
	fmt.Println(sliceIntWithMake)

	test := [4]int{1, 2, 3, 4}
	fmt.Println(test)

	fmt.Printf("##################### Contains duplicate")
	nums := []int{3, 3}
	fmt.Printf("%t", containsDuplicate(nums))

	arrInt := []int{1, 2, 3, 55, 6}
	fmt.Println(matchSum(arrInt, 5))
	arrInt = []int{-1, 1, 3, -5}
	fmt.Println(maximumSubArray(arrInt))
}

func containsDuplicate(nums []int) bool {
	countList := map[int]int{}
	for _, y := range nums {
		countList[y] = countList[y] + 1

	}

	for _, key := range countList {
		if key > 1 {
			return true
		}
	}
	return false
}

func matchSum(nums []int, target int) []int {
	for k, _ := range nums {
		for j, _ := range nums {
			if (nums[j] + nums[k]) == target {
				return []int{j, k}
			}
		}
	}

	return []int{}
}

func maximumSubArray(list []int) int {
	fmt.Println(list)
	//arrValues := []int{}
	var sum int
	//var maxSumTmp int
	for k, _ := range list {
		fmt.Println("ROUND - " + string(k))
		for i := k; i <= (len(list) - 1); i++ {
			fmt.Println(list[i])
			sum += list[i]
			fmt.Println(sum)

			fmt.Println("-----------------")
			//if maxSumTmp > maxSum {
			//	maxSum = maxSumTmp
			//}
		}
		//fmt.Println(maxSum)
	}
	return 1
}
