package main

import "fmt"

//<-
//var channelName chan type
// channelName:=make (chan channeltype,buffer)

func main() {
	x := make(chan string)

	go func() { // goroutine
		x <- "salam" //sent to cha

	}()
	salam := <-x //recieve from channel
	fmt.Println(salam)
}

/*

| ویژگی              | `make`                                    | `new`                        | Literal (`[]int{}`, `&Type{}`)                        |
| ------------------ | ----------------------------------------- | ---------------------------- | ----------------------------------------------------- |
| **کاربرد**         | فقط برای **map**، **slice** و **channel** | ساخت حافظه صفر شده از هر نوع | ساخت مقدار اولیه با داده مشخص یا خالی                 |
| **نوع خروجی**      | خود نوع (map، slice، channel)             | *pointer* به نوع داده        | خود نوع یا pointer (بسته به استفاده از `&`)           |
| **مقداردهی اولیه** | بله (ظرفیت، طول، و غیره)                  | همه فیلدها صفر میشن          | هر مقداری که خودت بدی                                 |
| **تفاوت اصلی**     | برای آماده کردن ساختارهای داخلی خاص       | فقط حافظه خام رو میده        | آزادی کامل در مقداردهی                                |
| **مثال**           | `make(chan int, 3)`                       | `p := new(int)`              | `nums := []int{1,2,3}` یا `p := &Person{Name: "Ali"}` |

*/
