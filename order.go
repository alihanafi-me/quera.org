package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	scanner.Scan()
	input := scanner.Text()
	fields := strings.Fields(input)

	for i := len(fields) - 1; i > -1; i-- {
		fmt.Print(fields[i] + " ")
	}

	fmt.Println()
}
