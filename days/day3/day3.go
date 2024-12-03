package day3

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/luizvarela/aoc2024/utils"
)

func Part1() {
	lines, err := utils.ReadLines("days/day3/data/day3_part1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sum := 0
	for _, line := range lines {
		sum += ParseMulExpressions(line)
	}

	fmt.Printf("Day 3 Part 1 - Sum: %d\n", sum)
}

func ParseMulExpressions(input string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllStringSubmatch(input, -1)
	sum := 0

	// fmt.Println(matches)

	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		sum += x * y
	}

	return sum
}
