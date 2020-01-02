package day9_test

import (
	"adventofcode2019/day9"
	"adventofcode2019/input"
	"testing"
)

func TestPart1(t *testing.T) {
	answer := day9.Part1(input.MainFileName)
	expected := "2350741403"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := day9.Part2(input.MainFileName)
	expected := "53088"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
