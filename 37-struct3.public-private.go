package main

import "fmt"

type car struct {
	Name, Model, Color string
	price              float64
}

func main() {
	c := car{Name: "bmw", Model: "2012", Color: "red", price: 2000}
	c.price = 1000
	fmt.Println("Car: ", c)
}
