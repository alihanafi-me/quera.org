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

	fmt.Println((count * (count + 1)) / 2)
}
