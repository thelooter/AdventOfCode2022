package main

import (
	"AdventOfCode/lib"
	"fmt"
)

func main() {

	lines := lib.ParseInputFile("/home/lenak/Dokumente/Coding/Go/AdventOfCode/assets/day3/input.txt")

	var compartments [][2]string

	for _, line := range lines {
		var length = len(line)
		var first = line[:length/2]
		var second = line[length/2:]
		compartments = append(compartments, [2]string{first, second})
	}

	var sameChars []string

NEXT:
	for _, element := range compartments {
		for _, character := range element[0] {
			for _, secondCharacter := range element[1] {
				if character == secondCharacter {
					sameChars = append(sameChars, string(character))
					continue NEXT
				}
			}
		}
	}

	priorities := make(map[string]int)

	lowercase := "abcdefghijklmnopqrstuvwxyz"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i, letter := range lowercase {
		priorities[string(letter)] = i + 1
	}

	for i, letter := range uppercase {
		priorities[string(letter)] = i + 27
	}

	var sum = 0

	for _, element := range sameChars {
		sum += priorities[element]
	}

	fmt.Println(sum)

	//-------------------------------------- Task 2 --------------------------------------

	var groups [][3]string

	var counter = 0
	var currentGroupIndex = 0
	for _, line := range lines {
		if counter == 0 {
			groups = append(groups, [3]string{line, "", ""})
			counter++
		} else if counter == 1 {
			groups[currentGroupIndex][1] = line
			counter++
		} else if counter == 2 {
			groups[currentGroupIndex][2] = line
			counter = 0
			currentGroupIndex++
		}
	}

	var sameGroupChars []string

OUTER:
	for _, element := range groups {
		for _, character := range element[0] {
			for _, character2 := range element[1] {
				if character == character2 {
					for _, character3 := range element[2] {
						if character == character3 {
							sameGroupChars = append(sameGroupChars, string(character))
							continue OUTER
						}
					}
				}
			}
		}
	}

	var groupSum = 0

	for _, element := range sameGroupChars {
		groupSum += priorities[element]
	}

	fmt.Println(groupSum)

}
