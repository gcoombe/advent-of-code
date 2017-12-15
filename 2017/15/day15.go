package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)


func main() {

	fmt.Println("----------------------------------------")
	fmt.Println("DAY 15")
	fmt.Println("----------------------------------------")

	fmt.Printf("\nPart 1 res: %d", runPartOne(289, 629, 16807, 48271, 40000000))
	fmt.Printf("\nPart 2 res: %d", runPartTwo(289, 629, 16807, 48271, 4, 8, 5000000))

	fmt.Println()
	fmt.Println("----------------------------------------")
	fmt.Println("END DAY 15")
	fmt.Println("----------------------------------------")

}

const DIV_FACTOR = int(2147483647)


func times(str string, n int) (out string) {
	for i := 0; i < n; i++ {
		out += str
	}
	return
}

// Left left-pads the string with pad up to len runes
// len may be exceeded if
func padLeft(str string, len int, pad string) string {
	return times(pad, len-utf8.RuneCountInString(str)) + str
}

func countMatches(iterationCount int, genA, genB []int) int {
	matchCount := 0
	for i := 0; i < iterationCount; i++ {
		binA := padLeft(strconv.FormatInt(int64(genA[i]), 2),16, "0")
		binB := padLeft(strconv.FormatInt(int64(genB[i]), 2),16, "0")

		if (binA[len(binA) - 16:] == binB[len(binB) - 16:]) {
			matchCount++
		}

	}
	return matchCount
}

func runPartOne(startA, startB, factorA, factorB, iterationCount int) int {

	genA := make([]int, iterationCount + 1)
	genB := make([]int, iterationCount + 1)

	currValA := startA
	currValB := startB
	for i := 0; i < iterationCount; i++ {
		genA[i] = (currValA * factorA)  % DIV_FACTOR
		genB[i] = (currValB * factorB)  % DIV_FACTOR

		currValA = genA[i]
		currValB = genB[i]
	}

	return countMatches(iterationCount, genA, genB)
}


func runPartTwo(startA, startB, factorA, factorB, criteriaA, criteriaB, iterationCount int) int {

	genA := make([]int, iterationCount)
	genB := make([]int, iterationCount)

	genA[iterationCount - 1] = -1
	genB[iterationCount - 1] = -1

	currValA := startA
	currValB := startB

	i := 0
	j := 0
	for i < iterationCount || j < iterationCount {
		if (i < iterationCount) {
			currValA = (currValA * factorA)  % DIV_FACTOR
			if (currValA % criteriaA == 0) {
				genA[i] = currValA
				i++
			}
		}

		if (j < iterationCount) {
			currValB = (currValB * factorB)  % DIV_FACTOR
			if (currValB % criteriaB == 0) {
				genB[j] = currValB
				j++
			}
		}
	}

	return countMatches(iterationCount, genA, genB)
}
