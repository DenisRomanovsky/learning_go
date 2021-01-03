package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//Some VERY custom to eb logic

	return "Hi There!"
}

func (spanishBot) getGreeting() string {
	//Some VERY custom to sp logic

	return "Hola!"
}
