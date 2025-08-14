package main

import "fmt"

func main() {

	var number int
	n1 := 0
	n2 := 1

	nextTerm := 0

	fmt.Println("enter the number of terms: ")
	fmt.Scanln(&number)

	fmt.Print("Fibonacci series is: ")

	for i := 1; i <= number; i++ {

		if i == 1 {

			fmt.Print(" ", n1)
			continue
		}
		if i == 2 {
			fmt.Print(" ", n2)
			continue
		}
		nextTerm = n1 + n2
		n1 = n2
		n2 = nextTerm

		fmt.Print(" ", nextTerm)
	}

}
