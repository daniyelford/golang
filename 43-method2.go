package main

import (
	"fmt"
)

// defining method with non struct reciever
type number int // creating custom type
func (n number) squareofNumber() number {
	return n * n
}
func main() {
	no1 := number(5)
	result := no1.squareofNumber()
	fmt.Printf("Square of number %d is %d", no1, result)
}
