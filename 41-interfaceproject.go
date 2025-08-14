package main

import "fmt"

type Book struct {
	Title  string
	Author string
}

type Library interface {
	AddBook(book Book)
	RemoveBook(title string)
	ListBooks()
}

type MyLibrary struct {
	books []Book
}

//add book method
func (lib *MyLibrary) AddBook(book Book) {
	lib.books = append(lib.books, book)
}

//remove book method

func (lib *MyLibrary) RemoveBook(title string) {
	for i, book := range lib.books {
		if book.Title == title {
			lib.books = append(lib.books[0:i], lib.books[i+1:]...)
			break
		}
	}
}

func (lib *MyLibrary) ListBooks() {
	if len(lib.books) == 0 {
		fmt.Println("no books available")
		return
	}
	fmt.Println("list of available books: ")
	for _, book := range lib.books {
		fmt.Printf("Title: %s, Author: %s\n", book.Title, book.Author)
	}
}

func addBook(lib Library) {
	var title, author string
	fmt.Println("Enter Title: ")
	fmt.Scanln(&title)
	fmt.Println("Enter Author: ")
	fmt.Scanln(&author)
	lib.AddBook(Book{Title: title, Author: author})
	fmt.Println("book added")
}

func removeBook(lib Library) {
	var title string
	fmt.Println("Enter Title you want to remove from list: ")
	fmt.Scanln(&title)
	lib.RemoveBook(title)
	fmt.Println("book is removed")
}

func showMenu() {
	fmt.Print("please select an option: \n1: add book\n 2:remove book\n 3:list books\n 4:exit\n ")
}

func main() {
	library := &MyLibrary{}
	for {
		showMenu()
		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			addBook(library)
		case 2:
			removeBook(library)
		case 3:
			library.ListBooks()
		case 4:
			return
		default:
			fmt.Println("error!!!")
		}
	}
}
