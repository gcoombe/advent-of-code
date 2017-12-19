package main

import (
	"runtime"
	"path"
	"fmt"
	"advent-of-code/2017/goUtils"
	"bufio"
)

func main() {

	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "input.txt")

	fmt.Println("----------------------------------------")
	fmt.Println("DAY 19")
	fmt.Println("----------------------------------------")

	fmt.Printf("\nPart 1 res: %s", runPartOne(filepath))
	fmt.Printf("\nPart 2 res: %d", runPartTwo(filepath))

	fmt.Println()
	fmt.Println("----------------------------------------")
	fmt.Println("END DAY 19")
	fmt.Println("----------------------------------------")
}

type GridSolver struct {
	grid [][]rune
	currX int
	currY int
	reachedEnd bool
	letterOrder string
	currDirection rune
	steps int
}

func (solver *GridSolver) canNavigate(direction rune) bool {
	if direction == 'N' {
		return solver.currY > 0 && solver.grid[solver.currY - 1][solver.currX] != ' '
	} else if direction == 'S' {
		return solver.currY < len(solver.grid) - 1 && solver.grid[solver.currY + 1][solver.currX] != ' '
	} else if direction == 'W' {
		return solver.currX > 0 && solver.grid[solver.currY][solver.currX - 1] != ' '
	} else {
		return solver.currX < len(solver.grid[solver.currY]) - 1 && solver.grid[solver.currY][solver.currX + 1] != ' '
	}
}

func (solver *GridSolver) currRune() rune {
	return solver.grid[solver.currY][solver.currX]
}

func (solver *GridSolver) navigate(direction rune) {
	if direction == 'N' {
		solver.currY = solver.currY - 1
	} else if direction == 'S' {
		solver.currY = solver.currY + 1
	} else if direction == 'W' {
		solver.currX = solver.currX - 1
	} else {
		solver.currX = solver.currX + 1
	}
	solver.currDirection = direction
	solver.steps++
}

func (solver *GridSolver) next() {
	currRune := solver.currRune()
	if currRune != '|' && currRune != '-' && currRune != '+' {
		solver.letterOrder = solver.letterOrder + string(solver.currRune())
	}

	if solver.currRune() == '+' {
		if solver.currDirection == 'N' || solver.currDirection == 'S' {
			if solver.canNavigate('W') {
				solver.navigate('W')
			} else {
				solver.navigate('E')
			}
		} else  {
			if solver.canNavigate('N') {
				solver.navigate('N')
			} else {
				solver.navigate('S')
			}
		}
	} else if solver.canNavigate(solver.currDirection) {
		solver.navigate(solver.currDirection)
	} else {
		solver.reachedEnd = true
	}
}

func buildGrid(filePath string) *GridSolver {
	grid := make([][]rune, 0)

	file := goUtils.GetFile(filePath)

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	maxWidth := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		runes := []rune(line)

		row := make([]rune, len(line))
		for i, val := range runes {
			row[i] = val
		}
		grid = append(grid, row)

		if len(line) > maxWidth {
			maxWidth = len(line)
		}
	}

	//Pad all rows to be the same width
	for i, row := range grid {
		if len(row) < maxWidth {
			originalLength := len(row)
			for j := originalLength; j < maxWidth; j++ {
				row = append(row, ' ')
			}

			grid[i] = row
		}
	}

	var startX int
	for x, val := range grid[0] {
		if val == '|' {
			startX = x
			break
		}
	}

	return &GridSolver{grid:grid, currX:startX, currY:0, reachedEnd:false, letterOrder:"", currDirection:'S', steps:1}

}

func runPartOne(filePath string) string {
	grid := buildGrid(filePath)

	for grid.reachedEnd == false {
		grid.next()
	}

	return grid.letterOrder
}

func runPartTwo(filePath string) int {
	grid := buildGrid(filePath)

	for grid.reachedEnd == false {
		grid.next()
	}

	return grid.steps
}