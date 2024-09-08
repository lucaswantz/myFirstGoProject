package main

import (
	"fmt"
	"myFirstGoProject/pacote"
)

func main() {
	fmt.Println("Hello, World")
	fmt.Println(pacote.Bar)

	pacote.PrintMinha()
	digaOi()

	a, b := swap(1, 2)
	fmt.Println(a, b)

	fmt.Println(somar(10, 11, 12))
}

func digaOi() {
	fmt.Println("oi")
}

func swap(a, b int) (int, int) {
	return b, a
}

func divider1(a int, b int) (int, int) {
	res := a / b
	rem := a % b
	return res, rem
}

func divider2(a, b int) (res, rem int) {
	res = a / b
	rem = a % b
	return res, rem
}

func divider3(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return
}

func somar(nums ...int) int {
	var out int
	for _, n := range nums {
		out += n
	}

	return out
}
