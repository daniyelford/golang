package main

import (
	"fmt"
)

func main() {
	var array_size int
	fmt.Println("Enter the size of the array: ")
	fmt.Scanln(&array_size)
	fmt.Println("enter the array element ")
	array := make([]int, array_size)
	for i := 0; i < array_size; i++ { //length of array
		fmt.Scanln(&array[i])
	}
	max := array[0]
	for i := 1; i < array_size; i++ {
		if max < array[i] {
			max = array[i]
		}
	}
	fmt.Printf("\n maximum element of array is:%d", max)
}
