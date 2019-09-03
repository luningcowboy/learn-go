package main

import "fmt"

func main() {
	var a int = 10
	for a < 20 {
		fmt.Printf("a = %d\n", a)
		a++
		if a > 15 {
			break
		}
	}
}
