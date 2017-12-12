package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
	"path"
	"runtime"
)

func main() {

	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "input.txt")

	fmt.Println("----------------------------------------")
	fmt.Println("DAY 2")
	fmt.Println("----------------------------------------")

	fmt.Printf("\nPart 1 res: %d", runPartOne(filepath))

	fmt.Println()
	fmt.Println("----------------------------------------")
	fmt.Println("END DAY 2")
	fmt.Println("----------------------------------------")

}


func getFile(path string) *os.File {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil
	}
	return file
}

func runPartOne(filePath string) int {
	file := getFile(filePath)
	if file == nil {
		return 0
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	checksum := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		values := strings.Split(line, "\t")

		lowest, _ := strconv.Atoi(values[0])
		highest := lowest

		for i := 1; i < len(values); i++ {
			value, _ := strconv.Atoi(values[i])
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