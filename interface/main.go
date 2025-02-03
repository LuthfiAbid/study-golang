package main

import "fmt"

type bot interface {
	getGreeting() string
}
type engBot struct{}
type inaBot struct{}

func main() {
	eb := engBot{}
	ib := inaBot{}

	printGreeting(eb)
	printGreeting(ib)
}

func (engBot) getGreeting() string {
	// For generating greeting word in english
	return "Hello there!"
}

func (inaBot) getGreeting() string {
	// For generating greeting word in indonesia
	return "Halo semua!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
