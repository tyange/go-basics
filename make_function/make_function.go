package main

import "fmt"

func main() {
	// slice type, slots, allocated spaces
	userNames := make([]string, 2, 5)

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)
}
