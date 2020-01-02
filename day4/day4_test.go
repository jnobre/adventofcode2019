package day4_test

import (
	"adventofcode2019/day4"
	"adventofcode2019/input"
	"testing"
)

func TestPart1(t *testing.T) {
	answer := day4.Part1(input.MainFileName)
	expected := ""
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := day4.Part2(input.MainFileName)
	expected := ""
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
