package main

import (
	"fmt"
	"time"
)

// این تابع یه پیام (m) می‌گیره
func Print(m string) {
	//  ۵ بار اون رو چاپ می‌کنه
	for i := 1; i <= 5; i++ {
		fmt.Println(m, ":", i)
		// بین هر بار چاپ هم یک ثانیه (1000 میلی‌ثانیه) صبر می‌کنه
		time.Sleep(time.Millisecond * 1000)
	}
}
func main() {
	// وقتی قبل از صدا زدن تابع از go استفاده می‌کنی، اون تابع در یک goroutine جدید اجرا میشه
	// goroutine یه thread خیلی سبک هست که توسط runtime خود Go مدیریت میشه.
	// یعنی همزمان (concurrent) با بقیه کدها اجرا میشه.
	// در اینجا، Print("goroutine 1") به صورت موازی (نه پشت سر هم) با کدهای دیگه اجرا میشه
	go Print("goroutine 1")
	// دومین goroutine هم ساخته میشه و همزمان با اولی اجرا میشه
	go Print("goroutine 2")
	// از اونجایی که main() وقتی تموم بشه، کل برنامه بسته میشه (حتی اگه goroutineها هنوز در حال اجرا باشن)، ما اینجا ۵ ثانیه sleep گذاشتیم تا دو تا goroutine فرصت کنن اجرا بشن
	time.Sleep(time.Second * 4) // اجرای برنامه کامل نمیشه و 5 بار نمایش داده نمیشه
	// time.Sleep(time.Second * 5)
	fmt.Println("hello world")

}
