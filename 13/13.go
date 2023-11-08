package main

import "fmt"

func swap(first *int, second *int) {

	*first = *first + *second
	*second = *first - *second
	*first = *first - *second
}

func main() {

	firstValue := 145
	secondValue := 632

	var first *int = &firstValue
	var second *int = &secondValue
	swap(first, second)
	fmt.Println(*first, *second)
}
