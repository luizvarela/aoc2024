package day2

func IsSafe(list []int) bool {
	return IsValidDistance(list) && (IsIncreasing(list) || IsDecreasing(list))
}

func IsSafePart2(list []int) bool {
	// First check if it's already safe without modifications
	if IsSafe(list) {
		return true
	}

	// Try removing each element one at a time
	for i := 0; i < len(list); i++ {
		// Create new list without current element
		newList := append(append([]int{}, list[:i]...), list[i+1:]...)
		if IsSafe(newList) {
			return true
		}
	}
	return false
}

func IsIncreasing(list []int) bool {
	for i := 1; i < len(list); i++ {
		if list[i] <= list[i-1] {
			return false
		}
	}
	return true
}

func IsDecreasing(list []int) bool {
	for i := 1; i < len(list); i++ {
		if list[i] >= list[i-1] {
			return false
		}
	}
	return true
}

func IsValidDistance(list []int) bool {
	for i := 1; i < len(list); i++ {
		if abs(list[i]-list[i-1]) > 3 {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
