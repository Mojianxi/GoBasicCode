package main

import "fmt"

type Point struct {
	px float32
	py float32
}

func (point *Point) setxy(px, py float32) {
	point.px = px
	point.py = py
}
func (point *Point) getxy() (float32, float32) {
	return point.px, point.py
}
func main() {
	point := new(Point)
	point.setxy(1.24, 4.25)
	px, py := point.getxy()
	fmt.Print(px, py)
}
