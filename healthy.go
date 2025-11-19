package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	if !strings.Contains(input, "G") {
		fmt.Println("nakhor lite")
	} else if strings.Count(input, "Y")+2*strings.Count(input, "R") > 5 {
		fmt.Println("nakhor lite")
	} else {
		fmt.Println("rahat baash")
	}
}
