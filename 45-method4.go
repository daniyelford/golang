package main

import (
    "fmt"
)
type person struct {
	name string
}

func(p*person)updateName(newName string){
	p.name = newName
}

func(p person)ShowInfo(){
	fmt.Println("Name: ", p.name)
}

func main() {
	x:=person{name:"John"}
	x.updateName("reza")
	fmt.Println(x.name)

	(&x).ShowInfo()

}