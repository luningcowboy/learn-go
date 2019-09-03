package main

import "fmt"

func main() {
	var a int = 10
	for a < 20 {
		if a == 15 {
			a += 1
			continue
		}
		fmt.Printf("a = %d\n", a)
		a++
	}
}
