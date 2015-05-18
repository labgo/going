package main

import "fmt"

func shift(bits uint) int {
	return 1 << bits

}

func pow(n int, exp uint) int {
	if exp == 0 {
		return 1
	}
	var res int = n
	for ; exp > 1; exp-- {
		res *= n
	}
	return res
}

func main() {
	var exp uint
	for exp = 30; exp < 34; exp++ {
		fmt.Println(exp, "shift:", shift(exp), "pow:", pow(2, exp))
	}
}
