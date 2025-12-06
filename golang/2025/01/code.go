package main

import (
	"aoc-in-go/2025/utils"
	"math"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return calculateDialPassword2(input)
	}
	// solve part 1 here
	return calculateDialPassword(input)
}

func calculateDialPassword(input string) int {
	position := 50
	zeroCount := 0

	for _, rotation := range strings.Split(input, "\n") {
		// Parse direction and amount in one pass
		direction := rotation[0]
		amount, _ := strconv.Atoi(rotation[1:])

		// Apply rotation
		if direction == 'L' {
			position -= amount
		} else { // 'R'
			position += amount
		}

		// Wrap to 0-99 range
		position = ((position % 100) + 100) % 100

		// Count zeros
		if position == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

func calculateDialPassword2(input string) int {
	position := 50
	zeroCount := 0

	for _, rotation := range strings.Split(input, "\n") {
		// Parse direction and amount in one pass
		direction := rotation[0]
		amount, _ := strconv.Atoi(rotation[1:])

		// Apply rotation
		zeroesPassed := 0
		if direction == 'R' {
			zeroesPassed = int(math.Abs(float64(position+amount)) / 100)
			position = utils.Mod(position+amount, 100)
		} else {
			inverted := 0
			if position > 0 {
				inverted = (100 - position)
			}
			zeroesPassed = int(math.Abs(float64(inverted+amount)) / 100)
			position = utils.Mod(position-amount, 100)
		}

		zeroCount += zeroesPassed
	}

	return zeroCount
}
