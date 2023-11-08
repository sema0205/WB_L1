package main

import "fmt"

func main() {

	slices := []int{1, 23, 42, 52, 362, 62}
	var index int = 4

	slices = append(slices[:index], slices[index+1:]...)

	fmt.Println(slices)
}
