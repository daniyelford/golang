package main

import (
	"fmt"
)

func helloworld() {
	fmt.Println("hello world")
}
func name(name string, age int) {
	fmt.Println(name, age)
}
func add1(a int, b int) int {
	return a + b
}
func add2(a, b int) int {
	return a + b
}

// اسم مقدار خروجی رو result گذاشتی — به این می‌گن named return value یا خروجی نامگذاری شده.
func add3(x int, y int) (result int) {
	result = x + y
	return
	// return result
}
func main() {
	helloworld()
	name("daniyal", 29)
	fmt.Println(add1(5, 6))
	fmt.Println(add2(4, 7))
	fmt.Println(add3(3, 8))
}
