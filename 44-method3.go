package main

import (
	"fmt"
)

//method with pointer reciever
// قانون کلی:
// اگر میخوای متدت داده‌ها رو تغییر بده → باید pointer receiver استفاده کنی.
// اگر میخوای فقط بخونی و تغییر ندی → می‌تونی value receiver استفاده کنی.

type person struct {
	name string
}

func (p *person) ChangeName(newName string) {
	p.name = newName
}

// ما updateName رو به صورت *person تعریف کردیم، ولی داریم روی مقدار (x) صدا میزنیم نه روی آدرس (&x).
// در Go، کامپایلر به طور خودکار مقدار رو به آدرس تبدیل می‌کنه اگر متد نیاز به pointer داشته باشه
func main() {
	x := person{name: "John"}
	fmt.Println(x.name)

	(&x).ChangeName("reza")
	fmt.Println(x.name)

}
