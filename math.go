package main

import "fmt"

func main() {
	resultado := Soma(1, 2)
	fmt.Println(resultado)
}

func Soma(a, b int) int {
	return a + b
}
