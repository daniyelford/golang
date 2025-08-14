package main // every program is part of a package. We define this using the package keyword

import "fmt" // let us import files included in the fmt package .The `fmt` package contains functions for formatted input and output

func main() { // This function is where the program's execution begins

	fmt.Println("hello world") //Println function is used for printing arguments  in a new line

}

/*
go run main.go


go build -o hello
./hello


go mod init myproject
go get github.com/gin-gonic/gin
go mod tidy


go fmt ./...
go vet ./...
go test -v


*/
