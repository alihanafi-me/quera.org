package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	result := "YES"

	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			result = "NO"
			break
		}
	}

	fmt.Println(result)
}
