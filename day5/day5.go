package day5

import (
	"fmt"
	"strconv"
	"strings"

	"aoc2024.com/util"
)

type Solver struct {
	grid    [100][100]bool
	updates [][]int
}

func New(lines []string) Solver {
	grid := [100][100]bool{}
	var updates [][]int

	for _, line := range lines {
		split_by_vert_bar := strings.Split(line, "|")
		if len(split_by_vert_bar) == 2 {
			before, err := strconv.Atoi(split_by_vert_bar[0])
			util.CheckWithFatalError(err, "Failed to read int %v")
			after, err := strconv.Atoi(split_by_vert_bar[1])
			util.CheckWithFatalError(err, "Failed to read int %v")
			grid[after][before] = true
		}
		split_by_comma := strings.Split(line, ",")
		if len(split_by_comma) > 1 {
			var ints []int
			for _, item := range split_by_comma {
				v, err := strconv.Atoi(item)
				util.CheckWithFatalError(err, "Failed to read int %v")
				ints = append(ints, v)
			}
			updates = append(updates, ints)
		}
	}
	return Solver{grid: grid, updates: updates}
}

func (s Solver) Part1() int {
	var middle int
	var total int
	for _, line := range s.updates {
		middle = line[len(line)/2]
		fail := false
		for i, item := range line {
			for _, comp := range line[:i] {
				if s.grid[comp][item] {
					fail = true
					break
				}
			}
			if fail {
				break
			}
		}
		if !fail {
			total += middle
		}
	}
	return total
}

func sort(arr []int, s Solver) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < i; j++ {
			if s.grid[arr[j]][arr[i]] {
				copy := arr[i]
				arr[i] = arr[j]
				arr[j] = copy
			}
		}
	}
}

func (s Solver) Part2() int {
	fmt.Println()
	var total int
	for _, line := range s.updates {
		fail := false
		for i, item := range line {
			for _, comp := range line[:i] {
				if s.grid[comp][item] {
					fail = true
					lineCopy := make([]int, len(line))
					copy(lineCopy, line)
					sort(lineCopy, s)
					middle := lineCopy[len(lineCopy)/2]
					fmt.Println(middle)
					total += middle
					break
				}
			}
			if fail {
				break
			}
		}
	}
	return total
}
