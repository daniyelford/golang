package main

import (
    "fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Welcome to home page")
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w,"Welcome to about page")
}

func contactHandler(w http.ResponseWriter, r *http.Request){
    fmt.Fprint(w,"Welcome to contact page contact us on github")
}


func main(){
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)	
	http.HandleFunc("/contact", contactHandler)
	fmt.Println("server started at :9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("error in server",err)
	}
}