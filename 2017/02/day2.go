package main

import (
	"bufio"
	"fmt"
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
	fmt.Println("DAY 2")
	fmt.Println("----------------------------------------")

	fmt.Printf("\nPart 1 res: %d", runPartOne(filepath))
	fmt.Printf("\nPart 2 res: %d", runPartTwo(filepath))

	fmt.Println()
	fmt.Println("----------------------------------------")
	fmt.Println("END DAY 2")
	fmt.Println("----------------------------------------")

}

func listStringsToInts(strings []string) []int {
	nums := make([]int, len(strings))
	for i, string := range strings {
		nums[i], _ = strconv.Atoi(string)
	}
	return nums
}

func runPartOne(filePath string) int {
	file := goUtils.GetFile(filePath)
	if file == nil {
		return 0
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	checksum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		values := listStringsToInts(strings.Split(line, "\t"))

		lowest := values[0]
		highest := lowest

		for i := 1; i < len(values); i++ {
			value := values[i]
			if value < lowest {
				lowest = value
			} else if value > highest {
				highest = value
			}
		}
		checksum += highest - lowest
	}

	return checksum
}

func runPartTwo(filePath string) int {
	file := goUtils.GetFile(filePath)
	if file == nil {
		return 0
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	checksum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		values := listStringsToInts(strings.Split(line, "\t"))

		for i, num1 := range values {
			for _, num2 := range values[i + 1:] {
				if num1 != 0 && num2 != 0 {
					if num1 % num2 == 0 {
						checksum += num1 / num2
					} else if num2 % num1 == 0 {
						checksum += num2 / num1
					}
				}
			}
		}
	}

	return checksum
}