package main

import (
	"fmt"
	"learning/interfaces/builder/user"
	"learning/interfaces/singleton"
	"learning/interfaces/strategy/payment"
)

func runPixPayment(value float32) float32 {
	return payment.RunPayment(value, payment.PixPayment{})
}

func runCardPaymeny(value float32) float32 {
	return payment.RunPayment(value, payment.CardPayment{})
}

func useStrategy() {
	var value float32 = 100.00
	pixValue := runPixPayment(value)
	cardValue := runCardPaymeny(value)
	fmt.Printf("Strategy\n")
	fmt.Printf("Final PIX value: %.2f\n", pixValue)
	fmt.Printf("Final Card value: %.2f\n\n", cardValue)
}

func useBuilder() {
	builder := user.CreateBuilder()
	user := builder.SetAge(10).SetHeight(1.80).SetName("Reidner").Build()
	fmt.Printf("Builder User: %v\n\n", user)
}

func useSingleton() {
	user1 := singleton.GetObject()
	user2 := singleton.GetObject()
	user3 := singleton.GetObject()
	user4 := singleton.GetObject()

	fmt.Printf("Singleton\n")
	fmt.Printf("User1 addr.: %p\n", user1)
	fmt.Printf("User2 addr.: %p\n", user2)
	fmt.Printf("User3 addr.: %p\n", user3)
	fmt.Printf("User4 addr.: %p\n", user4)

	// user.SetAge(25)
	// user.SetName("Reidner")
	// fmt.Printf("Singleton User: %v\n", user)
}

func main() {
	useBuilder()
	useStrategy()
	useSingleton()
}
