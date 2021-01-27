package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, Go language")

	// Print another package
	fmt.Println(quote.Go())
}
