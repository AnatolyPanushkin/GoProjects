package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var number string

	for true {
		fmt.Print("Input number: ")
		fmt.Fscan(os.Stdin, &number)

		check, err := strconv.Atoi(number)

		if err != nil || check < 0 || check%10 == 0 || check/1000 > 1 || check/100 < 1 {
			fmt.Println("invalid format of number")
			continue
			// panic("invalid format of number")
		}

		result := string(number[2]) + string(number[1]) + string(number[0])

		fmt.Println(result)
	}
}
