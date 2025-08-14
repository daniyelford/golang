package main

import "fmt"

func main() {

	var a [2]string
	a[0] = "first element"
	a[1] = "second element"

	fmt.Println(a)
	fmt.Println(a[0], a[1])

	primaryNumbers := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(primaryNumbers[0], primaryNumbers[1], primaryNumbers[2], primaryNumbers[3], primaryNumbers[4], primaryNumbers[5])

	//elipsis

	myarr := [...]int{2, 3, 5, 7, 11, 13, 14}

	for i := 0; i < len(myarr); i++ {
		fmt.Printf("%d\n", myarr[i])
	}

	//for  index ,   value := range array      returns both index and values in array

	names := [3]string{"ali", "reza", "daniyal"}
	for index, val := range names {
		fmt.Printf("%v\t%v\n", index, val)
	}

}
