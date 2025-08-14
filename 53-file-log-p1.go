package main

// log → برای لاگ‌گیری و نوشتن پیام‌های log استفاده میشه.
// os → برای کار با فایل‌ها و سیستم عامل (مثلاً باز کردن فایل) استفاده میشه
import (
	"log"
	"os"
)

func main() {
	// f, err := os.Create("test.txt") // می‌سازه یا خالی می‌کنه
	// f, err := os.Open("test.txt") // فقط باز کردن برای خواندن
	// f, err := os.OpenFile("test.txt", os.O_RDWR|os.O_CREATE, 0666)

	// os.O_RDONLY → فقط خواندن
	// os.O_WRONLY → فقط نوشتن
	// os.O_RDWR → خواندن و نوشتن
	// os.O_APPEND → نوشتن به انتهای فایل
	// os.O_CREATE → ایجاد فایل اگر وجود نداشت
	// os.O_TRUNC → پاک کردن محتوا قبل از نوشتن

	// info, err := os.Stat("test.txt")
	// fmt.Println("Name:", info.Name())
	// fmt.Println("Size:", info.Size())
	// fmt.Println("IsDir:", info.IsDir())

	// بررسی وجود یا نبود فایل
	// if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
	// fmt.Println("فایل وجود ندارد")
	// }

	// حذف فایل یا دایرکتوری
	// os.Remove("test.txt")         // حذف فایل
	// os.RemoveAll("folder_name")   // حذف همه محتویات پوشه

	// ساخت پوشه
	// os.Mkdir("mydir", 0755)
	// os.MkdirAll("a/b/c", 0755) // ساخت چندین سطح

	// os.Stdin → ورودی استاندارد (مثل fmt.Scanln)
	// os.Stdout → خروجی استاندارد (چاپ در ترمینال)
	// os.Stderr → خروجی خطا

	// متغیرهای محیطی (Environment Variables)
	// 🔹 گرفتن مقدار
	// value := os.Getenv("PATH")
	// fmt.Println(value)

	// 🔹 تنظیم مقدار
	// os.Setenv("MY_KEY", "12345")

	// 🔹 حذف مقدار
	// os.Unsetenv("MY_KEY")

	// 🔹 گرفتن همه متغیرها
	// for _, env := range os.Environ() {
	// 	fmt.Println(env)
	// }

	// 4️⃣ اطلاعات سیستم و کاربر
	// name, _ := os.Hostname()
	// fmt.Println("Hostname:", name)
	// dir, _ := os.Getwd()
	// fmt.Println("Current dir:", dir)

	// تغییر مسیر جاری:
	// os.Chdir("/tmp")

	// 5️⃣ بستن فایل
	// هر فایلی که باز می‌کنی باید ببندی:
	// f, _ := os.Open("test.txt")
	// defer f.Close()

	// 6️⃣ اجرا و مدیریت Process
	/*
		os.ProcAttr یک struct هست که ویژگی‌های پروسس جدید رو مشخص می‌کنه.
	*/
	// procAttr := &os.ProcAttr{

	/*
		Files یک آرایه از اشاره‌گرهای فایل (*os.File) هست که ورودی و خروجی استاندارد پروسس رو تعیین می‌کنه:
		os.Stdin → ورودی استاندارد (کیبورد یا input)
		os.Stdout → خروجی استاندارد (ترمینال)
		os.Stderr → خروجی خطا (ترمینال برای نمایش خطاها)
	*/
	// 	Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},

	// }
	// یعنی ما اینجا گفتیم که پروسس جدید هم از همون ورودی و خروجی ترمینال اصلی استفاده کنه
	/*
		آرگومان اول /bin/ls → مسیر باینری یا executable برنامه‌ای که می‌خوای اجرا کنی.
		(اینجا دستور ls که لیست فایل‌های دایرکتوری رو نشون میده.)
		آرگومان دوم → slice از stringها که شامل اسم برنامه و پارامترهاش هست.
		اینجا {"ls", "-l"} یعنی همون ls -l در لینوکس.
		آرگومان سوم → همون procAttr که تعیین می‌کنه پروسس جدید چه ورودی/خروجی داشته باشه
	*/
	// proc, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)

	/*
		این متد باعث میشه برنامه فعلی منتظر بمونه تا پروسس جدیدی که اجرا شده کامل بشه
	*/
	// proc.Wait()

	/*
	   os.StartProcess یک متد سطح پایین (low-level) هست. یعنی مستقیماً با سیستم‌عامل کار می‌کنه و امکاناتی مثل PATH رو به صورت خودکار مدیریت نمی‌کنه.
	   برای همین معمولاً توصیه میشه از پکیج os/exec استفاده کنی که راحت‌تر و high-level هست
	*/
	// cmd := exec.Command("ls", "-l")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// cmd.Run()

	// 7️⃣ خروج از برنامه
	// os.Exit(0)   // خروج موفق
	// os.Exit(1)   // خروج با خطا

	// در Go 1.16 به بعد، بیشتر توابع ioutil به os و io منتقل شده:
	// ioutil.ReadFile → os.ReadFile
	// ioutil.WriteFile → os.WriteFile
	// ioutil.ReadDir → os.ReadDir

	// os.OpenFile یه فایل باز می‌کنه یا ایجاد می‌کنه (O_CREATE).
	// O_WRONLY → فقط نوشتن
	// O_APPEND → اگر فایل موجود باشه، داده‌ها به انتهای فایل اضافه میشن
	// 0666 → سطح دسترسی فایل (خواندن و نوشتن برای همه)
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// اگر خطا باشه (err != nil) → log.Fatalln پیام خطا چاپ می‌کنه و برنامه متوقف میشه
		log.Fatalln("Error opening log file ")
	}
	// به جای نمایش لاگ در کنسول، همه‌ی logها در فایل app.log نوشته میشن.
	log.SetOutput(file)

	// پیام ساده لاگ می‌نویسه و یک خط جدید اضافه می‌کنه
	log.Println("this is a log message.")

	// مثل Printf تو fmt عمل می‌کنه و می‌تونی متغیرها رو داخل متن بگذاری
	log.Printf("logging a new variable%s\t", "reza")

	// Fatalf ترکیبی از Printf و os.Exit(1) هست
	log.Fatalf("this  is a fatal log massage")
	// بعد از log.Fatalf هیچ خطی از برنامه اجرا نمیشه
}

/*
   اگر بخوای لاگ‌ها رو همزمان در کنسول و فایل ببینی، باید از io.MultiWriter استفاده کنی.
   استفاده از log برای پروژه‌های ساده خوبه، ولی برای پروژه‌های بزرگ معمولاً از کتابخانه‌هایی مثل zap یا logrus استفاده می‌کنن
*/
