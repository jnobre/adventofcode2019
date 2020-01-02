package day3_test

import (
	"adventofcode2019/day3"
	"adventofcode2019/input"
	"testing"
)

func TestPart1(t *testing.T) {
	answer := day3.Part1(input.MainFileName)
	expected := "1211"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := day3.Part2(input.MainFileName)
	expected := "101386"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
