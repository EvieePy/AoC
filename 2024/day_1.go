package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func d1_splatter(data []string) ([]int, []int) {
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

func d1() {
	total_one := 0
	total_two := 0

	data, err := os.ReadFile("inputs/input_1.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\r\n")

	left, right := d1_splatter(lines)
	slices.Sort(left)
	slices.Sort(right)

	// left minus right (abs)....
	for i, n := range left {
		total_one += absInt(right[i] - n)
	}

	fmt.Printf("DAY ONE TOTAL_ONE: %d\n", total_one)

	// PART 2:
	mapping := make(map[int]int)
	for _, num := range right {
		mapping[num] = mapping[num] + 1
	}

	for _, num := range left {
		total_two += (num * mapping[num])
	}

	fmt.Printf("DAY ONE TOTAL_TWO: %d\n", total_two)
}
