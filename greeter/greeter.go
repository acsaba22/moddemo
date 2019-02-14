package greeter

import "rsc.io/quote"

func Name() string {
	return "Module"
}

func Talk() string {
	return quote.Hello()
}
