package main

import "fmt"

func main() {

	condition := true

	if condition {
		fmt.Println("shart dorost ast conditiojn => true")
	} else {
		fmt.Println("shart ghalat ast condition => false")
	}

	mojoody := 100000
	BtcPrice := 102000

	if mojoody > BtcPrice {
		fmt.Println("you can afford btc")
	} else {
		fmt.Println("you cant afford btc")
	}

	var x int = 20

	if x > 20 {
		fmt.Println("x is greater than 20 ")
	} else if x < 20 {
		fmt.Println("x is less than 20")
	} else {
		fmt.Println("x is equal to 20")
	}

	btc_price := 100
	amount := 200

	if btc_price > 100 && amount > 100 {

		fmt.Println("salam")

	} else {
		fmt.Println("you can buy btc")
	}

	z := 150
	w := 250

	if z < w || z < 100 {
		fmt.Println("hello world")
	} else if w == 250 && z == 150 {
		fmt.Println("x and y are equal")
	} else if w == 50 && z > 100 {
		fmt.Println("y is greater than 50")
	}

}
