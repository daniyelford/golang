package main

import "fmt"

func main() {
	const name string = "ali"
	// name2 := "reza"
	//name := "daniyal"
	age := 20
	height := 179.6

	// fmt.Println("example height", height)

	fmt.Print("my name is ", name, " my age is ", age, " my height is ", height)
	// %s → رشته (string)
	// %d → عدد صحیح
	// %f → عدد اعشاری
	// %t → بولین
	// %v → نمایش پیش‌فرض هر نوع
	// %+v → نمایش با نام فیلدها برای struct
	// %T → نوع متغیر
	// %q → رشته با کوتیشن
	// fmt.Printf(" my name is %s and my height is %f  and my age is %d", name2, height, age)
	// fmt.Printf(" nama : my name is %s ", name2)
	// fmt.Printf(" nama : my name is %f ", height)
	// fmt.Printf(" nama : my name is %d ", age)

}

/*
| تابع             | توضیح                                 |
| ---------------- | ------------------------------------- |
| `fmt.Sprintf()`  | مثل `Printf` ولی رشته رو برمی‌گردونه  |
| `fmt.Sprint()`   | مثل `Print` ولی رشته رو برمی‌گردونه   |
| `fmt.Sprintln()` | مثل `Println` ولی رشته رو برمی‌گردونه |

| تابع            | توضیح              | مثال                        |
| --------------- | ------------------ | --------------------------- |
| `fmt.Print()`   | چاپ بدون خط جدید   | `fmt.Print("Hello")`        |
| `fmt.Println()` | چاپ با خط جدید     | `fmt.Println("Hello")`      |
| `fmt.Printf()`  | چاپ با فرمت دلخواه | `fmt.Printf("عدد: %d", 42)` |

| تابع           | توضیح                  |
| -------------- | ---------------------- |
| `fmt.Scan()`   | گرفتن ورودی با فاصله   |
| `fmt.Scanln()` | گرفتن ورودی تا خط جدید |
| `fmt.Sscan()`  | گرفتن ورودی از رشته    |


*/
