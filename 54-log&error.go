package main

import (
	"fmt"
	"log"
	"os"
)

func ReadFile(filename string) (string, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return " ", fmt.Errorf("file %s not found", filename)
	}
	x, err := os.ReadFile(filename)
	if err != nil {
		return " ", fmt.Errorf("operation failed")
	}
	return string(x), nil
}

func WriteFile(filename, data string) error {
	err := os.WriteFile(filename, []byte(data), 0666)
	if err != nil {
		return fmt.Errorf("failed to write data ")
	}
	return nil

}

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("failed to open log file", err)
	}
	log.SetOutput(file)

	filename := "abcd.txt"
	data := "salam"

	err = WriteFile(filename, data)
	if err != nil {
		log.Println("failed to write to file", err)
	} else {
		log.Println("file written successfully")
	}

	x, err := ReadFile(filename)
	if err != nil {
		log.Println("error reading file", err)
	} else {
		log.Println("file content:", x)
	}

}
