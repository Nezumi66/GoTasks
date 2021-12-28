package main

import "fmt"

func contains(arr []int, elem int) bool {
	for _, value := range arr {
		if value == elem {
			return true
		}
	}
	return false
}

func duplicate(arr []int) {
	var dictionary []int
	for i := 0; i < len(arr); i++ {
		if contains(dictionary, arr[i]) == true {
			continue
		} else {
			found := []int{i}
			for j := i + 1; j < len(arr); j++ {
				if arr[i] == arr[j] {
					found = append(found, j)
				}
			}
			if len(found) > 1 {
				fmt.Printf("Element %d occured at indexes: %v\n", arr[i], found)
				dictionary = append(dictionary, arr[i])
			}
			found = []int{}
		}
	}
}

func main() {
	arr := []int{1, 5, 1, 1, 1, 2, 1, 2, 3, 3, 1, 5}
	duplicate(arr)
}
