package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	fmt.Fscan(os.Stdin, &input)

	numbers := strings.Split(input, "")

	for i := range numbers {
		number, err := strconv.Atoi(numbers[i])

		if err != nil {
			panic("invalud number")
		}

		fmt.Print(number * number)
	}
}
