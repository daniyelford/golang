package main
import(
	"time"
	"fmt"
)

func main() {


	ch1:=make(chan string)
	ch2:=make(chan string)

	go func (){

		time.Sleep(1*time.Second)
		ch1<-"hello world"
	}()

	go func(){

        time.Sleep(2*time.Second)
        ch2<-"slam donya "
    }()

	select {
	case a:=<-ch1:
		fmt.Println(a)
	case b:=<-ch2:
		fmt.Println(b)
	}






}