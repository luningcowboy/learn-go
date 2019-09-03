package main

import "fmt"
import "unsafe"
// 定义常量
const (
	d = "ddddd"
	e = len(d)
	f = unsafe.Sizeof(d)
)
func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "hello"

	area = LENGTH * WIDTH
	fmt.Printf("面积 : %d", area)
	println()
	println(a, b, c)
	println(d, e, f)
}
