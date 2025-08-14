package main

func main() {
	// fmt.Println(strings.IndexAny("golang fghfghwr rtgbwrtb tbwtb", "laeiou"))
	// fmt.Println(strings.LastIndexAny("golang", "laeiou"))
	// fmt.Println(strings.ToValidUTF8("aÿa\xffb", ""))
	// fmt.Println(strings.ReplaceAll("oink oink", "oin k", "moo"))
	// Sentence := "hello i am learning golang golang "
	// // بررسی می‌کنه که آیا 'reza' داخل sentence هست یا نه
	// fmt.Println(strings.Contains(Sentence, "reza"))

	// // آیا Sentence با golang تموم می‌شه؟
	// fmt.Println(strings.HasSuffix(Sentence, "golang"))

	// // آیا Sentence با golang شروع می‌شه؟
	// fmt.Println(strings.HasPrefix(Sentence, "hello"))

	// // چند بار golang در Sentence تکرار شده؟
	// fmt.Println(strings.Count(Sentence, "golang"))

	// //strings.cut         [x is a string string]    strings.cut(x,"phrase")   delete phrase from x string

	// fmt.Println(strings.Cut(Sentence, "reza"))

	// //len    length of arguments    len(args)

	// fmt.Println(len(Sentence))

	// //strings.repeat       [x is a string]      strings.repeat(x,"phrase")
	// fmt.Println(strings.Repeat("name ", 10))

	// //strings.replace    [x is a string]     strings.replace(x,"old phrase","new phrase ",number of replacement )

	// fmt.Println(strings.Replace(Sentence, "hello", "salam ", 1))
	// fmt.Println(strings.Replace(Sentence, "golang", "go", 1))

	// //strings.replaceAll        [x is a string]    strings.replaceAll(x,"old phrase","new phrase")
	// // همه‌ی golang ها رو با go جایگزین می‌کنه
	// x := strings.ReplaceAll(Sentence, "golang", "go")
	// fmt.Println(x)

	// //strings.compare       [x is a string]     strings.compare("phrase 1","phrase 2")   for key sensitive comparison

	// fmt.Println(strings.Compare("ali", "reza"))
	// fmt.Println(strings.Compare("reza", "reza"))
	// fmt.Println(strings.Compare("ali", "Reza"))

	// //strings.equalfold   [x is a string] strings.equalfold("phrase 1","phrase2")   key sensitive is not important

	// fmt.Println(strings.EqualFold("daniyal", "daniyal"))

	// //strings.Index     strings.index("phrase "," chosen character ")   return chosen character index

	// fmt.Println(strings.Index("daniyal", "a"))

	// //upper case  & lower case  & title

	// fmt.Println(strings.ToUpper("daniyal"))
	// fmt.Println(strings.ToLower("daniyal"))
	// fmt.Println(strings.Title("hello i live in iran"))

	// //strings.trim      [x is a string]    strings.trim("phrase "," chosen part ") remove chosen part

	// fmt.Println(strings.Trim("   golang is good  ", " "))
	// fmt.Println(strings.TrimLeft("IR626265462265", " IR"))
	// fmt.Println(strings.TrimRight("626265462265IR", " IR"))

	// //strings.split   [x is a string] strings.split(x,sptr)

	// fmt.Println(strings.Split(Sentence, " "))

	// fmt.Println(len(strings.Split(Sentence, " ")) - 1)

}

// ### 🔍 بررسی وجود و جستجو

// #### 1. `Contains`

// بررسی می‌کند آیا زیررشته‌ای در رشته‌ی اصلی وجود دارد یا نه.

// ```go
// strings.Contains("golang is cool", "cool") // خروجی: true
// ```

// #### 2. `ContainsAny`

// بررسی می‌کند آیا هر کدام از کاراکترهای داده شده در رشته وجود دارد یا نه.

// ```go
// strings.ContainsAny("team", "i") // خروجی: false
// strings.ContainsAny("fail", "ui") // خروجی: true
// ```

// #### 3. `ContainsRune`

// بررسی وجود یک `rune` خاص در رشته.

// ```go
// strings.ContainsRune("hello", 'e') // خروجی: true
// ```

// #### 4. `Count`

// تعداد تکرارهای بدون همپوشانی یک زیررشته را در رشته شمارش می‌کند.

// ```go
// strings.Count("cheese", "e") // خروجی: 3
// ```

// #### 5. `Index` و `LastIndex`

// مکان اولین/آخرین وقوع یک زیررشته را مشخص می‌کند.
// indexs
// c => 0
// h => 1
// i => 2
// c => 3
// k => 4
// e => 5
// n => 6
// ```go
// strings.Index("chicken", "ken")      // خروجی: 4
// strings.LastIndex("go gopher", "go")  // خروجی: 3
// ```

// #### 6. `IndexAny` و `LastIndexAny`

// توابع IndexAny و LastIndexAny بررسی می‌کنند که اولین یا آخرین کاراکتری که در مجموعه‌ای از کاراکترها تعریف شده، در رشته وجود دارد یا نه، و سپس موقعیت (ایندکس) آن را برمی‌گردانند.

