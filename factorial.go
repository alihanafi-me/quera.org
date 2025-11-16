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
	result := 1

	for i := 1; i <= count; i++ {
		result *= i
	}

	fmt.Println(result)
}
