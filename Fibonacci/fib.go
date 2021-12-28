package main

import "fmt"

func fibonacciRecur(index int) int {
	if index == 0 || index == 1 {
		return index
	} else {
		return fibonacciRecur(index-1) + fibonacciRecur(index-2)
	}
}

func fibonacciClosure() func() int {
	num1, num2 := 0, 1
	return func() int {
		num1, num2 = num2, num1+num2
		return num2 - num1
	}
}

func fibonacci(index int) int {
	x, y := 0, 1
	for i := 0; i < index; i++ {
		x, y = y, x+y
	}
	return x
}

func main() {
	fmt.Println(fibonacciRecur(13))
	fmt.Println(fibonacci(13))

	f := fibonacciClosure()
	index := 13
	for i := 0; i < index; i++ {
		f()
	}
	fmt.Println(f())
}
