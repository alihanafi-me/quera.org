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
	numbers := []int{0, 0, 0, 0}
	scanner.Scan()
	input := scanner.Text()
	fields := strings.Fields(input)

	numbers[0], _ = strconv.Atoi(fields[0])
	numbers[1], _ = strconv.Atoi(fields[1])
	numbers[2], _ = strconv.Atoi(fields[2])
	numbers[3], _ = strconv.Atoi(fields[3])

	if numbers[0] == numbers[2] {
		fmt.Println("Vertical")
	} else if numbers[1] == numbers[3] {
		fmt.Println("Horizontal")
	} else {
		fmt.Println("Try again")
	}
}
