package main

import (
	"fmt"
	"math"
)

// معادله درجه دو
func main() {
	var a, b, c, root1, root2, delta float64

	fmt.Println("enter a ,b,c of quadratic equationin a consecutive order: ")
	fmt.Scanln(&a, &b, &c)

	delta = (b * b) - (4 * a * c)

	if delta > 0 {
		root1 = (-b + math.Sqrt(delta)) / (2 * a)
		root2 = (-b - math.Sqrt(delta)) / (2 * a)
		fmt.Println("two distinct roots exist:root1 =", root1, "and root2 is: ", root2)
	} else if delta == 0 {
		root1 = (-b) / (2 * a)
		root2 = (-b) / (2 * a)
		fmt.Println("two equal roots exist: root1 = root2 =", root1)
	} else {
		fmt.Println("delta is negative")
	}

	fmt.Println("Pi:", math.Pi)
	fmt.Println("2 به توان 3:", math.Pow(2, 3))
	fmt.Println("ریشه دوم 25:", math.Sqrt(25))
	fmt.Println("سینوس 90 درجه:", math.Sin(90*math.Pi/180))
	fmt.Println("گرد به بالا:", math.Ceil(3.1))
	fmt.Println("گرد به پایین:", math.Floor(3.9))
	fmt.Println("بیشترین:", math.Max(10, 20))
}

/*

math

| ثابت          | مقدار             | توضیح             |
| ------------- | ----------------- | ----------------- |
| `math.Pi`     | 3.141592653589793 | عدد π             |
| `math.E`      | 2.718281828459045 | عدد e (عدد نپر)   |
| `math.Phi`    | 1.618033988749895 | نسبت طلایی        |
| `math.Sqrt2`  | √2                | ریشه دوم 2        |
| `math.SqrtE`  | √e                | ریشه دوم e        |
| `math.SqrtPi` | √π                | ریشه دوم π        |
| `math.Ln2`    | ln(2)             | لگاریتم طبیعی ۲   |
| `math.Ln10`   | ln(10)            | لگاریتم طبیعی ۱۰  |
| `math.Log2E`  | log₂(e)           | لگاریتم ۲ پایه e  |
| `math.Log10E` | log₁₀(e)          | لگاریتم ۱۰ پایه e |


| تابع             | توضیح       | مثال                  |
| ---------------- | ----------- | --------------------- |
| `math.Pow(x, y)` | x به توان y | `math.Pow(2, 3) // 8` |
| `math.Sqrt(x)`   | ریشه دوم    | `math.Sqrt(16) // 4`  |
| `math.Cbrt(x)`   | ریشه سوم    | `math.Cbrt(27) // 3`  |



| تابع            | توضیح                  | مثال                    |
| --------------- | ---------------------- | ----------------------- |
| `math.Exp(x)`   | e به توان x            | `math.Exp(1) // e`      |
| `math.Log(x)`   | لگاریتم طبیعی (پایه e) | `math.Log(math.E) // 1` |
| `math.Log10(x)` | لگاریتم پایه ۱۰        | `math.Log10(100) // 2`  |
| `math.Log2(x)`  | لگاریتم پایه ۲         | `math.Log2(8) // 3`     |


| تابع                                           | توضیح                            |
| ---------------------------------------------- | -------------------------------- |
| `math.Sin(x)`, `math.Cos(x)`, `math.Tan(x)`    | سینوس، کسینوس، تانژانت           |
| `math.Asin(x)`, `math.Acos(x)`, `math.Atan(x)` | معکوس‌های سینوس، کسینوس، تانژانت |
| `math.Atan2(y, x)`                             | زاویه از مختصات (به رادیان)      |

توجه: ورودی‌ها بر حسب رادیان هستن، نه درجه.
برای تبدیل:

radian := degree * math.Pi / 180
degree := radian * 180 / math.Pi



| تابع            | توضیح                 | مثال                   |
| --------------- | --------------------- | ---------------------- |
| `math.Floor(x)` | گرد به پایین          | `math.Floor(3.7) // 3` |
| `math.Ceil(x)`  | گرد به بالا           | `math.Ceil(3.1) // 4`  |
| `math.Round(x)` | گرد به نزدیک‌ترین عدد | `math.Round(3.5) // 4` |
| `math.Trunc(x)` | حذف اعشار             | `math.Trunc(3.9) // 3` |



| تابع             | توضیح                   | مثال                  |
| ---------------- | ----------------------- | --------------------- |
| `math.Abs(x)`    | قدر مطلق                | `math.Abs(-5) // 5`   |
| `math.Max(x, y)` | بیشترین مقدار           | `math.Max(3, 7) // 7` |
| `math.Min(x, y)` | کمترین مقدار            | `math.Min(3, 7) // 3` |
| `math.Mod(x, y)` | باقیمانده تقسیم (مثل %) | `math.Mod(7, 3) // 1` |



| ثابت/تابع             | توضیح                                |
| --------------------- | ------------------------------------ |
| `math.NaN()`          | تولید مقدار NaN (Not a Number)       |
| `math.IsNaN(x)`       | بررسی NaN بودن                       |
| `math.Inf(sign)`      | بی‌نهایت مثبت یا منفی (sign=1 یا -1) |
| `math.IsInf(x, sign)` | بررسی بی‌نهایت بودن                  |


*/
