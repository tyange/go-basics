package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	// fmt.Println("Future Value:", futureValue)
	fmt.Printf("Future Value: %.0f\nFuture Value (adjusted for Infaltion): %.0f", futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for Infaltion):", futureRealValue)
}
