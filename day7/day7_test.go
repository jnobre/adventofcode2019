package day7_test

import (
	"adventofcode2019/day7"
	"adventofcode2019/input"
	"testing"
)

func TestPart1(t *testing.T) {
	answer := day7.Part1(input.MainFileName)
	expected := "95757"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := day7.Part2(input.MainFileName)
	expected := "4275738"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
