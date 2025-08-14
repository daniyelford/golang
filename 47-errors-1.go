package main

import (
	"errors"
	"fmt"
)

func main() {
	// وقتی panic اتفاق بیفته
	// همون لحظه اجرای نرمال رو متوقف می‌کنه
	// Go شروع می‌کنه به unwind کردن stack (پایین اومدن از توابع).
	// همه‌ی deferهایی که تو مسیر هستن اجرا می‌شن.
	// در نهایت، اگر خطا مدیریت نشه، برنامه با پیام خطا بسته میشه

	// recover() یه تابع مخصوصه که فقط داخل یک defer کار می‌کنه.
	// وظیفه‌ش اینه که panic رو بگیره و برنامه رو به حالت عادی برگردونه

	// fmt.Println("شروع برنامه")

	// defer func() {
	//     if r := recover(); r != nil {
	//         fmt.Println("یک panic رخ داد:", r)
	//     }
	// }()

	// fmt.Println("قبل از panic")
	// panic("یه خطای جدی!") // ایجاد panic
	// fmt.Println("این خط هرگز اجرا نمیشه")

	fmt.Println(safeDivide(30, 5)) // خروجی: 6
	fmt.Println(safeDivide(30, 0)) // پیام خطا
	fmt.Println("برنامه ادامه پیدا کرد")

	x := "salam"
	MyError := errors.New("invalid option")
	if x != "goodbye" {
		fmt.Println(MyError)
	}

	result, err := divide(30, 0)
	if err != nil {
		fmt.Println("خطا:", err)
	} else {
		fmt.Println("نتیجه:", result)
	}

}
func safeDivide(a, b int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("تقسیم بر صفر اتفاق افتاد!")
		}
	}()
	return a / b
}
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("تقسیم بر صفر مجاز نیست")
	}
	return a / b, nil
}

//New()
//Erorf()
