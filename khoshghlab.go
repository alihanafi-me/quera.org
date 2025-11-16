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

	for i := 0; i < count; i++ {
		fmt.Println("man khoshghlab hastam")
	}
}
