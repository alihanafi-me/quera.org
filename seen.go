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
	sevenSeen := []string{"sabzeh", "senjed", "samanu", "serkeh", "seeb", "seer", "somagh"}

	for i := 0; i < count; i++ {
		fmt.Println(sevenSeen[i])
	}
}
