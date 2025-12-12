package main

import "fmt"

type User struct {
	name   string
	age    uint
	height float64
	weight float64
}

func localSwap(a, b int) {
	aux := a
	a = b
	b = aux
}

func refSwap(a, b *int) {
	aux := *a
	*a = *b
	*b = aux
}

// swap
func runSwap() {
	a, b := 10, 50
	fmt.Printf("Original Values: a: %d, b: %d\n", a, b)

	localSwap(a, b)
	fmt.Printf("Local Swap: a: %d, b: %d\n", a, b)

	refSwap(&a, &b)
	fmt.Printf("Ref Swap: a: %d, b: %d\n", a, b)
}

// struct + referencia/valor
func runStruct() {
	//new keyword -> retorna ponteiro!!
	user1 := new(User)
	user1.name = "Reidner Adnan"
	user1.age = 5
	user1.height = 1.95
	user1.weight = 100
	fmt.Printf("User1: %v\n", user1)

	//& keyword -> retorna ponteiro!!
	user2 := &User{name: "mock username 2", age: 10, height: 2.3, weight: 100.50}
	fmt.Printf("User2: %v\n", user2)

	//& keyword -> retorna valor!!
	user3 := User{name: "mock username 3", age: 10, height: 2.3, weight: 100.50}
	fmt.Printf("User3: %v\n", user3)
}

type Counter struct {
	Value uint
}

func (c *Counter) RefIncrement() {
	c.Value++
}

func (c Counter) LocalIncrement() {
	c.Value++
}

func main() {
	runSwap()
	runStruct()

	c1 := &Counter{Value: 0}
	c1.RefIncrement()
	c1.LocalIncrement()
	fmt.Printf("Counter 1 Value: %d\n", c1.Value)

	c2 := &Counter{Value: 0}
	c2.LocalIncrement()
	fmt.Printf("Counter 2 Value: %d\n", c2.Value)
}
