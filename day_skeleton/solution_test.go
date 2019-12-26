package day_skeleton_test

import (
	"adventofcode2019/day_skeleton"
	"adventofcode2019/input"
	"testing"
)

func TestPart1(t *testing.T) {
	answer := day_skeleton.Part1(input.MainFileName)
	expected := ""
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}

func TestPart2(t *testing.T) {
	answer := day_skeleton.Part2(input.MainFileName)
	expected := ""
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
