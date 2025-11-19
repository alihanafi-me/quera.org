package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	days := []string{"shanbe", "1shanbe", "2shanbe", "3shanbe", "4shanbe", "5shanbe", "jome"}

	scanner.Scan()
	scanner.Scan()
	input1 := scanner.Text()
	fields1 := strings.Fields(input1)

	scanner.Scan()
	scanner.Scan()
	input2 := scanner.Text()
	fields2 := strings.Fields(input2)

	scanner.Scan()
	scanner.Scan()
	input3 := scanner.Text()
	fields3 := strings.Fields(input3)

	merged := append(append(fields1, fields2...), fields3...)

	var result []string

	for _, day := range days {
		for i := 0; i < len(merged); i++ {
			if merged[i] == day {
				break
			}

			if i == len(merged)-1 {
				result = append(result, day)
			}
		}
	}

	fmt.Println(len(result))
}
