package main

// net/http → پکیج داخلی Go برای ساخت سرور و کلاینت HTTP
import (
	"fmt"
	"net/http"
)

// w http.ResponseWriter → مثل خروجی سرور هست، هر چی توش بنویسی، به مرورگر یا کلاینت فرستاده میشه.
// r *http.Request → اطلاعات درخواست اومده از کاربر (مثل URL، هدرها، پارامترها و ...)
func handler(w http.ResponseWriter, r *http.Request) {
	// رشته‌ی "Hello world" رو به پاسخ HTTP می‌فرسته
	fmt.Fprintf(w, "Hello world")
}

func main() {
	// هر وقت کاربر به مسیر / (صفحه اصلی) درخواست زد، تابع handler رو اجرا کن
	http.HandleFunc("/", handler)
	fmt.Println("server started at :8080")
	// روی پورت 8080 گوش بده (localhost:8080).
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("error starting server :", err)
	}
}
