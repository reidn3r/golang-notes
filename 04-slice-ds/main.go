package main

import (
	"fmt"
	"math/rand"
)

func createRandintSlice(capacity uint) []int {
	slice := make([]int, 0, capacity)
	for len(slice) < cap(slice) {
		slice = append(slice, rand.Intn(10))
	}
	return slice
}

func somaPares(slice []int) int {
	sum := 0
	for idx, val := range slice {
		if val%2 == 0 {
			fmt.Printf("index: %d ", idx)
			sum += val
		}
	}
	return sum
}

func reverseString(input string) string {
	if len(input) == 0 {
		return ""
	}

	runeSlice := []rune(input)
	l, r := 0, len(input)-1
	for l < r {
		auxRune := runeSlice[l]
		runeSlice[l] = runeSlice[r]
		runeSlice[r] = auxRune

		l++
		r--
	}

	return string(runeSlice)
}

func main() {
	slice := createRandintSlice(10)
	fmt.Println(slice)
	pares := somaPares(slice)
	fmt.Printf("Pares: %d\n", pares)

	testString := "banana"
	revere := reverseString(testString)
	fmt.Printf("Original: %s, Reverse: %s\n", testString, revere)

	//pega apenas indice
	for i := range slice {
		fmt.Printf("index: %d, valor: %d\n", i, slice[i])
	}
}
