package main

// go run 58-main.go BTC
import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// تو Go وقتی مقدار ثابت (constant) تعریف می‌کنی، می‌تونی نوع (type) رو ننویسی. دلیلش اینه که Go می‌تونه خودش از مقدار بفهمه چه نوعیه (type inference)
// فرقش اینه که وقتی type مشخص نکنی، Go اون constant رو به عنوان untyped constant می‌سازه، که انعطاف بیشتری داره

const apiKey = "d2a37fe2-3152-4915-8f08-070acfbdd5e6" // Replace with your actual API key
const apiUrl = "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest"

// تگ json:"price" یعنی وقتی JSON رو به struct تبدیل می‌کنیم، مقدار کلید price در JSON داخل Price ریخته میشه.
type Quote struct {
	Price float64 `json:"price"`
}

// این struct یعنی هر ارز یک quote داره که خودش یک map هست.
// مثلا: "quote": { "USD": { "price": 123.45 } }
type CryptoData struct {
	Quote map[string]Quote `json:"quote"`
}

// این struct برای نگه داشتن داده‌های اصلی پاسخ API ساخته شده. توی JSON، data یک map از symbol به CryptoData است.
type Response struct {
	Data map[string]CryptoData `json:"data"`
}

func getTokenPrice(symbol string) (float64, error) {
	// با fmt.Sprintf آدرس نهایی API ساخته میشه.
	url := fmt.Sprintf("%s?symbol=%s", apiUrl, symbol)
	// یک درخواست HTTP از نوع GET ساخته میشه.
	req, _ := http.NewRequest("GET", url, nil)
	// هدرها اضافه میشن:
	// Accepts: application/json → یعنی خروجی JSON باشه.
	req.Header.Add("Accepts", "application/json")
	// X-CMC_PRO_API_KEY → کلید API که برای احراز هویت لازم هس
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)

	// client := &http.Client{
	// 	Timeout: 5 * time.Second,
	// }
	// یک کلاینت HTTP ساخته میشه
	client := &http.Client{}
	// درخواست اجرا میشه
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	// یعنی وقتی تابع تموم شد، اتصال بسته بشه
	defer resp.Body.Close()
	// بدنه پاسخ (Body) به طور کامل خونده میشه
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	// یک متغیر response از نوع Response ساخته میشه
	var response Response
	// json.Unmarshal → JSON رو به struct های ما تبدیل می‌کنه
	// حالا باید بتونیم این JSON رو به ساختار Go تبدیل کنیم (unmarshal)
	// تو Go وقتی یه JSON از بیرون می‌گیری (مثلاً از API یا فایل)، اون یه رشته‌ی متنی هست ([]byte). اما تو برنامه نمی‌تونی مستقیم با رشته کار کنی، باید تبدیلش کنی به ساختار Go (struct یا map)
	// body->ورودی، یعنی داده‌ی JSON خام (به شکل []byte)
	// &response->خروجی، یعنی آدرس struct یا mapی که می‌خوای داده بریزه داخلش
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, err
	}

	if data, exists := response.Data[symbol]; exists {
		return data.Quote["USD"].Price, nil
	}

	return 0, fmt.Errorf("the token %s not found", symbol)
}

func main() {
	// اگه کاربر توی خط فرمان هیچ سمبلی وارد نکرده باشه، پیام راهنما چاپ میشه
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <token_symbol>")
		return
	}

	symbol := os.Args[1]
	fmt.Println(symbol)
	price, err := getTokenPrice(symbol)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("The price of %s is $%.2f\n", symbol, price)
}
