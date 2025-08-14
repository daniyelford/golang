package main

import (
	"fmt"
)

func sum(a, b int) int {
	result := a + b
	fmt.Println("result is ", result)
	return 0
}

// defer is lifo(last in first out) filo
func main() {

	// fmt.Println("hello world")
	// defer sum(5, 10)
	// defer sum(5, 7)
	// defer fmt.Println("این آخر اجرا می‌شود")
	// fmt.Println("این اول اجرا می‌شود")

	// f, err := os.Open("Readme.md")
	// if err != nil {
	// 	fmt.Println("خطا در باز کردن فایل:", err)
	// 	return
	// }
	// fmt.Println(f)
	// defer f.Close()
	// fmt.Println("فایل با موفقیت باز شد.")

}

// باعث می‌شه اجرای یک تابع یا دستور به تعویق بیفته و بعد از پایان اجرای تابع فعلی اجرا بشه.
// به عبارت ساده‌تر، وقتی در تابعی defer می‌نویسی، اون دستور یا تابع همیشه بعد از اجرای همه کدهای تابع جاری و قبل از خروج از اون تابع اجرا می‌شه

// برای پاک‌سازی منابع (مثل بستن فایل، اتصال شبکه، قفل‌ها و غیره) به صورت مطمئن، حتی اگر تابع زودتر یا با خطا خارج شود.
// برای کدهایی که باید حتماً در پایان اجرا بشن، بدون نیاز به نوشتن صدا زدن آنها در چند جای مختلف تابع
