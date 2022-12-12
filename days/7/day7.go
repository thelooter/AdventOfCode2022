package main

import (
	"AdventOfCode/lib"
	"strconv"
	"strings"
)

var commands = lib.ParseInputFile("/home/lenak/Dokumente/Coding/Go/AdventOfCode/assets/day7/input.txt")
var index = 0

var total = 0
var dirSizes []int

func main() {
	rootSize := computeNextDirSize()
	missingSpace := 30_000_000 - (70_000_000 - rootSize)
	best := 10000000000

	for _, size := range dirSizes {
		if size > missingSpace && size < best {
			best = size
		}
	}

	println("Part 1:", total, "Part 2:", best)
}

func computeNextDirSize() int {
	size := 0
	for {
		command := peek()
		index += 1
		if command == "" {
			return size
		}

		if strings.HasPrefix(command, "$ cd ") {
			dirName := command[5:]
			if strings.HasPrefix(dirName, "..") {
				return size
			}

			dirSize := computeNextDirSize()
			dirSizes = append(dirSizes, dirSize)
			if dirSize <= 100_000 {
				total += dirSize
			}
			size += dirSize
		}

		if strings.HasPrefix(command, "$ ls") {
			size += computeFileListSize()
		}
	}
}

func computeFileListSize() int {
	size := 0
	for {
		command := peek()
		if command == "" || strings.HasPrefix(command, "$") {
			return size
		}

		index += 1
		if strings.HasPrefix(command, "dir ") {
			continue
		}

		parts := strings.Split(command, " ")
		file, _ := strconv.Atoi(parts[0])
		size += file
	}
}

func peek() string {
	if index >= len(commands) {
		return ""
	}
	return commands[index]
}
