package main

import "fmt"

func isAnagram(x, y string) bool {
	if len(x) != len(y) {
		return false
	}

	mapper := make(map[rune]int)

	for _, char := range x {
		_, exists := mapper[char]
		if !exists {
			mapper[char] = 1
		} else {
			mapper[char]++
		}
	}

	for _, char := range y {
		_, exists := mapper[char]
		if !exists {
			return false
		} else {
			mapper[char]--
			updatedVal := mapper[char]
			if updatedVal < 0 {
				return false
			}
		}
	}

	for _, v := range mapper {
		if v != 0 {
			return false
		}
	}

	return true
}

func containsDuplicate(x []int) bool {
	mapper := make(map[int]bool)

	for _, val := range x {
		if mapper[val] {
			return true
		}
		mapper[val] = true
	}

	return false
}

func main() {
	result := isAnagram("aba", "aba")
	fmt.Printf("Result: %t\n", result)

	result = isAnagram("abba", "abba")
	fmt.Printf("Result: %t\n", result)

	result = isAnagram("xyz", "aaaaaaa")
	fmt.Printf("Result: %t\n", result)

	slice := []int{1, 2, 3, 3, 4, 5, 6}
	duplicateCheck := containsDuplicate(slice)
	fmt.Printf("contains duplicate: %t\n", duplicateCheck)

	//outras propriedades

	m := make(map[string]int)

	found := m["test"]
	fmt.Printf("Res: %d\n", found) //retorna 0

	val, exists := m["test"]
	fmt.Printf("val: %d, exists: %t\n\n", val, exists)

	m["test"] = 10
	found = m["test"]
	fmt.Printf("Res: %d\n", found) //retorna 10

	val, exists = m["test"]
	fmt.Printf("val: %d, exists: %t\n", val, exists)
}
