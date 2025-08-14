package main

import (
	"fmt"
)

func main() {

	var num int = 20

	if num >= 10 {
		fmt.Println("the number is greater than or equal to 10")

		if num > 15 {
			fmt.Println("the number is also r than 15")
		}

		if num > 19 {
			fmt.Println("the number is also r than 19")
			if num > 20 {
				fmt.Println("the number is also r than 20")
			} else {
				fmt.Println("the number is also re than 20")

			}
		}
		fmt.Println("the number is 20")
	} else {
		fmt.Println("num is less than 10")
	}

	fmt.Println("the number is 20")
}
