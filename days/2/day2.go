package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/home/lenak/Dokumente/Coding/Go/AdventOfCode/assets/day2/input.txt")
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

	var pairs []map[string]string
	pointsMap := map[string]int{
		"X": 1, //Rock
		"Y": 2, //Paper
		"Z": 3, //Scissors
	}

	for _, line := range lines {
		split := strings.Split(line, " ")
		pairs = append(pairs, map[string]string{
			split[0]: split[1],
		})
	}

	var points = 0

	for _, element := range pairs {
		for key, value := range element {
			switch key {
			case "A":
				switch value {
				case "X":
					points += pointsMap["X"]
					points += 3
				case "Y":
					points += pointsMap["Y"]
					points += 6
				case "Z":
					points += pointsMap["Z"]
					points += 0
				}
			case "B":
				switch value {
				case "X":
					points += pointsMap["X"]
					points += 0
				case "Y":
					points += pointsMap["Y"]
					points += 3
				case "Z":
					points += pointsMap["Z"]
					points += 6
				}
			case "C":
				switch value {
				case "X":
					points += pointsMap["X"]
					points += 6
				case "Y":
					points += pointsMap["Y"]
					points += 0
				case "Z":
					points += pointsMap["Z"]
					points += 3
				}
			}
		}
	}

	fmt.Println("First:", points)

	secondPoints := 0

	for _, element := range pairs {
		for key, value := range element {
			switch key {
			case "A": //Rocks
				switch value {
				case "X": //lose
					secondPoints += pointsMap["Z"]
					secondPoints += 0
				case "Y": //draw
					secondPoints += pointsMap["X"]
					secondPoints += 3
				case "Z": //win
					secondPoints += pointsMap["Y"]
					secondPoints += 6
				}
			case "B": //Paper
				switch value {
				case "X": //lose
					secondPoints += pointsMap["X"]
					secondPoints += 0
				case "Y": //draw
					secondPoints += pointsMap["Y"]
					secondPoints += 3
				case "Z": //win
					secondPoints += pointsMap["Z"]
					secondPoints += 6
				}
			case "C": //Scissors
				switch value {
				case "X": //lose
					secondPoints += pointsMap["Y"]
					secondPoints += 0
				case "Y": //draw
					secondPoints += pointsMap["Z"]
					secondPoints += 3
				case "Z": //win
					secondPoints += pointsMap["X"]
					secondPoints += 6
				}
			}
		}
	}

	fmt.Println("Second: ", secondPoints)
}
