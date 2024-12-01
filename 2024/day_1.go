package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func splatter(data []string) ([]int, []int) {
	var left []int
	var right []int

	for _, str := range data {
		value := strings.Split(str, "   ")

		leftN, _ := strconv.Atoi(value[0])
		rightN, _ := strconv.Atoi(value[1])

		left = append(left, leftN)
		right = append(right, rightN)
	}

	return left, right
}

func absInt(number int) int {
	if number < 0 {
		return 0 - number
	}
	return number
}

func main() {
	total_one := 0
	total_two := 0

	data, err := os.ReadFile("inputs/input_1.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\r\n")

	left, right := splatter(lines)
	slices.Sort(left)
	slices.Sort(right)

	// left minus right (abs)...
	for i, n := range left {
		total_one += absInt(right[i] - n)
	}

	fmt.Printf("TOTAL ONE: %d\n", total_one)

	// PART 2:
	mapping := make(map[int]int)
	for _, num := range right {
		mapping[num] = mapping[num] + 1
	}

	for _, num := range left {
		total_two += (num * mapping[num])
	}

	fmt.Printf("TOTAL TWO: %d\n", total_two)
}
