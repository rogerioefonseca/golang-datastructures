package main

import (
	"fmt"
	"strings"
)

func reverseString(str string, separator string) {
	if str == "" {
		return
	}

	arrStr := strings.Split(str, separator)
	str = ""

	//O(n^2) - Quadratic
	for i := (len(arrStr) - 1); i >= 0; i-- {
		if i == 0 {
			separator = ""
		}

		revertedWord := ""
		for j := (len(arrStr[i]) - 1); j >= 0; j-- {
			revertedWord += string(arrStr[i][j])
		}

		str += revertedWord + separator
	}
	fmt.Println(str)
}

// Linear Solution O(n)
func reverseStringLinear(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return string(r)
}

func checkRune(s string) {
	fmt.Printf("%T\n", s)
	fmt.Printf("%s\n", s)

	s_rune := []rune(s)

	for i := range s_rune {
		fmt.Printf("%d - %b - %T\n", s_rune[i], s_rune[i], s_rune[i])
	}
}

func main() {
	// How long can the the string be?
	// does the string has characters that split words?
	// The string contains just string..? yes since it is a string... SURE!
	// Do I need to check the paramters
	// Get the string
	originalString := "My name is Rogerio"
	reverseString(originalString, " ")
	reverseString("Code challenge", " ")
	reverseString("Code challenge 222", "-")
	fmt.Printf("%s\n", reverseStringLinear("Code challenge 222 LINEAR"))
	reverseString("", "")

	c := []rune("ø")
	fmt.Printf("%b %T\n", c, c)

	checkRune("ABCDE")

}
