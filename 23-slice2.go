package main

import (
	"fmt"
	"sort"
)

func main() {

	// array := [5]int{1, 2, 3, 4, 5}
	// slices := array[1:4]
	// // array[0] = 1
	// // {
	// 	// array[1] = 2
	// 	// array[2] = 3
	// 	// array[3] = 4
	// // }
	// // array[4] = 5
	// fmt.Println("Original array:", array)
	// fmt.Println("Sliced array:", slices)

	// append function add element to slice

	// slice := []int{1, 2, 3}
	// slice = append(slice, 4, 5, 6)
	// fmt.Println("Appended slice:", slice)

	// arr := [7]string{"hello ", "my", "first", "name", "is", "daniyal", "fard"}
	// fmt.Println("Array:", arr)
	// myslice := arr[1:6]
	// fmt.Println("slices:", myslice)
	// fmt.Printf("len is:  %d\n", len(myslice))
	// fmt.Printf("cap is: %d\n", cap(myslice))

	// var myslice = make([]int, 4, 7)

	// fmt.Printf("myslice is %v\n  len is %d   \ncap is %d\n", myslice, len(myslice), cap(myslice))

	// var myslice2 = make([]int, 4)

	// fmt.Printf("myslice is %v\n  len is %d   \ncap is %d\n", myslice2, len(myslice2), cap(myslice2))

	// myslice := []string{"this", "is", "the", "tutorial", "of", "golang", "language"}
	// fmt.Printf("myslice is %v\n ", myslice)
	// for i := 0; i < len(myslice); i++ {

	// 	fmt.Println(myslice[i])
	// }

	// for _, element := range myslice {
	// 	fmt.Printf("element is  %s\n ", element)
	// }

	a := []int{45, 23, 67, 90, 33, 87, 56, 57}
	fmt.Println("before sort :", a)

	sort.Ints(a) // sort

	fmt.Println("after sort :", a)

}

/*
| تابع / اینترفیس | امضاء (Signature)                                         | کاربرد                                                                |
| --------------- | --------------------------------------------------------- | --------------------------------------------------------------------- |
| `sort.Ints`     | `func Ints(a []int)`                                      | مرتب‌سازی اسلایس اعداد صحیح به صورت صعودی                             |
| `sort.Strings`  | `func Strings(a []string)`                                | مرتب‌سازی اسلایس رشته‌ها به صورت صعودی                                |
| `sort.Float64s` | `func Float64s(a []float64)`                              | مرتب‌سازی اسلایس اعداد اعشاری                                         |
| `sort.Slice`    | `func Slice(slice interface{}, less func(i, j int) bool)` | مرتب‌سازی اسلایس دلخواه با استفاده از تابع مقایسه سفارشی              |
| `sort.Sort`     | `func Sort(data Interface)`                               | مرتب‌سازی هر نوع داده که اینترفیس `sort.Interface` را پیاده کرده باشد |
| `sort.Stable`   | `func Stable(data Interface)`                             | مرتب‌سازی پایدار (حفظ ترتیب عناصر برابر)                              |
| `sort.IsSorted` | `func IsSorted(data Interface) bool`                      | بررسی اینکه داده‌ها مرتب هستند یا خیر                                 |
| `sort.Search`   | `func Search(n int, f func(int) bool) int`                | جستجوی دودویی در داده مرتب‌شده                                        |


| متد    | امضاء                 | توضیح                                     |
| ------ | --------------------- | ----------------------------------------- |
| `Len`  | `Len() int`           | تعداد کل عناصر                            |
| `Less` | `Less(i, j int) bool` | مشخص می‌کند عنصر i کوچکتر از j هست یا خیر |
| `Swap` | `Swap(i, j int)`      | جایگزینی مکان عناصر i و j                 |



*/
