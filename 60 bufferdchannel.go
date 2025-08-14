package main
import "fmt"

func main() {
	ch:=make(chan string,2 )
	ch<-"hello"
	ch<-"salam"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	


}