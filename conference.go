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

	row, _ := strconv.Atoi(fields[0])
	col, _ := strconv.Atoi(fields[1])

	if col < 11 {
		fmt.Println("Right " + strconv.Itoa(11-row) + " " + strconv.Itoa(col))
	} else {
		fmt.Println("Left " + strconv.Itoa(11-row) + " " + strconv.Itoa(21-col))
	}
}
