package main

import (
	"experiments/greetings"
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello go")
	fmt.Println(quote.Go())
	fmt.Println(greetings.Hello("Atul"))
	fmt.Println(greetings.Hello(""))

}
