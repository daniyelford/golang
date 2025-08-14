package main

import "fmt"

func main() {

	array := []int{1, 2, 3, 4}

	sum := 0

	// for _, val := range array {
	// 	sum += val
	// }

	for i := 0; i < len(array); i++ {
		sum += array[i]
	}

	avg := float64(sum) / float64(n)

	fmt.Println("Sum of array elements is:", sum, "and avg of elements is:", avg)

}
