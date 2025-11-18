package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func reverse(str string) string {
	strBytes := []byte(str)

	for i, j := 0, 0; i < len(strBytes)/2; i++ {
		j = len(strBytes) - 1 - i
		strBytes[i], strBytes[j] = strBytes[j], strBytes[i]
	}

	return string(strBytes)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	first := scanner.Text()
	modFirst, _ := strconv.Atoi(reverse(first + "1"))
	scanner.Scan()
	second := scanner.Text()
	modSecond, _ := strconv.Atoi(reverse(second + "1"))

	if modSecond < modFirst {
		fmt.Println(second + " < " + first)
	} else if modFirst < modSecond {
		fmt.Println(first + " < " + second)
	} else {
		fmt.Println(first + " = " + second)
	}
}
