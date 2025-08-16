package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	/*
		n یک عدد صحیح (int) است.
		float64(n)
		تابع math.Sqrt عدد با نوع float64 می‌خواد،
		پس n رو به float64 تبدیل می‌کنیم تا بتوانیم جذرش را حساب کنیم.
		math.Sqrt(float64(n))
		جذر عدد n را محاسبه می‌کنه.
		خروجی float64 است.
		int(math.Sqrt(float64(n)))
		جذر به int تبدیل میشه (قسمت اعشاری حذف میشه).
		یعنی فقط قسمت صحیح جذر باقی میمونه.
		i <= int(math.Sqrt(float64(n)))
		یعنی ما میخوایم i کمتر یا مساوی با جذر n باشه.
	*/
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func PrimeCreator(limit int, ch chan int) {
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			ch <- i
		}
	}
	close(ch)
}

func main() {
	limit := 50
	ch := make(chan int)
	go PrimeCreator(limit, ch)
	for prime := range ch {
		fmt.Println(prime)
	}
}
