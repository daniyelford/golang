package main

import (
	"fmt"
)

// powInt: توان صحیح (base^exp) با عدد صحیح — سریع و ایمن
func powInt(base, exp int) int {
	if exp == 0 {
		return 1
	}
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

// isArmstrong: بررسی می‌کند n عدد Armstrong هست یا نه
func isArmstrong(n int) bool {
	if n < 0 {
		return false
	}
	// شمارش تعداد رقم‌ها
	temp := n
	digits := 0
	if temp == 0 {
		digits = 1
	} else {
		for temp > 0 {
			digits++
			temp /= 10
		}
	}

	// محاسبه جمع توان‌های رقم‌ها
	sum := 0
	temp = n
	for temp > 0 {
		d := temp % 10
		sum += powInt(d, digits)
		temp /= 10
	}

	return sum == n
}

func main() {
	var t int
	fmt.Println("insert a number to check is Armstrong ")
	fmt.Scanln(&t)
	fmt.Printf("%d -> Armstrong? %v\n", t, isArmstrong(t))

	// لیست اعداد Armstrong تا 10000
	/*
		fmt.Println("\nArmstrong numbers up to 10000:")
		for i := 0; i <= 10000; i++ {
			if isArmstrong(i) {
				fmt.Println(i)
			}
		}
	*/

	// مثال‌های تکی
	/*
		tests := []int{0, 1, 5, 153, 370, 371, 9474, 9475}
		for _, t := range tests {
			fmt.Printf("%d -> Armstrong? %v\n", t, isArmstrong(t))
		}
	*/
}
