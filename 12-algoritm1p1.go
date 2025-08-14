package main

import "fmt"

func main() {

	var number, tempNumber, remainder int

	var result int = 0

	fmt.Println("enter a positive number")
	fmt.Scanln(&number)

	tempNumber = number

	for {
		remainder = tempNumber % 10
		result += remainder * remainder * remainder

		tempNumber /= 10

		if tempNumber == 0 {
			break //breaks the loop
		}

	}

	if result == number {

		fmt.Printf("%d is an armstrong number", number)
	} else {
		fmt.Printf("%d is not an armstrong number", number)
	}

}
