package main

import "fmt"

type Point struct {
	x int
	y int
}

func (p *Point) MoveBy(first int, second int) {
	p.x += first
	p.y += second
}

func main() {

	point := &Point{13, -42}
	pointSecond := &Point{4, 14}
	point.MoveBy(pointSecond.x, pointSecond.y)

	fmt.Println(point.x, point.y)
}
