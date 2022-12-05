package main

import (
	"AdventOfCode/lib"
	"fmt"
	"strings"
)

type action struct {
	amount int
	from   int
	to     int
}

func main() {

	lines := lib.ParseInputFile("/home/lenak/Dokumente/Coding/Go/AdventOfCode/assets/day5/input.txt")

	crates := initCrates()

	var actions []action

	for _, line := range lines {
		parts := strings.Split(line, " ")

		actions = append(actions, action{
			amount: lib.StringToInt(parts[1]),
			from:   lib.StringToInt(parts[3]) - 1,
			to:     lib.StringToInt(parts[5]) - 1,
		})
	}

	for _, action := range actions {
		for i := 0; i < action.amount; i++ {
			fromCrate := crates[action.from]
			toCrate := crates[action.to]

			crates[action.from] = fromCrate[:len(fromCrate)-1]
			crates[action.to] = toCrate + string(fromCrate[len(fromCrate)-1])
		}
	}

	var topValues []string

	for _, crate := range crates {
		topValues = append(topValues, string(crate[len(crate)-1]))
	}

	fmt.Println("Solution Part 1", strings.Join(topValues, ""))

	//------------------- Part 2 -------------------

	crates = initCrates()

	for _, action := range actions {
		fromCrate := crates[action.from]
		toCrate := crates[action.to]

		itemsToMove := fromCrate[len(fromCrate)-action.amount:]

		crates[action.from] = fromCrate[:len(fromCrate)-action.amount]
		crates[action.to] = toCrate + itemsToMove
	}

	var topValuesTwo []string

	for _, crate := range crates {
		topValuesTwo = append(topValuesTwo, string(crate[len(crate)-1]))
	}

	fmt.Println("Solution Part 2", strings.Join(topValuesTwo, ""))

}

func initCrates() [9]string {
	return [9]string{
		"PFMQWGRT",
		"HFR",
		"PZRVGHSD",
		"QHPBFWG",
		"PSMJH",
		"MZTHSRPL",
		"PTHNML",
		"FDQR",
		"DSCNLPH",
	}
}
