package main
import(
	"fmt"
	"time"
)

func Print(m string){
	for i:=1;i<=5;i++{
		fmt.Println(m ,":",i)
		time.Sleep(time.Millisecond*1000)
	}
}
func main(){
	go Print("goroutine 1")
	go Print("goroutine 2")

	time.Sleep(time.Second*5)
	fmt.Println("hello world")

}