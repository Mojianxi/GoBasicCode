package main

import "fmt"

type Animal interface {
	Fly()
	Run()
}
type Bird struct {
}

func (bird Bird) Fly() {
	fmt.Println("bird is flying")
}
func (bird Bird) Run() {
	fmt.Println("bird is Run")
}
func main() {
	var animal Animal
	bird := new(Bird)
	animal = bird
	animal.Fly()
	animal.Run()
	var v1 interface{}
	v1 = 6.78
	fmt.Print(v1)
	if v, ok := v1.(float64); ok {
		fmt.Println(v, ok)
	}
}
