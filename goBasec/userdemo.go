package main

import "fmt"

type Person struct {
	name string
	age  int
}
type Student struct {
	Person     //匿名字段,使用和类包含类似
	speciality string
}

func main() {
	person := Person{"zhangsan", 25}
	// person2 :=Person{age:25,name:"wangwu"}//指定字段赋值
	fmt.Printf("%v", person)
}
