package main

import (
"fmt"
"bufio"
"strings"
"strconv"
"path"
"runtime"
"advent-of-code/2017/goUtils"
)


func main() {

	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "input.txt")

	fmt.Println("----------------------------------------")
	fmt.Println("DAY 18")
	fmt.Println("----------------------------------------")

	fmt.Printf("\nPart 1 res: %d", runPartOne(filepath))
	fmt.Printf("\nPart 2 res: %d", runPartTwo(filepath))

	fmt.Println()
	fmt.Println("----------------------------------------")
	fmt.Println("END DAY 18")
	fmt.Println("----------------------------------------")

}

func getInstrumentComponents(instruction string, registers map[string] int) (string, string, int) {
	instructionComponents := strings.Split(instruction, " ")

	component1 := instructionComponents[0]
	component2 := instructionComponents[1]

	if len(instructionComponents) > 2 {
		intVal, err := strconv.Atoi(instructionComponents[2])
		if err == nil {
			return component1, component2, intVal
		}
		return component1, component2, registers[instructionComponents[2]]
	}

	return component1, component2, 0
}

func evaluateInstructionsPart1(instructions []string) int {
	registers :=  make(map[string] int)
	var lastSound int

	currInstructionIdx := 0

	for currInstructionIdx < len(instructions) {
		currInstruction := instructions[currInstructionIdx]
		nextInstructionIdx := currInstructionIdx + 1

		iComp1, iComp2, icomp3 := getInstrumentComponents(currInstruction, registers)

		switch iComp1 {

		case "snd":
			lastSound = registers[iComp2]
		case "set":
			registers[iComp2] = icomp3
		case "add":
			registers[iComp2] = icomp3 + registers[iComp2]
		case "mul":
			registers[iComp2] = icomp3 * registers[iComp2]
		case "mod":
			registers[iComp2] = registers[iComp2] % icomp3
		case "rcv":
			valToRec := registers[iComp2]
			if valToRec != 0 {
				registers[iComp2] = lastSound
			}
			return lastSound
		case "jgz":
			valInReg := registers[iComp2]
			if valInReg > 0 {
				nextInstructionIdx = currInstructionIdx + icomp3
			}
		}

		currInstructionIdx = nextInstructionIdx
	}

	return lastSound
}

func readInstructions(filePath string) []string {
	var instructions []string

	file := goUtils.GetFile(filePath)

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		instructions = append(instructions, line)
	}
	return instructions
}

func runPartOne(filePath string) int {
	instructions := readInstructions(filePath)
	return evaluateInstructionsPart1(instructions)
}

type Program struct  {
	programId int
	instructions []string
	registers map[string] int
	currInstructionIdx int
	sendCount int
	messageQueue []int
	deadLocked bool
}

func (program *Program) queueMessage(val int) {
	program.messageQueue = append(program.messageQueue, val)
}

//Returns 0 normally or 1 if we have deadlocked
func (program *Program) processInstruction(otherProgam *Program) {
	if program.currInstructionIdx >= len(program.instructions) {
		program.deadLocked = true
		return
	}

	program.deadLocked = false
	currInstruction := program.instructions[program.currInstructionIdx]
	nextInstructionIdx := program.currInstructionIdx + 1

	iComp1, iComp2, icomp3 := getInstrumentComponents(currInstruction, program.registers)

	switch iComp1 {

	case "snd":
		valToQueue, err := strconv.Atoi(iComp2)
		if err != nil {
			valToQueue = program.registers[iComp2]
		}
		otherProgam.queueMessage(valToQueue)
		program.sendCount++
	case "set":
		program.registers[iComp2] = icomp3
	case "add":
		program.registers[iComp2] = icomp3 + program.registers[iComp2]
		case "mul":
		program.registers[iComp2] = icomp3 * program.registers[iComp2]
	case "mod":
		program.registers[iComp2] = program.registers[iComp2] % icomp3
	case "rcv":
		if len(program.messageQueue) > 0 {
			var val int
			val, program.messageQueue = program.messageQueue[0], program.messageQueue[1:]
			program.registers[iComp2] = val
		} else {
			nextInstructionIdx = program.currInstructionIdx
			program.deadLocked = true
		}

	case "jgz":
		shouldJump, err := strconv.Atoi(iComp2)
		if err != nil {
			shouldJump = program.registers[iComp2]
		}
		if shouldJump > 0 {
			nextInstructionIdx = program.currInstructionIdx + icomp3
		}
	}

	program.currInstructionIdx = nextInstructionIdx
}

func createProgram(instructions []string, programId int) *Program {
	initialRegisters := make(map[string]int)
	initialRegisters["p"] = programId
	return & Program{instructions:instructions, currInstructionIdx: 0, registers:initialRegisters, sendCount:0, deadLocked:false, programId:programId}
}

func runPartTwo(filePath string) int {
	instructions := readInstructions(filePath)
	program0 := createProgram(instructions, 0)
	program1 := createProgram(instructions, 1)

	for program0.deadLocked == false || program1.deadLocked == false {
		program0.processInstruction(program1)
		program1.processInstruction(program0)
	}

	return  program1.sendCount
}

