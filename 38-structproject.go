package main

import (
	"fmt"
	"os"
)

type Product struct {
	Price    float64
	Name     string
	ID       int64
	Quantity int
}

func showMenu() {
	fmt.Println("1. Add product ")
	fmt.Println("2. update product ")
	fmt.Println("3. delete product ")
	fmt.Println("4. display list of products ")
	fmt.Println("5. Exit ")
}

func addProduct() {
	var name string
	var quantity int
	var price float64
	fmt.Println("enter product name: ")
	fmt.Scanln(&name)
	fmt.Println("enter product price: ")
	fmt.Scanln(&price)
	fmt.Println("enter product quantity: ")
	fmt.Scanln(&quantity)
	inventory = append(inventory, Product{ID: nextID, Name: name, Price: price, Quantity: quantity})
	nextID++
	fmt.Println("product added successfully")
}

func updateProduct() {
	var id int64
	fmt.Println("enter product id: ")
	fmt.Scanln(&id)
	for i, product := range inventory {
		if product.ID == id {
			var name string
			var quantity int
			var price float64
			fmt.Println("enter updated product name: ")
			fmt.Scanln(&name)
			fmt.Println("enter updated product price: ")
			fmt.Scanln(&price)
			fmt.Println("enter updated product quantity: ")
			fmt.Scanln(&quantity)
			inventory[i] = Product{ID: id, Name: name, Price: price, Quantity: quantity}
			fmt.Println("product updated")
			return
		}
	}
	fmt.Println("product not found")
}

func deleteProduct() {
	var id int64
	fmt.Println("enter product id: ")
	fmt.Scanln(&id)
	for i, product := range inventory {
		if product.ID == id {
			inventory = append(inventory[:i], inventory[i+1:]...)
			fmt.Println("product deleted ")
			return
		}
	}
	fmt.Println("product not found")
}

func displayList() {
	if len(inventory) == 0 {
		fmt.Println("inventory is empty")
		return
	}
	fmt.Println("list of products: ")
	for _, product := range inventory {
		fmt.Printf("ID: %d, Name: %s,quantity: %d,Price: %f\n", product.ID, product.Name, product.Quantity, product.Price)
	}
}

var inventory []Product

//	اسلایس‌ها و مپ‌ها به صورت پیش‌فرض با رفرنس منتقل میشن، یعنی:
//
// اگه توی تابع تغییرش بدی، در بیرون هم تغییر می‌کنه.
// نیازی به &inventory و *inventory نیست.
// 📌 اگر آرایه عادی ([5]Product) بود، باید از پوینتر استفاده می‌کردیم، چون آرایه‌ها به صورت کپی منتقل میشن.
var nextID int64 = 1

func main() {

	for {
		showMenu()
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			addProduct()
		case 2:
			updateProduct()
		case 3:
			deleteProduct()
		case 4:
			displayList()
		case 5:
			fmt.Println("exiting...")
			return
		default:
			fmt.Println("invalid choice")
			os.Exit(0)
		}
	}
}
