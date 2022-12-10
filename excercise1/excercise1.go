package main

import "fmt"

func main() {
	firstName := "Tyange"
	var lastName = "Yu"

	fmt.Println(firstName)
	fmt.Println(lastName)

	currentYear := 2022
	var birthYear int
	birthYear = 1990

	age := currentYear - birthYear
	fmt.Println(age)

	currentYear = currentYear + 1
	age = currentYear - birthYear
	fmt.Println(age)

	var firstRune rune = '&'
	fmt.Println(firstRune)
	fmt.Println(string(firstRune))

	var firstByte byte = 'a'
	fmt.Println(firstByte)
}
