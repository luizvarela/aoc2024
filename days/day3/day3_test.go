package day3_test

import (
	"testing"

	"github.com/go-playground/assert"
	"github.com/luizvarela/aoc2024/days/day3"
)

func TestParseMulExpressions(t *testing.T) {
	line := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	sum := day3.ParseMulExpressions(line)

	assert.Equal(t, 161, sum)
}
