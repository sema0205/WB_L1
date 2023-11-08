package main

import "fmt"

type MySet struct {
	storage map[string]bool
	set     []string
}

func (m *MySet) AddElement(value string) {
	if !m.storage[value] {
		m.storage[value] = true
		m.set = append(m.set, value)
	}
}

func (m *MySet) PrintElements() {
	fmt.Println(m.set)
}

func main() {

	massive := []string{"cat", "cat", "dog", "cat", "tree"}
	set := MySet{
		storage: make(map[string]bool),
		set:     []string{},
	}

	for _, val := range massive {
		set.AddElement(val)
	}

	set.PrintElements()
}
