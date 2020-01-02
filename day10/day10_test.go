package day10_test

import (
	"adventofcode2019/day10"
	"adventofcode2019/input"
	"testing"
)

func TestPart1(t *testing.T) {
	answer := day10.Part1(input.MainFileName)
	expected := "292"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := day10.Part2(input.MainFileName)
	expected := "317"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
