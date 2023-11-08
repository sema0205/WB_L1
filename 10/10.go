package main

import "fmt"

func main() {

	storage := make(map[int][]float64)
	massive := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	for _, value := range massive {
		index := int(value) - int(value)%10
		storage[index] = append(storage[index], value)
	}

	fmt.Println(storage)
}
