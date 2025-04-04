package main

import (
	"fmt"

	calculator "github.com/mnogu/go-calculator"
)

func main() {
	val, err := calculator.Calculate("((2+2))*2")
	if err == nil {
		fmt.Print(val)
	}
}
