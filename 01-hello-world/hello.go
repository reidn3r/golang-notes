package main

import (
	"fmt"
	"math/rand"
)

func sum(a int, b int) int {
	return a + b
}

func multiply(a float64, b float64) (float64, bool) {
	var result float64 = a * b

	if result > 0 {
		return result, true
	}

	return result, false
}

func main() {
	var s int = sum(2, 3)
	var msg string = fmt.Sprintf("Result: %d\n", s)

	res, b := multiply(-1, 2)
	var msg2 string = fmt.Sprintf("Result: %.2f, Greather than zero: %t\n", res, b)

	fmt.Println("Hello, world")
	fmt.Printf(msg)
	fmt.Printf(msg2)

	name := func(name string) {
		fmt.Printf("Your name is: %s\n", name)
	}

	randomNUmber := func(maxNumber int) int {
		return rand.Intn(maxNumber)
	}

	name("Reidner")
	rnd := randomNUmber(1000)

	fmt.Printf("RANDOM: %d\n", rnd)
}
