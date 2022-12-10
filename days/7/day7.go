package main

import (
	"AdventOfCode/lib"
	"fmt"
	"strings"
)

func main() {
	lines := lib.ParseInputFile("/home/lenak/Dokumente/Coding/Go/AdventOfCode/assets/day7/input.txt")

	root := New()

	var currentDir *Directory
	//totalSize := 0

Outer:
	for _, line := range lines {

		if strings.HasPrefix(line, "$ cd ") {
			parsedDir := strings.TrimPrefix(line, "$ cd ")
			if parsedDir == "/" {
				currentDir = root
			} else if parsedDir == ".." {
				currentDir = currentDir.Parent
			} else {
				if found, dir := currentDir.ContainsChild(parsedDir); found {
					currentDir = dir
				} else {

					panic("Directory not found")
				}
			}
			continue Outer
		}

		if strings.HasPrefix(line, "$ ls") {
			continue Outer
		}

		if strings.HasPrefix(line, "dir") {
			parsedDir := strings.TrimPrefix(line, "dir ")
			currentDir = currentDir.InsertDirectory(parsedDir, currentDir)
		} else {
			parsedFile := strings.Split(line, " ")
			currentDir.InsertFile(parsedFile[1], lib.StringToInt(parsedFile[0]))
		}

	}

	totalSize := 0

	for _, el := range root.Children {

		if el.Size() < 100000 {
			totalSize += el.Size()
		}
	}

	fmt.Println(totalSize)
}
