package main

import "fmt"

func main() {
	// 11 or 22 or 222 or 12321

	var number, remainder, temp int
	var reverse int = 0

	fmt.Println("please provide a positive number:")
	fmt.Scanln(&number)

	temp = number
	for {

		remainder = number % 10
		reverse = reverse*10 + remainder
		number /= 10

		if number == 0 {
			break
		}
	}

	if temp == reverse {

		fmt.Printf("%d is a palindrome number ", temp)
	} else {
		fmt.Printf("%d is not a palindrome number ", temp)
	}

}
