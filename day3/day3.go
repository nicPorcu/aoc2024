package day3

import (
	"strconv"
	"strings"
)

type Solver struct {
	lines []string
}

func New(lines []string) Solver {
	return Solver{lines}
}

func solveLine(line string) int {
	s := strings.Split(line, ")")
	if len(s) < 1 {
		return 0
	}
	vals := strings.Split(s[0], ",")
	if len(vals) != 2 {
		return 0
	}
	v0, err0 := strconv.Atoi(vals[0])
	v1, err1 := strconv.Atoi(vals[1])
	if err0 != nil || err1 != nil {
		return 0
	}
	return v0 * v1
}

func (s Solver) Part1() int {
	total := 0
	for _, line := range s.lines {
		parts := strings.Split(line, "mul(")
		for i, part := range parts {
			//fmt.Println(part)
			if i == 0 {
				continue
			}
			total += solveLine(part)
		}
	}
	return total
}

func (s Solver) Part2() int {
	enabled := true
	total := 0
	for _, line := range s.lines {

		parts := strings.Split(line, "do")

		for _, part := range parts {
			if strings.HasPrefix(part, "()") {
				enabled = true
			} else if strings.HasPrefix(part, "n't()") {
				enabled = false
			}
			if enabled {
				items := strings.Split(part, "mul(")
				for j, item := range items {
					if j == 0 {
						continue
					}
					solve := solveLine(item)
					total += solve
				}

			}

		}
	}

	return total
}
