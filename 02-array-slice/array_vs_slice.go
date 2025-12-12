package main

import (
	"fmt"
	"math/rand"
)

func modificaArray(array [20]int) [20]int {
	array[0] = 999
	return array
}

func modificaSlice(slice []int) []int {
	if len(slice) > 0 {
		slice[0] = 999
	} else {
		slice = append(slice, 999)
	}
	return slice
}

func populaSlice(slice []int) []int {
	for len(slice) < cap(slice) {
		slice = append(slice, rand.Intn(100))
	}
	return slice
}

func sliceInfo(slice []int) (int, int) {
	return cap(slice), len(slice)
}

func printSliceData(slice []int) {
	fmt.Println(slice)
	cap, len := sliceInfo(slice)
	msg := fmt.Sprintf("Slice info: %d, %d\n", cap, len)
	fmt.Println(msg)
}

func main() {
	const capacity = 20

	slice := make([]int, 0, capacity)
	printSliceData(slice)

	slice = modificaSlice(slice)
	printSliceData(slice)

	slice = populaSlice(slice)
	printSliceData(slice)

	var array [capacity]int
	array = modificaArray(array)

	fmt.Println(array)

	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d, ", slice[i])
	}

	for i := 0; i < capacity; i++ {
		fmt.Printf("%d, ", array[i])
	}
}
