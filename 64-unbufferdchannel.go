package main

import "fmt"

/*
	اگه کانال بدون ظرفیت (make(chan string)) ساخته بشه:
	وقتی چیزی توش send می‌کنی، تا وقتی که یه goroutine دیگه نیاد اون داده رو receive کنه، برنامه متوقف میشه (block میشه).
	برعکس، وقتی می‌خوای چیزی receive کنی ولی داخل کانال چیزی نیست، باز هم متوقف میشه.
	یعنی ارتباط همزمان (synchronous) هست
*/
func main() {
	x := make(chan string)
	go func(y string) {
		x <- y
	}("Hello, World!")
	m := <-x
	fmt.Println(m)
}
