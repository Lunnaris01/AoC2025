package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.ReadFile("day2_input.txt")
	if err != nil {
		panic(err)
	}

	text := string(file)
	lines := strings.Split(text, ",")
	wrongIDSum := 0

	for _, line := range lines {
		values := strings.Split(line, "-")
		start, _ := strconv.Atoi(values[0])
		end, _ := strconv.Atoi(values[1])
		for i := start; i <= end; i++ {
			numString := strconv.Itoa(i)
			numStringLen := len(numString)
			if numStringLen%2 != 0 {
				continue
			}
			if numString[0:numStringLen/2] == numString[numStringLen/2:] {
				wrongIDSum += i

			}

		}

	}

	fmt.Println(wrongIDSum)

}
