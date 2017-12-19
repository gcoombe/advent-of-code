package main

import (
"testing"
)

func TestPart1(t *testing.T) {
	result := runPartOne(3, 15)
	if (result != 638) {
		t.Errorf("Incorrest answer.  Expected %d got %d", 638, result)
	}
}

