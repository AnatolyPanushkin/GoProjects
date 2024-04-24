package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string
	fmt.Fscan(os.Stdin, &input)

	if len(input) <= 0 || len(input) > 1000 {
		panic("invalid input lenght")
	}

	fmt.Println(strings.Join(strings.Split(input, ""), "*"))
}
