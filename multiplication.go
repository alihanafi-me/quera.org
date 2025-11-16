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
	count, _ := strconv.Atoi(scanner.Text())

	for i := 1; i <= count; i++ {
		for j := 1; j <= count; j++ {
			fmt.Print(i*j, " ")
		}
		fmt.Println()
	}
}
