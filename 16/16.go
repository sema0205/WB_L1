package main

import (
	"fmt"
	"sort"
)

func main() {

	massive := []int{6436, 36, 34, 7626237, 247426, 1, 62473, 1, 63, 66, 262}

	// quicksort
	sort.Ints(massive)
	fmt.Println(massive)
}
