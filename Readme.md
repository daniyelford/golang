# Golang

Linux Install Golang 

```bash
	sudo apt update 
	sudo apt search golang-go 
	sudo apt search gccgo-go 
	sudo apt install golang-go
```

------------------------------------------

Linux Install Git

```bash
sudo apt update
sudo apt-get install git
```

------------------------------------------
Install vscode and git for windows

[Install vscode](https://code.visualstudio.com/download)

[Install Git](https://gitforwindows.org)

------------------------------------------

``` bash
go mod init golang
go mod tidy
```

------------------------------------------

| ابزار        | کاربرد                   |
| ------------ | ------------------------ |
| `go run`     | اجرای مستقیم برنامه      |
| `go build`   | ساخت باینری              |
| `go install` | نصب باینری در مسیر سیستم |
| `go get`     | دریافت/آپدیت پکیج        |
| `go mod`     | مدیریت وابستگی‌ها        |
| `go fmt`     | فرمت‌بندی کد             |
| `go vet`     | بررسی مشکلات منطقی       |
| `go test`    | اجرای تست‌ها             |
| `go doc`     | مشاهده مستندات           |
| `go env`     | مشاهده تنظیمات محیطی     |
| `go clean`   | پاکسازی خروجی‌ها         |

------------------------------------------

go tool

| ابزار       | کاربرد                                                     |
| ----------- | ---------------------------------------------------------- |
| **compile** | کامپایل فایل Go به object code (مرحله قبل از build نهایی). |
| **link**    | لینک کردن فایل‌های object برای ساخت باینری نهایی.          |
| **pprof**   | تحلیل پروفایل CPU و حافظه (برای بهینه‌سازی سرعت).          |
| **cover**   | گزارش میزان پوشش تست (Test Coverage).                      |
| **vet**     | بررسی مشکلات رایج (همون که با `go vet` هم اجرا میشه).      |
| **trace**   | تحلیل رفتار اجرای برنامه به صورت دقیق.                     |
 
------------------------------------------

```bash
go tool pprof myapp cpu.pprof
```

```go
package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	// ساخت فایل خروجی پروفایل CPU
	f, err := os.Create("cpu.pprof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(f) // شروع پروفایل
	defer pprof.StopCPUProfile() // توقف پروفایل در پایان

	// یک کد نمونه که CPU رو مصرف می‌کنه
	for i := 0; i < 50000000; i++ {
		fmt.Sprintf("num: %d", i)
	}

	time.Sleep(time.Second)
}
 ```
 
```bash
go run thisfile.go
```

------------------------------------------


 | دستور         | توضیح                                               |
| ------------- | --------------------------------------------------- |
| `top`         | توابعی که بیشترین مصرف CPU رو داشتن (به ترتیب درصد) |
| `list <func>` | نمایش کد تابع با هایلایت مصرف CPU                   |
| `web`         | تولید نمودار گرافیکی و باز کردنش در مرورگر          |
| `svg`         | ذخیره نمودار به فرمت SVG                            |
| `help`        | لیست همه دستورات                                    |

