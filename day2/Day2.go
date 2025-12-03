package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.ReadFile("day2_input.txt")
	if err != nil {
		panic(err)
	}

	text := string(file)
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		fmt.Println(line)
	}

}
