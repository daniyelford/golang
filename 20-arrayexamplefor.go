package main

import "fmt"

func main() {
	fruits := [8]string{"apple", "mango", "orange", "peach", "strawberry", "potato", "onion", "banana"}
	for idx, fruits := range fruits { //%v  points to value of sth
		fmt.Printf("%v\t%v\n", idx, fruits)
	}
}
