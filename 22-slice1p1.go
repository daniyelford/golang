package main

import (
	"fmt"
)

func main() {
	//   slice_name := [ ] datatype {value}

	//myslice := [] int {1,2,3}   //length =3    //capacity=3

	myslice := []int{}
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))

	myslice1 := []string{"my", "name", "is", "daniyal"}
	fmt.Println(myslice1)
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))

	//slice_name := array[start index:end index]

	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice2 := arr1[2:4]
	fmt.Printf("myslice =%v\n", myslice2)
	fmt.Printf("len = %d\n", len(myslice2))
	fmt.Printf("cap = %d\n", cap(myslice2))

	//slice_name := make ([] datatype ,len,cap)

	slice := make([]int, 5, 10)
	fmt.Printf("slice = %v\n", slice)
	fmt.Printf("cap = %d\n", cap(slice))
	fmt.Printf("len = %d\n", len(slice))

}
