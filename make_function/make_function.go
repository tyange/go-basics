package main

import "fmt"

func main() {
	// slice type, slots, allocated spaces
	userNames := make([]string, 2, 5)

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	counseRatings := make(map[string]float64, 3)

	counseRatings["go"] = 4.7
	counseRatings["react"] = 4.8
	counseRatings["angular"] = 4.5

	fmt.Println(counseRatings)
}
