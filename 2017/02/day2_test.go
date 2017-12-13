package main

import (
"testing"
"runtime"
"path"
)

func TestPart1(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "testInput.txt")

	checksum := runPartOne(filepath)
	if (checksum != 18) {
		t.Errorf("Incorrest answer.  Expected %d got %d", 18, checksum)
	}
}

func TestPart2(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "testInput2.txt")

	checksum := runPartTwo(filepath)
	if (checksum != 9) {
		t.Errorf("Incorrest answer.  Expected %d got %d", 9, checksum)
	}
}
