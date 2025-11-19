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

	scanner.Scan()
	input1 := scanner.Text()

	scanner.Scan()
	input2 := scanner.Text()

	errorCount := 0

	for i := 0; i < count; i++ {
		if input1[i] != input2[i] {
			errorCount++
		}
	}

	fmt.Println(errorCount)
}
