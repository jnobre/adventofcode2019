package day6_test

import (
	"adventofcode2019/day6"
	"adventofcode2019/input"
	"testing"
)

func TestPart1(t *testing.T) {
	answer := day6.Part1(input.MainFileName)
	expected := ""
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := day6.Part2(input.MainFileName)
	expected := ""
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
