package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadListsFromFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var leftList, rightList []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var left, right int
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &left, &right)
		if err != nil {
			fmt.Println("Error parsing line:", err)
			continue
		}
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	return leftList, rightList, nil
}

func ReadLinesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
