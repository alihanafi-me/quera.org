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

	if count > 100 {
		fmt.Println("Steam")
	} else if count < 0 {
		fmt.Println("Ice")
	} else {
		fmt.Println("Water")
	}
}
