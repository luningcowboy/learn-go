package main

func main() {
	println(max(1,2))
	a := "a"
	b := "b"
	a, b = swap(a, b)
	println(a, b)
	c := 100
	d := 200
	println(c, d)
	swapInt(&c, &d)
	println(c, d)
}
/* 返回两个数中的最大值 */
func max(n1, n2 int) int {
	var ret int
	if n1 > n2 {
		ret = n1
	} else{
		ret = n2
	}
	return ret
}
/* 交换两个值 */
func swap(x, y string) (string, string) {
	return y, x
}
/* 引用参数 */
func swapInt(x *int, y *int) {
	var tmp int
	tmp = *x
	*x = *y
	*y = tmp
}
