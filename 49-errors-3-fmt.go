package main

import(
	"fmt"
	
)
//errorf
func divide(num1,num2 int) error {
	if num2==0{
		return fmt.Errorf("%d  %d\n number cant be divided into zero ",num1,num2)
	}
	return nil
}

func main(){

	err:=divide(10,0)
	if err !=nil{
        fmt.Println(err)
    }else{
        fmt.Println("no error  valid operation")
    }
}