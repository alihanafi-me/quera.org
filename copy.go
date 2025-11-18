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

	count, _ := strconv.Atoi(fields[0])
	name := fields[1]

	fmt.Print(strings.Repeat("copy of ", count))
	fmt.Println(name)
}
