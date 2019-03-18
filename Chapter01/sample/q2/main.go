package main

import (
	"fmt"
	"strings"
)

func main() {
	var height int
	fmt.Scan(&height)

	fmt.Print(buildPyramid(height))
	fmt.Print(buildRPyramid(height))
}

func buildPyramid(height int) string {
	var sb strings.Builder
	for i := 1; i <= height; i++ {
		space := height - i
		asta := i*2 - 1
		sb.WriteString(createLine(" ", "*", space, asta) + "\n")
	}
	return sb.String()
}

func buildRPyramid(height int) string {
	var sb strings.Builder
	for i := 0; i < height; i++ {
		space := i
		asta := (height-i)*2 - 1
		text := fmt.Sprintf("%d", (i+1)%10)
		sb.WriteString(createLine(" ", text, space, asta) + "\n")
	}
	return sb.String()
}

func createLine(s1 string, s2 string, num1 int, num2 int) string {
	var sb strings.Builder
	for i := 0; i < num1; i++ {
		sb.WriteString(s1)
	}
	for i := 0; i < num2; i++ {
		sb.WriteString(s2)
	}
	return sb.String()
}
