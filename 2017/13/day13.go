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
	fmt.Println("DAY 13")
	fmt.Println("----------------------------------------")

	fmt.Printf("\nPart 1 res: %d", runPartOne(filepath))
	fmt.Printf("\nPart 2 res: %d", runPartTwo(filepath))

	fmt.Println()
	fmt.Println("----------------------------------------")
	fmt.Println("END DAY 13")
	fmt.Println("----------------------------------------")

}



type Depth struct  {
	depthVal, scannerRange int
}

func (depth *Depth) posAtTime(delay int) int {
	return  (depth.depthVal + delay) % ((depth.scannerRange - 1) * 2)
}

func newDepth(depthVal, scannerRange int) *Depth {
	return &Depth{scannerRange: scannerRange, depthVal: depthVal}
}

func buildDepths(filePath string) map[int]*Depth {
	depths := make(map[int]*Depth)

	file := goUtils.GetFile(filePath)

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		values := strings.Split(line, ": ")
		scannerRange, _ := strconv.Atoi(values[1])

		depthVal, _ := strconv.Atoi(values[0])
		depths[depthVal] = newDepth(depthVal, scannerRange)

	}
	return depths

}

func runPartOne(filePath string) int {
	depths := buildDepths(filePath)

	severity := 0
	for _, currDepth := range depths{
		if currDepth.posAtTime(0) == 0 {
			severity += currDepth.depthVal * currDepth.scannerRange
		}
	}
	return severity
}


func runPartTwo(filePath string) int {

	depths := buildDepths(filePath)

	var hitScanner bool
	//Random limit.
	for delay := 0; ; delay++ {
		hitScanner = false

		for _, currDepth := range depths {

			if currDepth.posAtTime(delay) == 0 {
				hitScanner = true
				break
			}

		}

		if hitScanner == false {
			return delay
		}
	}

	return -1
}