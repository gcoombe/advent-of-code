package main

import (
	"runtime"
	"path"
	"fmt"
	"io/ioutil"
	"strings"
	"strconv"
)

func main() {

	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "input.txt")

	fmt.Println("----------------------------------------")
	fmt.Println("DAY 16")
	fmt.Println("----------------------------------------")

	fmt.Printf("\nPart 1 res: %s", runPartOne(filepath, "abcdefghijklmnop"))
	fmt.Printf("\nPart 2 res: %s", runPartTwo(filepath, "abcdefghijklmnop"))

	fmt.Println()
	fmt.Println("----------------------------------------")
	fmt.Println("END DAY 16")
	fmt.Println("----------------------------------------")
}

func getInstructions(filePath string) [][]string {
	b, err := ioutil.ReadFile(filePath) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	instructions := make([][]string, 0)
	for _, instruction := range strings.Split(string(b), ",") {
		if instruction[0] == 's' {
			instructions = append(instructions, []string{string(instruction[0]), instruction[1:]})
		} else {
			operandComponents := strings.Split(instruction[1:], "/")
			instructions = append(instructions, []string{string(instruction[0]), operandComponents[0], operandComponents[1]})
		}
	}
	return  instructions

}

func swap(programs string, indexA, indexB int) string {
	programSlice := []rune(programs)
	firstElm := programSlice[indexA]
	programSlice[indexA] = programSlice[indexB]
	programSlice[indexB] = firstElm
	return string(programSlice)
}

func performDance(programs string, instructions [][]string) string {
	for _, instruction := range instructions {
		if instruction[0] == "s" {
			lenToMove, _ := strconv.Atoi(instruction[1])
			programs = programs[len(programs) - lenToMove:] + programs[0:len(programs) - lenToMove]
		} else if instruction[0] == "x" {
			targetA, _ := strconv.Atoi(string(instruction[1]))
			targetB, _ := strconv.Atoi(string(instruction[2]))
			programs = swap(programs, targetA, targetB)
		} else {
			targetA := strings.IndexRune(programs, rune(instruction[1][0]))
			targetB := strings.IndexRune(programs, rune(instruction[2][0]))
			programs = swap(programs, targetA, targetB)
		}
	}

	return programs
}

func runPartOne(filePath string, startState string) string {
	programs := startState
	instructions := getInstructions(filePath)

	return  performDance(programs, instructions)
}

func runPartTwo(filePath string, startState string) string {
	programs := startState
	instructions := getInstructions(filePath)

	for i := 0; i < 1000000000; i++ {
		programs = performDance(programs, instructions)
	}

	return programs
}


