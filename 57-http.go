package main

// go get github.com/gorilla/mux
import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// مسیر صفحه اصلی
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to the homepage!")
	})

	r.HandleFunc("/user/{name}/age/{age}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		age := vars["age"]
		fmt.Fprintf(w, "Hello %s! You are %s years old.", name, age)
		// پارامترها همیشه رشته برمی‌گردن. اگه بخوای با age عملیات ریاضی انجام بدی، باید با strconv.Atoi(age) تبدیلش کنی به عدد
		fmt.Fprintf(os.Stdout, "Hello %s! You are %s years old.", name, age)
		// fmt.Fprintf(file, "Hello %s! You are %s years old.", name, age)
	})

	// مسیر صفحه درباره ما
	r.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to about page")
	})

	// مسیر صفحه تماس
	r.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to contact page. Contact us on GitHub")
		// fmt.Fprintf(os.Stdout, "Welcome to contact page. Contact us on GitHub")
	})

	// شروع سرور
	fmt.Println("Server started at :9090")
	http.ListenAndServe(":9090", r)
}
