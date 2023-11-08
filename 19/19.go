package main

import (
	"fmt"
)

func main() {

	var value string
	fmt.Scanln(&value)

	runes := []rune(value)
	runesSecond := make([]rune, len(runes))

	for i := 0; i < len(runes); i++ {
		runesSecond[len(runes)-i-1] = runes[i]
	}

	var answer string = string(runesSecond)
	fmt.Println(answer)
}
