package main

import "fmt"

func main() {
	age := 32

	agePointer := &age // type with '*' (ex. *int): pointer type

	fmt.Println("Age:", *agePointer) // pointer type with '*': get real value

	editAgeToAdultAge(agePointer) // edit orginal value(age)
	fmt.Println(age)
}

func editAgeToAdultAge(age *int) {
	*age = *age - 18
}
