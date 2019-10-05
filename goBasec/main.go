package main

import "fmt"

// func main() {
// 	var x int
// 	x = 1
// 	//这种用法须在函数体内
// 	y, z := "go demo", 2
// 	fmt.Printf("%d %s %d", x, y, z)
// }
func main() {
	// sum := 0
	// x := 1
	// for x <= 100 {
	// 	sum += x
	// 	x++
	// }
	// fmt.Print(sum)
	// x:=1
	// switch x{
	// 	case 1:
	// 		fmt.Print("demo1")
	// 		//GO默认添加了break
	// 		break
	// 	case 2:
	// 	//继续需要添加
	// 	fallthrough
	// 	case 3:
	// 		fmt.Print("demo2")
	// 		break
	// 	default 3:
	// 		fmt.Print("demo3")
	// }
	// x := [5]int{1, 2, 3, 4, 5}
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	// x := make(map[string]int)
	// x["zhangsan"] = 78
	// x["lisi"] = 90
	// x["wangwu"] = 100
	// for i, v := range x {
	// 	fmt.Println(i, v)
	// }
	x := "zhangsan"
	for _, v := range x {
		fmt.Printf("%c\n", v)
	}
}
