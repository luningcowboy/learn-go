package main

var x, y int
var ( // 这种方法一般用了申明全局变量
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"
// 这种不带声明的格式只能出现在函数体中
//g, h := 123, "hello"
func main() {
	g, h := 123, "hello"
	println(x, y, a, b, c, d, e, f, g, h)
}
