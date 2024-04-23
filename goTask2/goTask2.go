package main

import (
	"fmt"
	"os"
)

func main() {
	var a int
	var b int
	var c int

	fmt.Print("Input a: ")
	fmt.Fscan(os.Stdin, &a)

	fmt.Print("Input b: ")
	fmt.Fscan(os.Stdin, &b)

	fmt.Print("Input c: ")
	fmt.Fscan(os.Stdin, &c)

	if a > b && b > c {
		panic("invalid input")
	}

	if c*c == a*a+b*b {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}
