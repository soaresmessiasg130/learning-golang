package main

import "fmt"
import "messiasspp.com/calculator"

func UseCalculator() {
	a64 := float64(12.99)
	b64 := float64(32.74)

	sum64 := currency.SumTwoCurrencysUsingFloat64(a64, b64)

	message64 := fmt.Sprintf("Using Float 64: sum of '%.16f' and '%.16f' is %.16f!", a64, b64, sum64)

	fmt.Println(message64)

	fmt.Printf("sads %.16f", float64(float64(1.99) + float64(2.74)))
}

func main() {
	fmt.Println("Hello from interface")

	UseCalculator()
}
