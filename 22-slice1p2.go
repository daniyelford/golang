package main

import "fmt"

func main() {
	nums := make([]int, 2, 5) // طول 2، ظرفیت 5
	fmt.Println("nums:", nums)
	fmt.Println("len:", len(nums)) // 2
	fmt.Println("cap:", cap(nums)) // 5

	// اضافه کردن عنصر
	nums = append(nums, 10, 20, 30)
	fmt.Println("nums:", nums)
	fmt.Println("len:", len(nums)) // 5
	fmt.Println("cap:", cap(nums)) // 5 (هنوز ظرفیت پر نشده)
}
