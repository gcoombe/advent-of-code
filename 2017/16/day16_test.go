package main

import (
"testing"
"runtime"
"path"
)

func TestPart1(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), "inputTest.txt")

	programs := runPartOne(filepath, "abcde")
	if (programs != "baedc") {
		t.Errorf("Incorrest answer.  Expected %s got %s", "baedc", programs)
	}
}
