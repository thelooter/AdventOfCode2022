package lib

import (
	"bufio"
	"os"
	"strconv"
)

func ParseInputFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func StringToInt(input string) int {
	atoi, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return atoi
}
