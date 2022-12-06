package main

import (
	"AdventOfCode/lib"
	"fmt"
)

func main() {
	input := lib.ParseInputFile("/home/lenak/Dokumente/Coding/Go/AdventOfCode/assets/day6/input.txt")[0]

	var buffer []rune
	for _, char := range input {
		buffer = append(buffer, char)
	}

	var counter = 0

	for i := 4; i < len(input); i++ {
		prev4 := string(buffer[i-4 : i])

		var prevBuffer []string
		for _, char := range prev4 {
			prevBuffer = append(prevBuffer, string(char))
		}

		uniqueSet := unique(prevBuffer)
		if len(uniqueSet) == 4 {
			fmt.Println("Found 4 unique chars: ", prev4)
			fmt.Println("Index: ", i)
			break
		}

		counter++
	}

	//--------------Part 2------------------

	var counter2 = 0

	for i := 14; i < len(input); i++ {
		prev14 := string(buffer[i-14 : i])

		var prevBuffer []string
		for _, char := range prev14 {
			prevBuffer = append(prevBuffer, string(char))
		}

		uniqueSet := unique(prevBuffer)
		if len(uniqueSet) == 14 {
			fmt.Println("Found 14 unique chars: ", prev14)
			fmt.Println("Index: ", i)
			break
		}

		counter2++
	}
}

func unique(slice []string) []string {
	// create a map with all the values as key
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	// turn the map keys into a slice
	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}
