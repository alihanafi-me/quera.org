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

	a, _ := strconv.Atoi(fields[0])
	b, _ := strconv.Atoi(fields[1])
	l, _ := strconv.Atoi(fields[2])
	y := l / 2
	x := l - y

	fmt.Println((x * a) + (y * b))
}
