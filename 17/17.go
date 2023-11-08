package main

import (
	"fmt"
	"sort"
)

func main() {

	var lookup int = 62473
	massive := []int{6436, 36, 34, 7626237, 247426, 1, 62473, 1, 63, 66, 262}
	sort.Ints(massive)

	// binary search
	index := sort.Search(len(massive), func(i int) bool {
		return massive[i] == lookup
	})

	if index < len(massive) && massive[index] == lookup {
		fmt.Println("found")
	} else {
		fmt.Println("not found")
	}
}
