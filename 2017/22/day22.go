package main

import (
	"runtime"
	"path"
	"fmt"
	"advent-of-code/2017/goUtils"
	"bufio"
	"strings"
	"strconv"
)


func main() {

	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "input.txt")

	fmt.Println("----------------------------------------")
	fmt.Println("DAY 21")
	fmt.Println("----------------------------------------")

	fmt.Printf("\nPart 1 res: %d", runPartOne(filepath, 10000))
	fmt.Printf("\nPart 2 res: %d", runPartTwo(filepath, 10000000))

	fmt.Println()
	fmt.Println("----------------------------------------")
	fmt.Println("END DAY 21")
	fmt.Println("----------------------------------------")
}

type GridSolver struct {
	grid map[string]rune
	currX, currY int
	direction rune
	infectionCount int
}

func (solver GridSolver) getKey()string {
	return strconv.Itoa(solver.currX) + "," + strconv.Itoa(solver.currY)
}

func (solver GridSolver) currNode()rune {
	if val, ok := solver.grid[solver.getKey()]; ok {
		return val
	}
	return '.'
}

func (solver *GridSolver) setNode(x, y int, value rune) {
	solver.grid[strconv.Itoa(x) + "," + strconv.Itoa(y)] = value
}

func (solver *GridSolver) setCurrNode(value rune) {
	solver.grid[solver.getKey()] = value
}

func (solver *GridSolver) rotate() {
	if solver.direction == 'N' && solver.currNode() == '#' || solver.direction == 'S' && solver.currNode() == '.' ||
			solver.direction == 'W' && solver.currNode() == 'F' {
		solver.direction = 'E'
	} else if solver.direction == 'E' && solver.currNode() == '#' || solver.direction == 'W' && solver.currNode() == '.' ||
			solver.direction == 'N' && solver.currNode() == 'F' {
		solver.direction = 'S'
	} else if solver.direction == 'S' && solver.currNode() == '#' || solver.direction == 'N' && solver.currNode() == '.' ||
			solver.direction == 'E' && solver.currNode() == 'F' {
		solver.direction = 'W'
	} else if solver.direction == 'W' && solver.currNode() == '#' || solver.direction == 'E' && solver.currNode() == '.' ||
			solver.direction == 'S' && solver.currNode() == 'F'{
		solver.direction = 'N'
	}
}

func (solver *GridSolver) update2() {
	if solver.currNode() == '.' {
		solver.setCurrNode( 'W')
	} else if solver.currNode() == 'W' {
		solver.infectionCount++
		solver.setCurrNode('#')
	} else if solver.currNode() == '#' {
		solver.setCurrNode('F')
	} else {
		solver.setCurrNode( '.')
	}
}

func (solver *GridSolver) update1() {
	if solver.currNode() == '.' {
		solver.infectionCount++
		solver.setCurrNode( '#')
	} else {
		solver.setCurrNode( '.')
	}
}

func (solver *GridSolver) move() {
	if solver.direction == 'N' {
		solver.currY = solver.currY - 1
	} else if solver.direction == 'E' {
		solver.currX = solver.currX + 1
	} else if solver.direction == 'S' {
		solver.currY = solver.currY + 1
	} else {
		solver.currX = solver.currX - 1
	}
}

func buildSolver(filePath string) *GridSolver {
	file := goUtils.GetFile(filePath)

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	grid := make(map[string]rune, 0)
	solver := &GridSolver{grid:grid, currX:0, currY:0, direction:'N', infectionCount:0}

	lineCount := 0
	lineLength := 0
	for fileScanner.Scan() {
		line := strings.Replace(fileScanner.Text(), " ", "", -1)
		lineLength = len(line)
		for i, char := range line {
			if char == '#' {
				solver.setNode(i, lineCount, char)
			}
		}
		lineCount++
	}

	solver.currX = (lineLength-1) /2
	solver.currY = (lineCount-1) /2
	return solver
}

func runPartOne(filePath string, rounds int) int {
	gridSolver := buildSolver(filePath)
	for i := 0; i < rounds; i++ {
		gridSolver.rotate()
		gridSolver.update1()
		gridSolver.move()
	}
	return gridSolver.infectionCount
}

func runPartTwo(filePath string, rounds int) int {
	gridSolver := buildSolver(filePath)
	for i := 0; i < rounds; i++ {
		gridSolver.rotate()
		gridSolver.update2()
		gridSolver.move()
	}
	return gridSolver.infectionCount
}

