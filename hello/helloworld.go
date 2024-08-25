package main

import (
	"experiments/greetings"
	"fmt"
	"log"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello go")
	fmt.Println(quote.Go())
	fmt.Println(greetings.Hello("Atul"))

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)

}
