package main
import "fmt"

func main() {
	x:=make (chan int)

	go func(){
		for i:=0; i<=5; i++ {
			x <- i
	}
	close(x)
}()

for num := range x{
	fmt.Println(num)
	}
}