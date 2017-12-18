package main

import (
"testing"
"runtime"
"path"
)

func TestPart1(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "inputTest.txt")

	lastPlayed := runPartOne(filepath)
	if (lastPlayed != 4) {
		t.Errorf("Incorrest answer.  Expected %d got %d", 4, lastPlayed)
	}
}

func TestPart2(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "testInput2.txt")

	sends := runPartTwo(filepath)
	if (sends != 3) {
		t.Errorf("Incorrest answer.  Expected %d got %d", 3, sends)
	}
}
