package main

import "fmt"

func main() {
	pi := 3.14

	radius := 5

	circumference := 2 * pi * float64(radius)

	fmt.Println(circumference)

	fmt.Printf("For a raius of %v, the circle circumference is %.2f.", radius, circumference)
}
