package main

import (
	"AdventOfCode/lib"
	"fmt"
	"strings"
)

func main() {

	lines := lib.ParseInputFile("/home/lenak/Dokumente/Coding/Go/AdventOfCode/assets/day4/input.txt")

	counter := 0

	for _, line := range lines {

		parts := strings.Split(line, ",")
		left, right := parts[0], parts[1]

		leftParts := strings.Split(left, "-")
		rightParts := strings.Split(right, "-")

		leftStart, leftEnd := lib.StringToInt(leftParts[0]), lib.StringToInt(leftParts[1])
		rightStart, rightEnd := lib.StringToInt(rightParts[0]), lib.StringToInt(rightParts[1])

		var numArray [2][]int

		for i := leftStart; i <= leftEnd; i++ {
			numArray[0] = append(numArray[0], i)
		}

		for i := rightStart; i <= rightEnd; i++ {
			numArray[1] = append(numArray[1], i)
		}

		if contains(numArray[0], numArray[1]) {
			counter++
		} else if contains(numArray[1], numArray[0]) {
			counter++
		}

	}

	fmt.Println("Contains all: ", counter)

	counter = 0

Outer:
	for _, line := range lines {

		parts := strings.Split(line, ",")
		left, right := parts[0], parts[1]

		leftParts := strings.Split(left, "-")
		rightParts := strings.Split(right, "-")

		leftStart, leftEnd := lib.StringToInt(leftParts[0]), lib.StringToInt(leftParts[1])
		rightStart, rightEnd := lib.StringToInt(rightParts[0]), lib.StringToInt(rightParts[1])

		for i := leftStart; i <= leftEnd; i++ {
			for j := rightStart; j <= rightEnd; j++ {
				if i == j {
					counter++
					continue Outer
				}
			}
		}
	}

	fmt.Println("Contains at least one: ", counter)

}

func contains(ints []int, ints2 []int) bool {
	containsAll := true
	for _, i := range ints {
		if !containsInt(ints2, i) {
			containsAll = false
		}
	}

	return containsAll
}

func containsInt(ints2 []int, i int) bool {
	for _, j := range ints2 {
		if i == j {
			return true
		}
	}
	return false
}
