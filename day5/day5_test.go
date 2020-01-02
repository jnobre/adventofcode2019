package day5_test

import (
	"adventofcode2019/day5"
	"adventofcode2019/input"
	"testing"
)

func TestPart1(t *testing.T) {
	answer := day5.Part1(input.MainFileName)
	expected := "5346030"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := day5.Part2(input.MainFileName)
	expected := "513116"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
