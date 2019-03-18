package main

import (
	"fmt"
	"strings"
)

func main() {
	var start, end int
	fmt.Scan(&start, &end)

	fmt.Print(fizzBuzz(start, end))
}

func fizzBuzz(start int, end int) string {
	var sb strings.Builder
	for i := start; i <= end; i++ {
		if i%15 == 0 {
			sb.WriteString("FizzBuzz\n")
		} else if i%3 == 0 {
			sb.WriteString("Fizz\n")
		} else if i%5 == 0 {
			sb.WriteString("Buzz\n")
		} else {
			sb.WriteString(fmt.Sprintf("%d\n", i))
		}
	}
	return sb.String()
}
