package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	// 1)
	hobbies := [3]string{"Came", "Saw", "Won"}
	fmt.Println(hobbies)

	// 2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3)
	// mainHobbies := hobbies[0:2]
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// 4)
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5)
	courseGoals := []string{"Learn Go!", "Learn all the basics"}
	fmt.Println(courseGoals)

	// 6)
	courseGoals[1] = "Learn all the details"
	courseGoals = append(courseGoals, "Learn all the basics")
	fmt.Println(courseGoals)

	// 7)
	products := []Product{{"first-product", "A First Product", 10.11}, {"second-product", "A Second Product", 11.11}}
	fmt.Println(products)

	newProduct := Product{"third-product", "3", 3.33}
	products = append(products, newProduct)
	fmt.Println(products)
}
