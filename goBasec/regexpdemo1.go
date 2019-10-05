package main

import (
	"fmt"
	"regexp"
)

func main() {
	isok, _ := regexp.Match("[a-zA-Z]{3}", []byte("zhl"))
	fmt.Printf("%v", isok)
	reg1 := regexp.MustCompile(`^z(.*)l$`)
	result1 := reg1.FindAllString("zhangsan", -1)
	fmt.Printf("%v\n", result1)
	reg2 := regexp.MustCompile(`z(.{1})(.{1})(.*)l$`)
	result2 := reg2.FindAllStringSubmatch("zhangsan", -1)
	fmt.Printf("%v\n", result2)

}
