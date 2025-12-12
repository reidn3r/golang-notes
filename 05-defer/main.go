package main

import "fmt"

func teste1() {
	defer fmt.Println("Ok1")
	defer fmt.Println("Ok2")
	defer fmt.Println("Ok3")
	panic("Teste")
	//Panic: Executa todos os defers registrados em ordem LIFO
	//Em seguida, mata o processo.

}

func teste2() {
	panic("Teste")
	defer fmt.Println("Ok")
}

func teste3() {
	defer fmt.Println("Ok1")
	defer fmt.Println("Ok2")
	defer fmt.Println("Ok3")
	//Executa em ordem LIFO
}

func main() {
	teste1()
	// teste2()
	teste3()
}
