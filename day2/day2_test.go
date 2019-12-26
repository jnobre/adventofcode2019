package day2_test

import (
	"adventofcode2019/day2"
	"adventofcode2019/input"
	"testing"
)

func TestPart1(t *testing.T) {
	answer := day2.Part1(input.MainFileName)
	expected := "2894520"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := day2.Part2(input.MainFileName)
	expected := "9342"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
