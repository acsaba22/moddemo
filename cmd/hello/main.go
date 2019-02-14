package main

import (
	"fmt"

	"github.com/acsaba22/moddemo/greeter"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello " + greeter.Name())
	fmt.Println("V1: " + greeter.Talk() + quote.Hello())
}
