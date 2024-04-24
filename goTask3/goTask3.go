package main

import (
	"fmt"
	"os"
)

func main() {
	var seconds int

	fmt.Print("Input a: ")
	fmt.Fscan(os.Stdin, &seconds)

	hours := seconds / 3600
	minutes := (seconds - hours*3600) / 60

	fmt.Printf("It is %s hours %d minutes.", fmt.Sprint(hours), minutes)
}
