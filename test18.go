package main

import "fmt"

func main() {
	var a int = 10
	LOOP: for a < 20 {
		if a == 15 {
			a += 1
			goto LOOP
		}
		fmt.Printf("a = %d\n", a)
		a++
	}
}
