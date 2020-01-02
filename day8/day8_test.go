package day8_test

import (
	"adventofcode2019/day8"
	"adventofcode2019/input"
	"testing"
)

func TestPart1(t *testing.T) {
	answer := day8.Part1(input.MainFileName)
	expected := "2500"
	if answer != expected {
		t.Errorf("Error, expected %s got %s", expected, answer)
	}
}
