package day4

type Solver struct {
	grid [][]rune
}

const searchString = "XMAS"
const searchStringInv = "SAMX"

const searchString2 = "MAS"
const searchString2Inv = "SAM"

func New(lines []string) Solver {
	var grid [][]rune
	for _, s := range lines {
		grid = append(grid, []rune(s))
	}
	return Solver{grid: grid}
}

func (s Solver) Part1() int {
	height := len(s.grid)
	width := len(s.grid[0])
	cnt := 0
	for i, row := range s.grid {
		for j, _ := range row {
			var localVals [][4]rune
			var search string
			if s.grid[i][j] == rune(searchString[0]) {
				search = searchString
			} else if s.grid[i][j] == rune(searchStringInv[0]) {
				search = searchStringInv
			} else {
				continue
			}
			if j >= 3 && i+4 <= height {
				localVals = append(localVals, [4]rune{s.grid[i][j], s.grid[i+1][j-1], s.grid[i+2][j-2], s.grid[i+3][j-3]})
			}
			if i+4 <= height && j+4 <= width {
				localVals = append(localVals, [4]rune{s.grid[i][j], s.grid[i+1][j+1], s.grid[i+2][j+2], s.grid[i+3][j+3]})
			}
			if i+4 <= height {
				localVals = append(localVals, [4]rune{s.grid[i][j], s.grid[i+1][j], s.grid[i+2][j], s.grid[i+3][j]})
			}
			if j+4 <= width {
				localVals = append(localVals, [4]rune{s.grid[i][j], s.grid[i][j+1], s.grid[i][j+2], s.grid[i][j+3]})
			}

			for _, val := range localVals {
				if string(val[:]) == search {
					cnt += 1
				}
			}
		}
	}
	return cnt
}

func (s Solver) Part2() int {
	cnt := 0
	height := len(s.grid)
	width := len(s.grid[0])
	for i, row := range s.grid {
		for j, _ := range row {
			if i <= 0 || i >= height-1 || j <= 0 || j >= width-1 {
				continue
			}
			if s.grid[i][j] == 'A' {
				val1 := [3]rune{s.grid[i-1][j-1], s.grid[i][j], s.grid[i+1][j+1]}
				val2 := [3]rune{s.grid[i+1][j-1], s.grid[i][j], s.grid[i-1][j+1]}
				//fmt.Printf("%v, %v \n", i, j)
				//fmt.Printf("%v, %v \n", string(val1[:]), string(val2[:]))
				if string(val1[:]) == searchString2 || string(val1[:]) == searchString2Inv {
					if string(val2[:]) == searchString2 || string(val2[:]) == searchString2Inv {
						cnt += 1
						//fmt.Printf("MATCH %v, %v \n", i, j)
					}
				}
			}
		}
	}
	return cnt
}
