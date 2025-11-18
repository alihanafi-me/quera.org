package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fields := strings.Fields(input)
	valid := []int{1, 1, 2, 2, 2, 8}
	num := 0

	for i, _ := range valid {
		num, _ = strconv.Atoi(fields[i])
		fmt.Print(valid[i] - num)
		fmt.Print(" ")
	}

	fmt.Println()
}
