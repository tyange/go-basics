package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	var years float64 = 10
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	fmt.Print(futureValue)
}
