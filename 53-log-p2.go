package main

import (
	"io"
	"log"
	"os"
)

func main() {
	// باز کردن فایل لاگ
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Error opening log file:", err)
	}

	// ایجاد MultiWriter برای کنسول و فایل
	multi := io.MultiWriter(file, os.Stdout)

	// تنظیم log برای نوشتن روی MultiWriter
	log.SetOutput(multi)

	// نمونه لاگ‌ها
	log.Println("this is a log message")
	log.Printf("logging a variable: %s\n", "reza")
}
