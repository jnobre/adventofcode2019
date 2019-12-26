package day1

import (
	"adventofcode2019/input"
	"log"
	"strconv"
)

// Part1 of puzzle
func Part1(filename string) string {
	var sum int64

	input.ReadFile(filename, func(line string) {
		mass, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
		sum += CalcFuel(mass)
	})

	return strconv.FormatInt(sum, 10)
}

// Part2 of puzzle
func Part2(filename string) string {
	var sum int64

	input.ReadFile(filename, func(line string) {
		mass, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
		sum += CalcFuelIncludingWeightOfFuel(mass)
	})

	return strconv.FormatInt(sum, 10)
}

//CalcFuel: part1
func CalcFuel(mass int64) int64 {
	v := (mass / 3) - 2

	if v <= 0 {
		return 0
	}

	return v
}

//CalcFuelIncludingWeightOfFuel: part2
func CalcFuelIncludingWeightOfFuel(mass int64) int64 {
	v := (mass / 3) - 2

	if v <= 0 {
		return 0
	}

	return v + CalcFuelIncludingWeightOfFuel(v)
}
