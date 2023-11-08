package main

import (
	"fmt"
	"strconv"
)

func SetBit(value int64, index int) int64 {
	length := len(strconv.FormatInt(value, 2))
	var newValue int64 = 1 << (length - index)
	value |= newValue

	return value
}

func ClearBit(value int64, index int) int64 {
	length := len(strconv.FormatInt(value, 2))
	var newValue int64 = ^(1 << (length - index))
	value &= newValue

	return value
}

func main() {

	var value int64
	var setBit int
	var clearBit int

	value = 113235252362
	setBit = 3
	clearBit = 2

	fmt.Println(strconv.FormatInt(value, 2))
	fmt.Println(strconv.FormatInt(SetBit(value, setBit), 2))
	fmt.Println(strconv.FormatInt(ClearBit(value, clearBit), 2))

}
