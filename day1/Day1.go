package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("day1_input.txt")
	if err != nil {
		panic(err)
	}

	saveIndex := 50
	zeroCounter := 0

	text := string(file)

	lines := strings.Split(text, "\n")

	for _, line := range lines {
		direction := line[0]
		steps, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		if direction == 'R' {
			for steps > 0 {
				saveIndex = saveIndex + 1
				saveIndex = saveIndex % 100
				steps = steps - 1
				if saveIndex == 0 {
					zeroCounter++
				}

			}
		}
		if direction == 'L' {
			for steps > 0 {
				saveIndex = saveIndex - 1
				saveIndex = saveIndex % 100
				steps = steps - 1
				if saveIndex == 0 {
					zeroCounter++
				}

			}
		}

	}

	fmt.Println(zeroCounter)

}
