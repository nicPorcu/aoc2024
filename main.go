package main

import (
	"flag"
	"fmt"

	"aoc2024.com/day1"
)

type SolverInterface interface {
	Part1() int
	Part2() int
}

var day int

func main() {
	flag.IntVar(&day, "day", 1, "The day of AOC to run")
	flag.Parse()
	tail := flag.Args()
	if len(tail) > 1 {
		panic("Invalid arg length")
	}

	dayToSolver := map[int]func() SolverInterface{
		1: func() SolverInterface {
			return day1.New()
		},
	}
	solver := dayToSolver[day]()
	fmt.Println(solver.Part1())
	fmt.Println(solver.Part2())
}
