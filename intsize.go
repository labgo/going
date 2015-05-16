package main

import "fmt"

func main() {
	var i int
	var bits uint
	for bits = 30; bits < 34; bits++ {
		i = 1 << bits
		fmt.Println(bits, "->", i)
	}
}
