package main

import (
	"testing"
	"runtime"
	"path"
)

func TestPart1(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "testInput.txt")

	order := runPartOne(filepath)
	if order != "ABCDEF" {
		t.Errorf("Incorrest answer.  Expected %s got %s", "ABCDEF", order)
	}
}

func TestPart2(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "testInput.txt")

	steps := runPartTwo(filepath)
	if steps != 38 {
		t.Errorf("Incorrest answer.  Expected %d got %d", 38, steps)
	}
}
