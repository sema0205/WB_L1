package main

import (
	"fmt"
	"math/big"
)

func main() {

	firstBig := new(big.Int)
	firstBig.SetString("9185196812063415983469346515349810032", 10)
	secondBig := new(big.Int)
	secondBig.SetString("3269938619361938619361936130619968161", 10)

	multiply := new(big.Int)
	multiply.Mul(firstBig, secondBig)
	fmt.Println("multiply", multiply)

	divide := new(big.Int)
	divide.Div(firstBig, secondBig)
	fmt.Println("divide", divide)

	add := new(big.Int)
	add.Add(firstBig, secondBig)
	fmt.Println("add", add)

	substract := new(big.Int)
	substract.Sub(firstBig, secondBig)
	fmt.Println("substract", substract)
}
