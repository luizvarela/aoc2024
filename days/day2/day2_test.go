package day2_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/luizvarela/aoc2024/days/day2"
)

func TestIsSafe(t *testing.T) {
	assert.Equal(t, day2.IsSafe([]int{1, 2, 3, 4}), true)
	assert.Equal(t, day2.IsSafe([]int{4, 3, 2, 1}), true)
	assert.Equal(t, day2.IsSafe([]int{1, 3, 2, 4}), false)
	assert.Equal(t, day2.IsSafe([]int{7, 6, 4, 2, 1}), true)
	assert.Equal(t, day2.IsSafe([]int{1, 2, 7, 8, 9}), false)
	assert.Equal(t, day2.IsSafe([]int{9, 7, 6, 2, 1}), false)
	assert.Equal(t, day2.IsSafe([]int{1, 3, 2, 4, 5}), false)
	assert.Equal(t, day2.IsSafe([]int{8, 6, 4, 4, 1}), false)
	assert.Equal(t, day2.IsSafe([]int{1, 3, 6, 7, 9}), true)
}

func TestIsValidDistance(t *testing.T) {
	assert.Equal(t, day2.IsValidDistance([]int{1, 2, 7, 8, 9}), false)
	assert.Equal(t, day2.IsValidDistance([]int{1, 2, 3, 4, 5}), true)
}

func TestIsDecreasing(t *testing.T) {
	assert.Equal(t, day2.IsDecreasing([]int{1, 2, 7, 8, 9}), false)
	assert.Equal(t, day2.IsDecreasing([]int{5, 4, 3, 2, 1}), true)
}

func TestIsIncreasing(t *testing.T) {
	assert.Equal(t, day2.IsIncreasing([]int{1, 2, 7, 8, 9}), true)
	assert.Equal(t, day2.IsIncreasing([]int{5, 4, 3, 2, 1}), false)
}

func TestIsSafePart2(t *testing.T) {
	assert.Equal(t, day2.IsSafePart2([]int{7, 6, 4, 2, 1}), true)
	assert.Equal(t, day2.IsSafePart2([]int{1, 2, 7, 8, 9}), false)
	assert.Equal(t, day2.IsSafePart2([]int{9, 7, 6, 2, 1}), false)
	assert.Equal(t, day2.IsSafePart2([]int{1, 3, 2, 4, 5}), true)
	assert.Equal(t, day2.IsSafePart2([]int{8, 6, 4, 4, 1}), true)
	assert.Equal(t, day2.IsSafePart2([]int{1, 3, 6, 7, 9}), true)
}
