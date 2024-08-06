package main

import (
	"fmt"

	"github.com/myuser/calculator"
	"rsc.io/quote"
)

func main() {
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version: ", calculator.Version)
	fmt.Println(quote.Hello())
}
