package main

import "fmt"

type Integer struct {
	value int
}

func (a Integer) compare(b Integer) bool {
	return a.value < b.value
}
func main() {
	// var a int =1
	// var b int =2
	a := Integer{1}
	b := Integer{2}
	fmt.Printf("%v", a.compare(b))
}
