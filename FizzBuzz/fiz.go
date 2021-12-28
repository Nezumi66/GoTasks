package main

import "fmt"

func fizzBuzz(num int) {
	var str string
	for i := 1; i <= num; i++ {
		if i%3 == 0 {
			str += "Fizz"
		}
		if i%5 == 0 {
			str += "Buzz"
		}
		if str == "" {
			str += fmt.Sprintf("%d", i)
		}
		fmt.Println(str)
		str = ""
	}
}
func main() {
	fizzBuzz(100)
}
