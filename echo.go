package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := []rune(scanner.Text())

	fmt.Println(string(input))

	for i := 1; i < len(input); i++ {
		for j := 0; j < i; j++ {
			input[j] = input[i]
		}

		fmt.Println(string(input))
	}
}
