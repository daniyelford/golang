package main

import (
	"fmt"
)

var i int = 1
var n int

func factorial(n int) uint64 {
	if n < 0 {
		fmt.Println("invalid input")
		return 0
	}
	var factval uint64 = 1
	for i := 1; i <= n; i++ {
		factval *= uint64(i)
	}
	return factval
}

func main() {
	fmt.Print("enter a positive integer number: ")
	fmt.Scan(&n)
	fmt.Println(factorial(n))

}
