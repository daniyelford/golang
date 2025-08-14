package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
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
