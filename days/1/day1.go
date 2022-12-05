package main

import (
	"AdventOfCode/lib"
	"fmt"
	"sort"
	"strconv"
)

func main() {

	lines := lib.ParseInputFile("/home/lenak/Dokumente/Coding/Go/AdventOfCode/assets/day1/input.txt")

	var elves []int
	var current = 0

	for _, element := range lines {
		value, err := strconv.Atoi(element)

		if err != nil {
			elves = append(elves, current)
			current = 0
		}

		current += value
	}

	sort.Ints(elves)
	reversed := reverse(elves)
	fmt.Println("First: ", reversed[0])
	fmt.Println("First 3 Elves combined: ", reversed[0]+reversed[1]+reversed[2])

}

func reverse(input []int) []int {
	inputLen := len(input)
	out := make([]int, inputLen)

	for i, n := range input {
		j := inputLen - i - 1
		out[j] = n
	}

	return out
}
