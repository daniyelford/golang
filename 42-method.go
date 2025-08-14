package main

import (
	"fmt"
)

/*
func(reciever_name Type)mathodName(parameter list)(return type){
//code
}
*/
type person struct {
	name string
	age  int
}

// defining method with struct reciever
// پیاده‌سازی Interface
func (p person) displayInfo() {
	fmt.Printf("Name: %s, Age: %d \n", p.name, p.age)
}
func main() {
	x := person{name: "John", age: 25}
	x.displayInfo()
}
