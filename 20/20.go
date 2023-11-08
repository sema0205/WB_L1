package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {

	var input string
	var answer string
	buffer := bufio.NewReader(os.Stdin)
	input, _ = buffer.ReadString('\n')

	slice := strings.Fields(input)
	slices.Reverse(slice)
	for _, s := range slice {
		answer += s
		answer += " "
	}

	fmt.Println(answer)
}
