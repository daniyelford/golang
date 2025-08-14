package main
import "fmt"

func main() {
	x:=make (chan string)

	go func() {
        x <- "Hello, World!"
    }()
	m:=<-x
	fmt.Println(m)
}