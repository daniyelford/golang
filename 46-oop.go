package main

import (
	"fmt"
)

// use interface for demonstrating abstraction
// چند نوع مختلف از veichle با رفتار متفاوت داشته باشیم
// هر چیزی که start() و stop() رو پیاده‌سازی کنه، من اون رو یک وسیله نقلیه می‌شناسم
// با یک قرارداد کلی کار کنیم، بدون اینکه لازم باشه بدونیم دقیقا Car هست یا Bike

type veichle interface {
	start()
	stop()
}

// encapsulation
// از بیرون فقط از طریق متدها میشه باهاش کار کرد
type Car struct {
	brand string
	year  int
}

// constructor method for car
// نمونه جدید Car می‌سازه و آدرسش (*Car) رو برمی‌گردونه
func NewCar(brand string, year int) *Car {
	// مقدار دهی آدرس Car
	return &Car{brand: brand, year: year}
}

func (c Car) start() {
	fmt.Println(c.brand, "car is starting")
}

func (c Car) stop() {
	fmt.Println(c.brand, "car is stopping")
}

// encapsulation
type bike struct {
	brand string
	year  int
}

// constructor method for bike
func NewBIKE(brand string, year int) *bike {
	return &bike{brand: brand, year: year}
}

func (b bike) start() {
	fmt.Println(b.brand, "bike is starting")
}

func (b bike) stop() {
	fmt.Println(b.brand, "bike is stopping")
}

// polymorphism
// اینجا x یک تابع هست که هر چیزی رو می‌پذیره به شرطی که interface veichle رو پیاده‌سازی کرده باشه
func x(v veichle) {
	v.start()
	v.stop()
}

// creating two instance/objects in main function for car and bike

func main() {
	car := NewCar("bwm", 2015)
	bike := NewBIKE("honda", 2020)
	x(car)
	x(bike)
}
