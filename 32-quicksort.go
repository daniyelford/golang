package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	left, right := 0, len(arr)-1
	pivot := arr[len(arr)/2]
	// left = 0
	// right = 4
	// pivot = arr[2] = 7
	for i := range arr {
		if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	/*
		   | گام | i | arr\[i] | شرط (arr\[i] < pivot) | آرایه قبل جابه‌جایی | جابه‌جایی انجام میشه؟                | left بعد |
		   | --- | - | ------- | --------------------- | ------------------- | ------------------------------------ | -------- |
		   | 1   | 0 | 8       | 8 < 7 ❌               | \[8, 3, 7, 4, 2]    | ❌ هیچی انجام نمی‌شه                  | 0        |
		   | 2   | 1 | 3       | 3 < 7 ✅               | \[8, 3, 7, 4, 2]    | arr\[1] ↔ arr\[0] → \[3, 8, 7, 4, 2] | 1        |
		   | 3   | 2 | 7       | 7 < 7 ❌               | \[3, 8, 7, 4, 2]    | ❌ هیچی انجام نمی‌شه                  | 1        |
		   | 4   | 3 | 4       | 4 < 7 ✅               | \[3, 8, 7, 4, 2]    | arr\[3] ↔ arr\[1] → \[3, 4, 7, 8, 2] | 2        |
		   | 5   | 4 | 2       | 2 < 7 ✅               | \[3, 4, 7, 8, 2]    | arr\[4] ↔ arr\[2] → \[3, 4, 2, 8, 7] | 3        |




		   	شروع: left = 0
			[ 8,  3,  7,  4,  2 ]
			^
			left

			گام 2: بعد از جابه‌جایی 3 و 8 → left = 1
			[ 3,  8,  7,  4,  2 ]
					^
				left

			گام 4: بعد از جابه‌جایی 4 و 8 → left = 2
			[ 3,  4,  7,  8,  2 ]
						^
					left

			گام 5: بعد از جابه‌جایی 2 و 7 → left = 3
			[ 3,  4,  2,  8,  7 ]
							^
							left

		   =>
		   	{
				[3, 4, 2, 8, 7]
				left = 3
			}
	*/
	arr[left], arr[right] = arr[right], arr[left]
	// [3, 4, 2, 7, 8]
	quickSort(arr[:left])
	// 3, 4, 2
	quickSort(arr[left+1:])
	// 8
	return arr
}

func main() {
	nums := []int{8, 3, 7, 4, 2}
	fmt.Println(nums)            //unsorted array
	fmt.Println(quickSort(nums)) //sort array

}
