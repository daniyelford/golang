package main

import (
	"fmt"
)

func bubblesort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
func main() {
	nums := []int{65, 34, 25, 12}
	fmt.Println("unsorted array: ", nums)
	bubblesort(nums)
	fmt.Println("sorted array: ", nums)
}

/*
n=4
0<=i<3
0<=j<3-i

i=0
{
    j=0
    [34,65,25,12]
    j=1
    [34,25,65,12]
    j=2
    [34,25,12,65]
}
i=1
{
    j=0
    [25,34,12,65]
    j=1
    [25,12,34,65]
}
i=2
{
    j=0
    [12,25,34,65]
}
*/
