package main

import "fmt"

func main() {
	const (
		a = iota // 0
		b // 1
		c // 2
		d = "ha" // ha iota += 1
		e // ha iota += 1
		f = 100 // 100 iota += 1
		g // 100 iota += 1
		h = iota // 7
		i // 8
	)
	fmt.Println(a,b,c,d,e,f,g,h,i)
	//0 1 2 ha ha 100 100 7 8
}
