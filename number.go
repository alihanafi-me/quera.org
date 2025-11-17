package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	for _, item := range input {
		value := string(item)
		num, _ := strconv.Atoi(value)
		fmt.Print(value + ": ")

		for i := 0; i < num; i++ {
			fmt.Print(value)
		}

		fmt.Println()
	}
}
