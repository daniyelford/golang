package main

import "fmt"

/*
channel یه سازوکار برای ارتباط بین goroutineها هست.
می‌تونی تصورش کنی مثل یه صف (queue) که چند goroutine می‌تونن داخلش چیزی بذارن (send) و چند goroutine دیگه می‌تونن ازش چیزی بردارن (receive).
*/
// make(chan channel_type,capacity)
func main() {
	ch := make(chan string, 2)
	ch <- "hello"     //ارسال
	ch <- "salam"     //ارسال
	fmt.Println(<-ch) //دریافت
	fmt.Println(<-ch) //دریافت
}
