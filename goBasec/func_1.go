package main

import "fmt"

// func swap(a int, b int) (int, n int) {
// 	return b, a
// }
func main() {
	// a := 1
	// b := 2
	// a, b = swap(a, b)
	// fmt.Printf("%d %d", a, b)
	defer func() {
		fmt.Println("After defer")
	}()
	fmt.Println("befor defer")
}
