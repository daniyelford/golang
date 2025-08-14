package main

import "fmt"

/*
در Go، interface فقط وقتی لازم میشه که بخوای چند struct مختلف یک قرارداد (contract) مشترک رو پیاده‌سازی کنن یا بخوای polymorphism داشته باشی
Go Implicit Interface داره، نه Explicit
type interface_name interface{
function_name(list of argument types)(list of return types)
}
*/

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreeting() string {
	return "hello welcome to bot"
}

func (spanishBot) getGreeting() string {
	return "hola"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {

	englishPart := englishBot{}
	spanishPart := spanishBot{}

	printGreeting(spanishPart)
	printGreeting(englishPart)
}
