package main

import (
	"bufio"         // برای خوندن ورودی از کاربر (خط به خط)
	"bytes"         // برای ساختن buffer بایت‌ها (بدنۀ POST request)
	"encoding/json" // برای تبدیل map یا struct به JSON
	"fmt"           // برای چاپ و فرمت‌دهی متن
	"io"            // برای خوندن و نوشتن فایل و response
	"log"           // برای مدیریت خطا و لاگ گرفتن
	"net/http"      // برای ارسال درخواست HTTP (GET, POST و ...)
	"os"            // برای کار با سیستم‌عامل (مثلا stdin)
	"strings"       // برای کار با رشته‌ها (trim, split و ...)
)

func main() {
	apiKey := ""
	model := "gemini-1.5-pro"
	apiURL := fmt.Sprintf("https://generativelanguage.googleapis.com/v1/models/%s:generateContent?key=%s", model, apiKey)
	// یک Reader می‌سازه که از ورودی ترمینال می‌خونه
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Ask a Question: ")
	// تا وقتی کاربر Enter بزنه می‌خونه
	question, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading input:%v", err)
	}
	// برای حذف کاراکترهای اضافی مثل \n یا فاصله
	question = strings.TrimSpace(question)
	requestBody := map[string]interface{}{
		"contents": []map[string]interface{}{
			{"parts": []map[string]string{
				{"text": question},
			}},
		},
	}
	// Marshal → map رو به JSON تبدیل می‌کنه (رشته‌ی باینری)
	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		log.Fatalf("Error creating json files:%v", err)
	}
	/*
		http.NewRequest("POST", ...) → درخواست POST به اون آدرس.
		bytes.NewBuffer(jsonData) → بدنه‌ی JSON رو به buffer بایت تبدیل می‌کنه تا بشه ارسال کرد.
	*/
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatalf("Error creating request POST:%v", err)
	}
	// req.Header.Set → میگه که این داده application/json هست.
	req.Header.Set("Content-type", "application/json")
	// http.Client → یک کلاینت HTTP می‌سازه.
	client := &http.Client{}
	// client.Do(req) → درخواست POST رو ارسال می‌کنه.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request:%v", err)
	}
	// defer resp.Body.Close() → بعد از پایان کار، connection بسته میشه.
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response:%v", err)
	}
	outputFile := "output.txt"
	err = os.WriteFile(outputFile, body, 0644)
	if err != nil {
		log.Fatalf("Error writing to file:%v", err)
	}
	fmt.Printf("\n Response saved in '%s' \n", outputFile)
}
