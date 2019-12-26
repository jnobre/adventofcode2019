package day1_test

import (
	"adventofcode2019/day1"
	"adventofcode2019/input"
	"testing"
)

func TestPart1(t *testing.T) {
	answer := day1.Part1(input.MainFileName)
	expected := "3382136"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := day1.Part2(input.MainFileName)
	expected := "5070314"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
