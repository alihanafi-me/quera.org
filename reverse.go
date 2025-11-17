package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var nums []string
	var reverse []string

	scanner.Scan()
	input := scanner.Text()

	for input != "0" {
		nums = append(nums, input)
		scanner.Scan()
		input = scanner.Text()
	}

	for i := 0; i < len(nums); i++ {
		reverse = append(reverse, nums[len(nums)-i-1])
	}

	for _, num := range reverse {
		fmt.Println(num)
	}
}
