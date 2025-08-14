package main

import "fmt"

func main() {

	/* var x int = 20

	if x == 1 {
		fmt.Println("hi")
	} else if x == 2 {
		fmt.Println("hi2")

	} else if x == 3 {
		fmt.Println("hi3")

	} else if x == 20 {
		fmt.Println("hi4")
	} else {
		fmt.Println("bye")
	}

	switch x {
	case 1:
		fmt.Println("hi")
	case 2:
		fmt.Println("hi")
	case 3:
		fmt.Println("hi")
	case 20:
		fmt.Println("hi4")
	default:
		fmt.Println("bye")
	}

	fmt.Println("please provide a number between one and seven")
	var number, x int
	fmt.Scanln(&number)

	switch number {

	case 1:
		fmt.Println("today is saturday")
	case 2:
		fmt.Println("today is sunday")
	case 3:
		fmt.Println("today is monday")
	case 4:
		fmt.Println("today is tuesdayy")
	case 5:
		fmt.Println("today is wednesday")
	case 6:
		fmt.Println("today is thursday")
	case 7:
		fmt.Println("today is friday")

	default:
		fmt.Println("input numbewr is not approved")
		fmt.Scanln(&x)
	}*/

	fmt.Println("please provide a number between one and seven")
	var number int
	n, err := fmt.Scanln(&number)
	fmt.Println("first scan result:", n, "error:", err)

	switch number {
	case 1:
		fmt.Println("today is saturday")
	case 2:
		fmt.Println("today is sunday")
	case 3:
		fmt.Println("today is monday")
	case 4:
		fmt.Println("today is tuesday")
	case 5:
		fmt.Println("today is wednesday")
	case 6:
		fmt.Println("today is thursday")
	case 7:
		fmt.Println("today is friday")
	default:
		fmt.Println("input number is not approved")
		// fmt.Print("please enter any number again: ")
		// n2, err2 := fmt.Scanln(&x)
		// fmt.Println("second scan result:", n2, "error:", err2, "x:", x)
	}

	/*
		var op string
		var a, b int
		fmt.Println("enter a and b in a consecutive order : ")
		fmt.Scanln(&a, &b)
		fmt.Println("enter operator: ")
		fmt.Scanln(&op)

		switch op {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		default:
			fmt.Println("invalid input")

		}
	*/
}