// اگر هیچ کاراکتری از مجموعه پیدا نشود، مقدار -1 برگردانده می‌شود.

// IndexAny(s, chars) → موقعیت اولین کاراکتر موجود در chars

// LastIndexAny(s, chars) → موقعیت آخرین کاراکتر موجود در chars

// 📌 کاربرد: مفید برای بررسی اینکه کدام یک از چندین کاراکتر خاص (مثلاً واکه‌ها یا علامت‌ها) در رشته وجود دارد.

// strings.IndexAny("golang", "laeiou")      // خروجی: 1 → کاراکتر 'o'
// strings.LastIndexAny("golang", "aeiou")  // خروجی: 3→ کاراکتر 'a'
// strings.IndexAny("golang", "xyz")        // خروجی: -1 → هیچ کدام از این حروف در رشته نیست

// 💡 نکته: ترتیب اهمیت دارد، یعنی اولین یا آخرین در رشته اصلی بررسی می‌شود، نه ترتیب حروف داده شده.

// #### 7. `IndexByte`, `LastIndexByte`

// این دو تابع برای یافتن مکان (ایندکس) اولین یا آخرین وقوع یک بایت خاص در رشته استفاده می‌شوند. بایت همان معادل عددی یک کاراکتر در جدول
// ASCII است، بنابراین این توابع مخصوص کاراکترهای تک‌بایتی (مانند حروف انگلیسی ساده) هستند.

// IndexByte(s, c) → موقعیت اولین بایت c در رشته s

// LastIndexByte(s, c) → موقعیت آخرین بایت c در رشته s

// 📌 تفاوت با IndexRune:
// اگر رشته شامل کاراکترهای چندبایتی (مثل حروف فارسی یا کاراکترهای یونیکد خاص) باشد، باید از
// IndexRune استفاده کنید؛ چون
// IndexByte فقط روی تک‌بایتی‌ها جواب درست می‌دهد.

// 🧪 مثال:
// strings.IndexByte("golang", 'a')        // خروجی: 3 → 'a' در موقعیت 3 است
// strings.LastIndexByte("hello", 'l')     // خروجی: 3 → آخرین 'l' در موقعیت 3
// strings.IndexByte("سلام", 'س')          // خروجی: -1 → چون 'س' چندبایتی است

// 💡 نکته: این توابع بسیار سریع هستند چون مستقیماً با بایت‌ها کار می‌کنند. برای کاراکترهای لاتین مناسب‌ترین انتخاب‌اند.

// #### 8. `IndexRune`

// مکان اولین occurrence از یک `rune`.

// ```go
// strings.IndexRune("gé", 'é') // خروجی: 1
// ```

// #### 9. `IndexFunc` و `LastIndexFunc`

// جستجوی اولین/آخرین کاراکتری که تابع داده شده true برگرداند.

// ```go
// isDigit := func(r rune) bool { return r >= '0' && r <= '9' }
// strings.IndexFunc("a3b2c", isDigit) // خروجی: 1
// strings.LastIndexFunc("a1b2c", isDigit) // خروجی: 3
// ```

// ---

// ### ✂️ برش رشته و تبدیل آن به آرایه

// #### 10. `Split`, `SplitN`, `SplitAfter`, `SplitAfterN`

// رشته را با جداکننده تقسیم می‌کند.
// این توابع برای تقسیم رشته به بخش‌های کوچکتر بر اساس جداکننده (separator) استفاده می‌شوند:

// Split → رشته را بر اساس جداکننده می‌شکند و خود جداکننده حذف می‌شود.

// SplitN → مشابه Split ولی حداکثر N بخش تولید می‌کند.

// SplitAfter → جداکننده را در انتهای هر بخش نگه می‌دارد.

// SplitAfterN → مشابه SplitAfter ولی با حداکثر N بخش.

// 📌 کاربرد: زمانی که بخواهیم رشته را به اجزای معین تقسیم کنیم، مثلا برای تجزیهٔ CSV یا پارس کردن آدرس‌ها.

// 🔎 تفاوت‌ها:

// Split و SplitN جداکننده را حذف می‌کنند.

// SplitAfter و SplitAfterN جداکننده را نگه می‌دارند.

// 🧪 مثال:
// func(x,y) x=>"a,b,c" , y=>","
// strings.Split("a,b,c", ",")           // ["a" "b" "c"]
// strings.SplitN("a,b,c", ",", 2)       // ["a" "b,c"]
// strings.SplitAfter("a,b,c", ",")     // ["a," "b," "c"]
// strings.SplitAfterN("a,b,c", ",", 2) // ["a," "b,c"]

// 💡 نکته: در SplitN و SplitAfterN اگر n <= 0 باشد، نتیجه هیچ محدودیتی ندارد (مانند Split و SplitAfter).

// #### 11. `Fields`

// تقسیم رشته بر اساس فاصله‌ها و فضای سفید.

// ```go
// strings.Fields("  foo bar  baz  ") // ["foo" "bar" "baz"]
// ```

