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
	count, _ := strconv.Atoi(scanner.Text())

	fmt.Println(strings.Repeat("*", count))

	for i := 0; i < count-2; i++ {
		fmt.Println("*" + strings.Repeat(" ", count-2) + "*")
	}

	fmt.Println(strings.Repeat("*", count))
}
