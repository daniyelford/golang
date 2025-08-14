package main

import "fmt"

type employee struct {
	firstname, lastname, Address string
	age, salary                  int
}

func main() {
	emp := &employee{"reza", "amiri", "krj", 21, 100}
	fmt.Println("first name is ", (*emp).firstname)
	fmt.Println(emp.firstname)
}
