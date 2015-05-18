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

func factorial(n uint64) uint64 {
	if n < 2 {
		return 1
	}
	var res uint64 = 2
	for ; n > 2; n-- {
		res *= n
	}
	return res

}

func main() {
	var exp uint
	for exp = 30; exp < 34; exp++ {
		fmt.Println(exp, "shift:", shift(exp), "pow:", pow(2, exp))
	}
	for exp = 62; exp < 66; exp++ {
		fmt.Println(exp, "shift:", shift(exp), "pow:", pow(2, exp))
	}
	var prev uint64 = 0;
	for n := uint64(0); n < 100; n++ {
		fmt.Println(n, "!", factorial(n))
		if factorial(n) < prev {
			break
		}
		prev = factorial(n)
	}
}
