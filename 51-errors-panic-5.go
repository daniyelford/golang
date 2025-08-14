package main

import "fmt"

func division(num1, num2 int) {
	if num2 == 0 {
		panic("division by zero is not possible")
	} else {
		result := num1 / num2
		fmt.Println("result is ", result)
	}
}

func main() {
	//panic
	division(5, 1)
	division(5, 0)
	division(8, 2) // اجرا نمیشه
}
