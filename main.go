package main

import (
	"fmt"
)

var x, y int
var (
	a int
	b bool
)
var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g,h :=123,"hello"

func main() {
	fmt.Println("hello", "world")
	g, h := 123, "hello"
	//./main.go:21:6: aa declared and not used
	//var aa string = "abc"
	println(x, y, a, b, c, d, e, f, g, h)
	println(c, d)
	//c和d交换结果
	c, d = d, c
	println(c, d)

	const LENGTH = 10
	const WIDTH = 5
	var area int
	const l, m, n = 1, false, "str"
	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d", area)
	println()
	println(l, m, n)
}
