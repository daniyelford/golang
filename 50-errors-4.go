package main

import "fmt"

//creating custom error
// interface استاندارد error در Go
type error interface {
	Error() string
}

type DivisionByZero struct {
	message string
}

func (z DivisionByZero) Error() string {
	return "invalid operation"
}

func divide(n1 int, n2 int) (int, error) {
	if n2 == 0 {
		return 0, &DivisionByZero{}
	} else {
		return n1 / n2, nil
	}
}

func main() {
	number1 := 20
	number2 := 0

	result, err := divide(number1, number2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The result is: %d", result)
	}

}
