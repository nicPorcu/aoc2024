package day6

import "fmt"

type Direction int
type Result int

const (
	UP Direction = iota
	DOWN
	LEFT
	RIGHT
)

const (
	GOT_OUT Result = iota
	TRAPPED
)

type Solver struct {
	grid      [][]rune
	row       int
	col       int
	facing    Direction
	start_row int
	start_col int
}

type Offset struct {
	x int
	y int
}

func New(lines []string) Solver {
	var grid [][]rune
	var row int
	var col int
	var facing Direction

	for r, line := range lines {
		grid = append(grid, []rune(line))
		vals := []rune(line)

		for c, v := range vals {
			switch v {
			case '.':
				continue
			case '#':
				continue
			case 'v':
				facing = DOWN
			case '^':
				facing = UP
			case '>':
				facing = RIGHT
			case '<':
				facing = LEFT
			}
			col = c
			row = r
		}
	}
	return Solver{grid: grid, facing: facing, row: row, col: col, start_row: row, start_col: col}
}

func findOffset(dir Direction) Offset {
	var offset Offset
	switch dir {
	case UP:
		offset = Offset{x: 0, y: -1}
	case DOWN:
		offset = Offset{x: 0, y: 1}
	case LEFT:
		offset = Offset{x: -1, y: 0}
	case RIGHT:
		offset = Offset{x: 1, y: 0}
	}
	return offset
}

func findNextDir(dir Direction) Direction {
	var next_dir Direction
	switch dir {
	case UP:
		next_dir = RIGHT
	case DOWN:
		next_dir = LEFT
	case LEFT:
		next_dir = UP
	case RIGHT:
		next_dir = DOWN
	}
	return next_dir
}

func (s Solver) Part1() int {
	height := len(s.grid)
	width := len(s.grid[0])
	s.grid[s.start_row][s.start_col] = 'X'
	fmt.Printf("(%d, %d) \n", s.start_row, s.start_col)
	cnt := 1

	for {
		if s.grid[s.row][s.col] == '.' {
			cnt += 1
			fmt.Printf("(%d, %d) \n", s.row, s.col)
			s.grid[s.row][s.col] = 'X'
		}
		offset := findOffset(s.facing)
		other_row := s.row + offset.y
		other_col := s.col + offset.x
		if other_row < 0 || other_col < 0 || other_col >= width || other_row >= height {
			break
		} else if s.grid[other_row][other_col] == '#' {
			s.facing = findNextDir(s.facing)
		} else {
			s.row = other_row
			s.col = other_col
		}
	}
	s.grid[s.start_row][s.start_col] = '^'
	return cnt
}

func runTheThing(s Solver) Result {

	type Point struct {
		row    int
		col    int
		facing Direction
	}

	height := len(s.grid)
	width := len(s.grid[0])

	seen := make(map[Point]struct{})
	for {
		offset := findOffset(s.facing)
		other_row := s.row + offset.y
		other_col := s.col + offset.x
		if other_row < 0 || other_col < 0 || other_col >= width || other_row >= height {
			return GOT_OUT
		} else if s.grid[other_row][other_col] == '#' {
			s.facing = findNextDir(s.facing)
		} else {
			s.row = other_row
			s.col = other_col
			point := Point{row: s.row, col: s.col, facing: s.facing}
			_, ok := seen[point]
			if ok {
				return TRAPPED
			}
			seen[point] = struct{}{}
		}
	}
}

func (s Solver) Part2() int {
	cnt := 0
	for i, row := range s.grid {
		for j, _ := range row {
			if s.grid[i][j] != 'X' {
				continue
			}
			s.grid[i][j] = '#'
			res := runTheThing(s)
			if res == TRAPPED {
				cnt += 1
				//fmt.Printf("(%d, %d) \n", i, j)
			}
			s.grid[i][j] = 'X'
		}
	}
	return cnt
}
