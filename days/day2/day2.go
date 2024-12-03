package day2

import (
	"fmt"

	"github.com/luizvarela/aoc2024/utils"
)

func Part1() {
	lines, err := utils.ReadLinesFromFile("days/day2/data/day_2_part_1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	safeReportsCount := 0
	for _, line := range lines {
		numbers := utils.SplitLine(line)
		intList := utils.ConvertStringListToIntList(numbers)
		if IsSafe(intList) {
			safeReportsCount++
		}
	}

	fmt.Printf("Day 2 Part 1 - Safe reports count: %d\n", safeReportsCount)
}

func Part2() {
	lines, err := utils.ReadLinesFromFile("days/day2/data/day_2_part_1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	safeReportsCount := 0
	for _, line := range lines {
		numbers := utils.SplitLine(line)
		intList := utils.ConvertStringListToIntList(numbers)
		if IsSafePart2(intList) {
			safeReportsCount++
		}
	}

	fmt.Printf("Day 2 Part 2 - Safe reports count: %d\n", safeReportsCount)
}
