package utils

import (
	"strconv"
	"strings"
)

func ConvertStringListToIntList(list []string) []int {
	intList := make([]int, len(list))
	for i, value := range list {
		intList[i], _ = strconv.Atoi(value)
	}
	return intList
}

func SplitLine(line string) []string {
	return strings.Split(line, " ")
}

func ReadLines(filename string) ([]string, error) {
	return ReadLinesFromFile(filename)
}
