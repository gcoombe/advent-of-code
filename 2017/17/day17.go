package main

import (
	"fmt"
)


func main() {

	fmt.Println("----------------------------------------")
	fmt.Println("DAY 17")
	fmt.Println("----------------------------------------")

	fmt.Printf("\nPart 1 res: %d", runPartOne(367, 2017))
	fmt.Printf("\nPart 2 res: %d", runPartTwo(367, 50000000))

	fmt.Println()
	fmt.Println("----------------------------------------")
	fmt.Println("END DAY 17")
	fmt.Println("----------------------------------------")

}

func buildBuffer(input, iterations int) ([]int, int) {
	buffer := []int{0}
	currPos := 0

	for i := 1; i < iterations + 1; i++ {
		insertionPoint := ((currPos + input) % len(buffer)) + 1

		nextBuffer := make([]int, insertionPoint)

		copy(nextBuffer, buffer[0:insertionPoint])
		nextBuffer = append(nextBuffer, i)
		buffer = append(nextBuffer, buffer[insertionPoint:]...)
		currPos = insertionPoint
	}

	return buffer, currPos
}

func runPartOne(input, iterations int) int {
	buffer, currPos := buildBuffer(input, iterations)
	return buffer[currPos + 1]
}


func runPartTwo(input, iterations int) int {
	currPos := 0
	returnVal := 0

	for i := 1; i < iterations + 1; i++ {
		currPos = ((currPos + input) % i) + 1
		if currPos == 1 {
			returnVal = i
		}
	}

	return returnVal
}