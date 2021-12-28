package main

import (
	"fmt"
)

func palindrome(word string) bool {
	str, check := "", false
	for i := len(word) - 1; i >= 0; i-- {
		str += string(word[i])
	}
	if word == str {
		check = true
	}
	return check
}

func main() {
	fmt.Println(palindrome("madam"))
}
