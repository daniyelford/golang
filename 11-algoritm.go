package main

import "fmt"

func main() {

	fmt.Println("please provide b & h in a consecutive order ")
	var b, h, area float32
	fmt.Scanln(&b, &h)

	area = b * h / 2
	fmt.Println("s is ", area)

	var num1, num2, temp int
	fmt.Println("enter once number: ")

	fmt.Scanln(&num1)

	fmt.Println("enter twoice number: ")

	fmt.Scanln(&num2)

	temp = num1
	num1 = num2
	num2 = temp

	fmt.Println("enter num1: ")
	fmt.Println("num1 is ", num1)
	fmt.Println("num2 is ", num2)

}
