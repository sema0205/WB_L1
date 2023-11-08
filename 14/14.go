package main

import (
	"fmt"
	"reflect"
)

func main() {

	var intVal int = 4
	var stringVal string = "test"
	var boolVal bool = true
	var chanVal = make(chan int)

	var massive []interface{}
	massive = append(massive, intVal)
	massive = append(massive, stringVal)
	massive = append(massive, boolVal)
	massive = append(massive, chanVal)

	for _, v := range massive {
		fmt.Print(v, " ")
		fmt.Println(reflect.TypeOf(v).String())
	}

}
