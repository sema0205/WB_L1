package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SameCharacters(value string) bool {
	var answer bool = true
	storage := make(map[string]bool)
	runes := []rune(value)

	for _, r := range runes {
		char := strings.ToLower(strconv.QuoteRune(r))
		if storage[char] {
			answer = false
			break
		}
		storage[char] = true
	}

	return answer
}

func main() {
	var input string
	fmt.Scanln(&input)
	fmt.Println(SameCharacters(input))
}
