package main

import (
	"fmt"
	"log"
)

func WithLogging(f func()) func() {
	return func() {
		log.Println("Entering function")
		defer log.Println("Exiting function")
		f()
	}
}

func ExampleFunction() {
	fmt.Println("Inside the example function.")
}

func AnotherFunction() {
	fmt.Println("Inside another function.")
}

func main() {
	decoratedFunc1 := WithLogging(ExampleFunction)
	decoratedFunc2 := WithLogging(AnotherFunction)

	fmt.Println("Calling decoratedFunc1:")
	decoratedFunc1()

	fmt.Println("\nCalling decoratedFunc2:")
	decoratedFunc2()
}
