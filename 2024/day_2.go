package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// ...
func d2_check(nums []int) int {
	prev := 0
	prev_dist := 0

	for _, n := range nums {
		if prev == 0 {
			prev = n
			continue
		}

		dist := prev - n
		if absInt(dist) > 3 || dist == 0 {
			return 0
		} else if dist > 0 && prev_dist < 0 || dist < 0 && prev_dist > 0 {
			return 0
		}

		prev_dist = dist
		prev = n
	}

	return 1
}

func getInts(lines []string) [][]int {
	var inner []int
	var outter [][]int

	for _, s := range lines {
		for _, n := range strings.Split(s, " ") {
			num, _ := strconv.Atoi(n)
			inner = append(inner, num)
		}
		outter = append(outter, inner)
		inner = []int{}
	}

	return outter
}

func d2() {
	data, err := os.ReadFile("inputs/input_2.txt")
	// data, err := os.ReadFile("EXAMPLE.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\r\n")
	new := getInts(lines)

	p1 := 0
	p2 := 0

	for _, n := range new {
		p1 += d2_check(n)
	}

	for _, n := range new {
		if d2_check(n) == 1 {
			p2 += 1
			continue
		}

		for i := range len(n) {
			thonk := slices.Concat(n[0:i], n[i+1:])

			if d2_check(thonk) == 1 {
				p2 += 1
				break
			}
		}
	}

	fmt.Printf("DAY TWO TOTAL_ONE: %d\n", p1)
	fmt.Printf("DAY TWO TOTAL_TWO: %d\n", p2)
}
