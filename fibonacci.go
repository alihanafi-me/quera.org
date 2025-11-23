package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func inArray[T comparable](needle T, haystack []T) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	var result []int

	for i := 1; ; i++ {
		if i == 1 || i == 2 {
			result = append(result, 1)
		} else {
			result = append(result, result[i-2]+result[i-3])
		}

		if result[len(result)-1] >= count {
			break
		}
	}

	for i := 1; i <= count; i++ {
		if inArray(i, result) {
			fmt.Print("+")
		} else {
			fmt.Print("-")
		}
	}

	fmt.Println()
}
