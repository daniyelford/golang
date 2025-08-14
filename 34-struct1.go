package main

import "fmt"

type Person struct {
	name   string
	age    int
	salary int
	job    string
}

func main() {
	var pers1 Person
	pers1.name = "reza"
	pers1.age = 25
	pers1.salary = 1000
	pers1.job = "mechanic"
	fmt.Println(pers1)
	fmt.Println("fisrt person Name: ", pers1.name)
	p := Person{"Ali", 25, 500, "developer"}
	fmt.Println("fisrt person age: ", p.age)
	p1 := Person{name: "mmd", age: 27, salary: 4000, job: "post"}
	fmt.Println("fisrt person salary: ", p1.salary)
	fmt.Println(".........................................")
	printPersonData(pers1)
	// p1.Greet()
}
func printPersonData(pers Person) {
	fmt.Println(" person Name: ", pers.name)
	fmt.Println(" person age: ", pers.age)
	fmt.Println(" person salary: ", pers.salary)
	fmt.Println(" person job: ", pers.job)
}

// func (p Person) Greet() {
// 	fmt.Printf("Hi, I'm %s and I'm %d years old\n", p.name, p.age)
// }
