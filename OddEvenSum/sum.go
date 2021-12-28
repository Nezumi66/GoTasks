package main

import "fmt"

func sum(num int) {
	odd, even := 0, 0
	for i := 1; i <= num; i++ {
		if i%2 == 0 {
			even += i
		} else {
			odd += i
		}
	}
	fmt.Println(odd, even)
}

func main() {
	sum(10)
}
