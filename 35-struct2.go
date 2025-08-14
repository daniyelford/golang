package main

import "fmt"

// public and private
type Address struct {
	name, street, city string
	pincode            int
}

func main() {
	var a Address
	fmt.Println(a)
	address1 := Address{"reza", "mirdamad", "tehran", 5}
	fmt.Println("address1:", address1)
	address2 := Address{name: "ali", street: "vanak", city: "shiraz", pincode: 10}
	fmt.Println("address2:", address2)
}
