package main

import "fmt"

func main() {
	// 배열의 크기를 명시하지 않으면, dynamic array.
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	// 아래처럼 값이 없는 자리를 조작할 수 없음.
	// price[2] = 10.11

	// 대신 append 메서드로 배열에 새로운 요소를 넣을 수 있음.
	// 그러나 append 메서드는 original array를 수정하지는 않고, 늘 새로운 배열을 반환함.
	// 그러므로 원래의 prices 배열을 재선언해주어야 함.
	prices = append(prices, 5.99)
	fmt.Println(prices)
}

// func main() {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	// including
// 	featuredPrices := prices[1:]
// 	// override original array
// 	featuredPrices[0] = 199.99
// 	// excluding
// 	highlightedPrices := featuredPrices[:1]

// 	fmt.Println(highlightedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))

// 	highlightedPrices = highlightedPrices[:3]
// 	fmt.Println(highlightedPrices)
// 	fmt.Println(len(highlightedPrices), cap(highlightedPrices))
// }
