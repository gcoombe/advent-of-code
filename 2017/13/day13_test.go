package main

import (
	"testing"
	"runtime"
	"path"
)

func TestPart1(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "testInput.txt")

	severity := runPartOne(filepath)
	if (severity != 24) {
		t.Errorf("Incorrest answer.  Expected %d got %d", 24, severity)
	}
}


func TestPart2(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "testInput.txt")

	delay := runPartTwo(filepath)
	if (delay != 10) {
		t.Errorf("Incorrest answer.  Expected %d got %d", 10, delay)
	}
}
