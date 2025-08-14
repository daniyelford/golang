package main

// seed
func main() {
	// rand.Seed(time.Now().UnixNano())
	// randomNumber := rand.Intn(100) + 1
	// fmt.Println(randomNumber)
	// از Go 1.20 به بعد دیگه rand.Seed() عملاً تاریخ مصرفش گذشته و Go خودش به صورت پیش‌فرض یه Seed تصادفی داخلی برای math/rand ست می‌کنه
	// r := rand.New(rand.NewSource(42))
	// فقط وقتی می‌خوای یه دنباله‌ی تکراری و کنترل‌شده از اعداد تصادفی بسازی (برای تست و شبیه‌سازی) از rand.New(rand.NewSource(seed)) استفاده کن.

	// rand.Seed(123)
	// for i := 0; i < 5; i++ {
	// 	fmt.Println(rand.Intn(100)) // عدد بین 0 و 99
	// }

	// src := rand.NewSource(123) // seed ثابت
	// r := rand.New(src)         // مولد مستقل

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(r.Intn(100))
	// }

	// rand.Seed(time.Now().UnixNano())
	// id := fmt.Sprintf("%d-%d", time.Now().Unix(), rand.Intn(10000))
	// // // id := fmt.Sprintf("%d", rand.Intn(time.Now().Unix()))
	// id1 := fmt.Sprintf("%d", rand.Intn(int(time.Now().Unix())))
	// id2 := fmt.Sprintf("%d", rand.Int63n(time.Now().Unix()))
	// fmt.Println(id, ' ', id1, ' ', id2)

	// seed2 := int64(123)
	// r := rand.New(rand.NewSource(seed2))

	// num := r.Int63()                      // یک عدد تصادفی ثابت
	// id4 := num % int64(time.Now().Unix()) // شبیه Intn
	// id5 := num % time.Now().Unix()        // شبیه Int63n

	// fmt.Println(id4, id5) // دقیقا یکسان

	// rand.Seed(time.Now().UnixNano())

	// fmt.Println("عدد بین 0 تا 99:", rand.Intn(100))
	// fmt.Println("عدد اعشاری:", rand.Float64())

	// // مثال: تولید رمز ساده
	// letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	// b := make([]rune, 8)
	// fmt.Println(b)
	// for i := range b {
	// 	b[i] = letters[rand.Intn(len(letters))]
	// }
	// fmt.Println("رمز:", string(b))

	// now := time.Now()
	// fmt.Println("الان:", now)

	// fmt.Println("سال:", now.Year())
	// fmt.Println("ماه:", now.Month())
	// fmt.Println("روز:", now.Day())

	// fmt.Println("فرمت:", now.Format("2006-01-02 15:04:05"))

	// time.Sleep(2 * time.Second)
	// fmt.Println("بعد از دو ثانیه:", time.Now())
}

/*

time

| تابع/متد                             | توضیح                        | مثال                              |
| ------------------------------------ | ---------------------------- | --------------------------------- |
| `time.Now()`                         | گرفتن زمان فعلی              | `t := time.Now()`                 |
| `t.Year()` / `t.Month()` / `t.Day()` | گرفتن اجزای تاریخ            | `year := t.Year()`                |
| `t.Format(layout)`                   | فرمت‌دهی تاریخ               | `t.Format("2006-01-02 15:04:05")` |
| `time.Sleep(d)`                      | توقف برنامه برای مدت مشخص    | `time.Sleep(2 * time.Second)`     |
| `time.Since(t)`                      | فاصله زمانی از یک زمان مشخص  | `time.Since(start)`               |
| `time.Until(t)`                      | فاصله تا یک زمان خاص         | `time.Until(deadline)`            |
| `t.Add(d)`                           | اضافه کردن مدت به یک زمان    | `t.Add(24*time.Hour)`             |
| `t.Sub(t2)`                          | کم کردن دو زمان → `Duration` | `t1.Sub(t2)`                      |


math/rand

| تابع              | توضیح                             | مثال                               |
| ----------------- | --------------------------------- | ---------------------------------- |
| `rand.Int()`      | عدد تصادفی بین 0 و حداکثر int     | `rand.Int()`                       |
| `rand.Intn(n)`    | عدد تصادفی بین 0 و n-1            | `rand.Intn(10)`                    |
| `rand.Float64()`  | عدد اعشاری بین 0 و 1              | `rand.Float64()`                   |
| `rand.Seed(seed)` | مقدار اولیه برای تولید عدد تصادفی | `rand.Seed(time.Now().UnixNano())` |


crypto/rand

| تابع                                  | توضیح                                | خروجی            |
| ------------------------------------- | ------------------------------------ | ---------------- |
| `rand.Read(b []byte)`                 | پر کردن آرایه با بایت‌های تصادفی امن | تعداد بایت + خطا |
| `rand.Int(rand.Reader, max *big.Int)` | عدد تصادفی امن بین 0 و max-1         | `*big.Int` + خطا |
| `rand.Prime(rand.Reader, bits int)`   | تولید عدد اول امن با تعداد بیت مشخص  | `*big.Int` + خطا |



| ویژگی        | `math/rand`        | `crypto/rand`              |
| ------------ | ------------------ | -------------------------- |
| سرعت         | سریع‌تر            | کندتر                      |
| امنیت        | غیرامن             | امن                        |
| مناسب برای   | شبیه‌سازی، بازی‌ها | پسورد، توکن، کلید رمزنگاری |
| نیاز به Seed | بله                | خیر                        |


*/
