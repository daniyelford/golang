package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {

		time.Sleep(1 * time.Second)
		ch1 <- "hello world"
	}()

	go func() {

		time.Sleep(2 * time.Second)
		ch2 <- "slam donya "
	}()
	// select دقیقاً مثل switch عمل می‌کنه، با این تفاوت که برای کانال‌ها استفاده میشه
	select {
	case a := <-ch1:
		fmt.Println(a)
	case b := <-ch2:
		fmt.Println(b)
	}

	// select {
	// 	case msg1 := <-ch1:
	// 		fmt.Println("دریافت از ch1:", msg1)
	// 	case msg2 := <-ch2:
	// 		fmt.Println("دریافت از ch2:", msg2)
	// 	// case ch3 <- "salam":
	// 	// 	fmt.Println("ارسال روی ch3 انجام شد")
	// 	default:
	// 		fmt.Println("هیچ کانالی آماده نبود")
	// }

	/*
		هر case مربوط به یک عملیات روی کاناله (دریافت یا ارسال).
		select منتظر میشه ببینه کدوم کانال آماده است.
		وقتی یکی آماده شد → همون case اجرا میشه.
		اگه چندتا همزمان آماده باشن → یکی به صورت تصادفی انتخاب میشه.
		اگه هیچکدوم آماده نباشن و default داشته باشیم → همون اجرا میشه.
		اگه default نداشته باشیم → بلاک میشه تا یکی آماده بشه

		وقتی چند کانال داریم و نمی‌دونیم کدوم زودتر آماده میشه، با select منتظرشون میمونیم.
		خیلی به درد همزمانی (concurrency) و ارتباط بین goroutineها می‌خوره.
		مثل epoll یا select در سیستم‌عامل‌ها عمل می‌کنه، ولی به شکل ساده‌تر.
	*/
}
