package day1

import "sort"

func calculateTotalDistance(leftList, rightList []int) int {
	sort.Ints(leftList)
	sort.Ints(rightList)

	totalDistance := 0
	for i := 0; i < len(leftList); i++ {
		distance := abs(leftList[i] - rightList[i])
		totalDistance += distance
	}

	return totalDistance
}

func calculateTotalSimilarityScore(leftList, rightList []int) int {
	totalSimilarityScore := 0
	for i := 0; i < len(leftList); i++ {
		totalSimilarityScore += leftList[i] * countOccurrences(rightList, leftList[i])
	}
	return totalSimilarityScore
}

func countOccurrences(list []int, target int) int {
	count := 0
	for _, value := range list {
		if value == target {
			count++
		}
	}
	return count
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
