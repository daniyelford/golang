package main

import (
	"fmt"
	"time"
)

/*
channel یه سازوکار برای ارتباط بین goroutineها هست.
می‌تونی تصورش کنی مثل یه صف (queue) که چند goroutine می‌تونن داخلش چیزی بذارن (send) و چند goroutine دیگه می‌تونن ازش چیزی بردارن (receive).
*/
// وقتی کانال با ظرفیت (buffer) ساخته بشه
// make(chan channel_type,capacity)
func main() {
	/*
		یعنی ظرفیت کانال = ۲.
		می‌تونی تا ۲ تا مقدار داخلش بذاری بدون اینکه نیاز باشه کسی همون لحظه اونارو بخونه.
		وقتی ظرفیت پر شد، send بلاک میشه تا یکی از مقادیر خونده بشه.
		وقتی کانال خالی شد و کسی بخواد بخونه، دوباره بلاک میشه تا چیزی فرستاده بشه
	*/
	ch := make(chan string, 2)
	ch <- "hello"     //ارسال
	ch <- "salam"     //ارسال
	fmt.Println(<-ch) //دریافت
	fmt.Println(<-ch) //دریافت

	go func() {
		ch <- "msg 1"
		ch <- "msg 2"
		fmt.Println("sent 2 messages")
	}()

	time.Sleep(time.Second) // صبر کن goroutine مقدارها رو بفرسته

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	go worker(1, ch)
	go worker(2, ch)
	// go worker(3, ch)

	// فرستادن ۶ پیام
	for i := 1; i <= 6; i++ {
		ch <- fmt.Sprintf("Task %d", i)
		fmt.Println("sent task", i)
	}

	close(ch) // بعد از ارسال همه کارها کانال رو می‌بندیم
	time.Sleep(time.Second * 5)
}
func worker(id int, ch chan string) {
	// هرچی از کانال بیاد، بخون و داخل msg بذار
	// وقتی کانال بسته بشه (close(ch))، حلقه خودش متوقف میشه
	for msg := range ch {
		fmt.Printf("Worker %d received: %s\n", id, msg)
		time.Sleep(time.Second)
	}
}
