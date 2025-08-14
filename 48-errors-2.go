package main

import (
	"errors"
	"fmt"
)

func checkName(name string) error {
	newError := errors.New("invalid name")
	if name != "reza" {
		return newError
	}
	return nil
}

func main() {

	name := "reza"
	err := checkName(name)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Name is valid")
	}

}
