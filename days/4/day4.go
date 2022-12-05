package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("/home/lenak/Dokumente/Coding/Go/AdventOfCode/assets/day4/input.txt")
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

	counter := 0

	for _, line := range lines {

		parts := strings.Split(line, ",")
		left, right := parts[0], parts[1]

		leftParts := strings.Split(left, "-")
		rightParts := strings.Split(right, "-")

		leftStart, leftEnd := toInt(leftParts[0]), toInt(leftParts[1])
		rightStart, rightEnd := toInt(rightParts[0]), toInt(rightParts[1])

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

		leftStart, leftEnd := toInt(leftParts[0]), toInt(leftParts[1])
		rightStart, rightEnd := toInt(rightParts[0]), toInt(rightParts[1])

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

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
