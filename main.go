package main

import (
	"flag"
	"fmt"
	"strconv"

	"aoc2024.com/day1"
	"aoc2024.com/day2"
	"aoc2024.com/day3"
	"aoc2024.com/day4"
	"aoc2024.com/day5"
	"aoc2024.com/day6"
	"aoc2024.com/day7"
	"aoc2024.com/util"
)

type SolverInterface interface {
	Part1() int
	Part2() int
}

var day int
var short bool

func main() {
	flag.IntVar(&day, "day", 1, "The day of AOC to run")
	flag.BoolVar(&short, "short", false, "whether to run short")
	flag.Parse()
	tail := flag.Args()
	if len(tail) > 1 {
		panic("Invalid arg length")
	}
	folder := "day" + strconv.Itoa(day)
	filename := "input.txt"
	if short {
		filename = "short.txt"
	}
	path := folder + "/" + filename
	f := util.ReadFile(path)

	dayToSolver := map[int]func() SolverInterface{
		1: func() SolverInterface {
			return day1.New(f)
		},
		2: func() SolverInterface {
			return day2.New(f)
		},
		3: func() SolverInterface {
			return day3.New(f)
		},
		4: func() SolverInterface {
			return day4.New(f)
		},
		5: func() SolverInterface {
			return day5.New(f)
		},
		6: func() SolverInterface {
			return day6.New(f)
		},
		7: func() SolverInterface {
			return day7.New(f)
		},
	}
	solver := dayToSolver[day]()
	fmt.Println(solver.Part1())
	fmt.Println(solver.Part2())
}
