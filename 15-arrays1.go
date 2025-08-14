package main

import "fmt"

func main() {

	arr := [4]string{"reza", "ali", "farhad", "samane"}
	fmt.Println("elements of array is ", arr)

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	Myarr := [3][3]string{{"csharp", "c", "python"}, {"golang", "java", "ruby"}}

	fmt.Println("Elements of array1 is: ")

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(Myarr[i][j], i, j, " ")
		}
	}

	var array [2][2]int
	array[0][0] = 100
	array[0][1] = 200
	array[1][0] = 300
	array[1][1] = 400

	fmt.Println("Elements of array2 is: ")
	for x := 0; x < 2; x++ {
		for y := 0; y < 2; y++ {
			fmt.Print(array[x][y], " ")
		}

	}

}
