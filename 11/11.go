package main

import "fmt"

func main() {

	massive1 := []int{4, 5, 26, 27, 848, 273, 4573, 7}
	massive2 := []int{5135, 136, 23, 848, 235, 62, 23632, 27}

	var answer []int

	storage := make(map[int]int)

	for _, val := range massive1 {
		storage[val] += 1
	}

	for _, val := range massive2 {
		if storage[val] == 1 {
			answer = append(answer, val)
		}
	}

	fmt.Println(answer)
}
