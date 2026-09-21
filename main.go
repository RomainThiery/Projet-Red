package main

import (
	"os"

	"fmt"
)

func main() {
	args := os.Args[1:]
	for _, arg := range args {
		for _, char := range arg {
			fmt.Print(char)
		}
		fmt.Print('\n')
	}
}