// #### 12. `FieldsFunc`

// تقسیم رشته با شرط دلخواه.

// ```go
// f := func(r rune) bool { return r == ',' || r == ';' }
// strings.FieldsFunc("a,b;c", f) // ["a" "b" "c"]
// ```

// ---

// ### 🔗 ترکیب و ساخت رشته

// #### 13. `Join`

// ترکیب رشته‌ها با جداکننده مشخص شده.

// ```go
// array => ["a" "b" "c"]
// array to string => {"a","b","c"}

// strings.Join([]string{"a", "b", "c"}, ",") // خروجی: "a,b,c" string
// ```

// #### 14. `Repeat`

// تکرار رشته به تعداد مشخص.

// ```go
// strings.Repeat("na", 2) // خروجی: "nana"
// ```

// ---

// ### 🔤 تغییر حالت حروف

// #### 15. `ToUpper`, `ToLower`, `ToTitle`

// تبدیل حروف به بزرگ، کوچک یا حالت عنوان.

// ```go
// strings.ToUpper("go")   // "GO"
// strings.ToLower("Go")   // "go"
// strings.ToTitle("go lang") // "Go Lang"
// ```

// #### 16. `ToUpperSpecial`, `ToLowerSpecial`, `ToTitleSpecial`

// تغییرات با درنظر گرفتن زبان خاص (با استفاده از پکیج `x/text`).

// #### 17. `ToValidUTF8`

// حذف یا جایگزینی کاراکترهای نامعتبر UTF-8.

// ```go
// strings.ToValidUTF8("aÿa\xffb", "") // خروجی: "aÿab"
// ```

// ---

// ### ✂️ حذف و برش اطراف رشته

// #### 18. `Trim`, `TrimLeft`, `TrimRight`

// حذف کاراکترهای مشخص‌شده از ابتدا/انتها/دوطرف.

// ```go
// strings.Trim(" !!! He ! llo !!! ", " !") // "Hello"
// ```

// #### 19. `TrimSpace`

// حذف فاصله‌های اطراف رشته.

// ```go
// strings.TrimSpace("  Hello  ") // "Hello"
// ```

// #### 20. `TrimPrefix`, `TrimSuffix`

// حذف پیشوند/پسوند از رشته.

// ```go
// strings.TrimPrefix("foobar", "foo") // "bar"
// strings.TrimSuffix("foobar", "bar") // "foo"
// ```

// #### 21. `TrimFunc`, `TrimLeftFunc`, `TrimRightFunc`

// حذف بر اساس شرط دلخواه.

// ```go
// f := func(r rune) bool { return r == '!' }
// strings.TrimFunc("!!!Hello!!!", f) // "Hello"
// ```

// ---

// ### 🪄 سایر توابع

// #### 22. `HasPrefix`, `HasSuffix`

// بررسی شروع/پایان رشته با مقدار خاص.

// ```go
// strings.HasPrefix("golang", "go")  // true
// strings.HasSuffix("golang", "ang") // true
// ```

// #### 23. `Map`

// اعمال تغییر روی هر کاراکتر با تابع دلخواه.

// ```go
// rot13 := func(r rune) rune {
// 	if r >= 'a' && r <= 'z' {
// 		return 'a' + (r-'a'+13)%26
// a=1
// f=6
// 1 + (((6-1+13)=18)%26=18)=19
// 	}
// 	return r
// }
// strings.Map(rot13, "hello") // "uryyb"
// ```

// #### 24. `EqualFold`

// مقایسه برابر بودن رشته‌ها بدون حساسیت به حروف بزرگ/کوچک.

// ```go
// strings.EqualFold("Go", "go") // true
// ```

// #### 25. `Replace`, `ReplaceAll`

// جایگزینی زیررشته‌ها در رشته.

// ```go
// strings.Replace("oink oink oink", "k", "ky", 2)   // "oinky oinky oink"
// strings.ReplaceAll("oink oink", "oink", "moo")     // "moo moo"
// ```

// #### 26. `Cut`, `CutPrefix`, `CutSuffix` (از Go 1.18 به بعد)

// بریدن رشته در محل خاص یا حذف پیشوند/پسوند در قالب خروجی دو مقداری.

// ```go
// before, after, found := strings.Cut("key:value", ":")
// before = "key"
// after = "value"
// found = true
// newStr, found := strings.CutPrefix("unhappy", "un")
// newStr = "happy"
// found = true
// newStr, found := strings.CutSuffix("filename.txt", ".txt")
// newStr = "filename"
// found = true

// ---

// ### ⚙️ توابع داخلی پکیج (غیر عمومی)

// تابع‌هایی مانند:

// * `makeASCIISet`
// * `indexFunc`
// * `lastIndexFunc`
// * `trimLeftByte`
// * `trimRightByte`
// * `trimLeftUnicode`
// * `trimRightUnicode`
// * `isSeparator`

// توابع داخلی پکیج هستند و برای استفاده مستقیم طراحی نشده‌اند.
