package main

import "fmt"

func handlePanic() {
	x := recover()
	if x != nil {
		fmt.Println("recover", x)
	}
}

func division(num1, num2 int) {
	defer handlePanic()
	if num2 == 0 {
		panic("cant divide number to zero")
	} else {
		result := num1 / num2
		fmt.Println("result: ", result)
	}
}

func main() {
	division(4, 2)
	division(8, 0)
	division(10, 2)
}
