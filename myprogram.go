package main

import (
	"fmt"

	"go-coverage/greetings"
	"go-coverage/mymath"
	"rsc.io/quote"
)

func main() {
	fmt.Printf("I say %q and %q\n", quote.Hello(), greetings.Goodbye())

	// fmt.Printf("Is 1001 a big number? %v\n", mymath.IsBigNum(1001))
	// fmt.Printf("Is 999 a big number? %v\n", mymath.IsBigNum(999))

	fmt.Printf("Is 1001 a small number? %v\n", mymath.IsSmallNum(1001))
	fmt.Printf("Is 999 a small number? %v\n", mymath.IsSmallNum(999))

	fmt.Printf("...")
}
