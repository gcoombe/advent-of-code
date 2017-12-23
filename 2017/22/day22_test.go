package main

import (
	"testing"
	"runtime"
	"path"
)

func TestPart1Short(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "testInput.txt")

	count := runPartOne(filepath, 7)
	if count != 5 {
		t.Errorf("Incorrest answer.  Expected %d got %d", 5, count)
	}
}

func TestPart1(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "testInput.txt")

	count := runPartOne(filepath, 70)
	if count != 41 {
		t.Errorf("Incorrest answer.  Expected %d got %d", 41, count)
	}
}

func TestPart2(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "testInput.txt")

	count := runPartTwo(filepath, 100)
	if count != 26 {
		t.Errorf("Incorrest answer.  Expected %d got %d", 100, count)
	}
}