package main

import "fmt"

func main() {

	var bike, motorcycle, ashkan string

	fmt.Println("please provide your input: ")

	fmt.Scanln(&bike, &motorcycle, &ashkan)
	fmt.Printf("Your inputs are: %s  %s  %s \n", bike, motorcycle, ashkan)

}
