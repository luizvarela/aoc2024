package day1

import (
	"fmt"

	"github.com/luizvarela/aoc2024/utils"
)

func Part1() {
	leftList, rightList, err := utils.ReadListsFromFile("days/day1/data/day_1_part_1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	totalDistance := calculateTotalDistance(leftList, rightList)
	fmt.Printf("Day 1 Part 1 - Total distance: %d\n", totalDistance)
}

func Part2() {
	leftList, rightList, err := utils.ReadListsFromFile("days/day1/data/day_1_part_1.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	totalSimilarityScore := calculateTotalSimilarityScore(leftList, rightList)
	fmt.Printf("Day 1 Part 2 - Total similarity score: %d\n", totalSimilarityScore)
}
