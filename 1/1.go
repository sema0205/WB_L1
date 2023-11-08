package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h Human) CountSomething() int {
	return 100 - h.age
}

type Action struct {
	Human
}

func main() {

	var action Action
	action.name = "name"
	action.age = 46
	fmt.Println(action.CountSomething())
}
