package main

import (
	"adventofcode2019/day1"
	"adventofcode2019/day2"
	"adventofcode2019/input"
	"adventofcode2019/logger"
	"flag"
	"fmt"
	"strconv"
)

var day = flag.Int("day", 0, "number of the day")
var test = flag.Bool("test", false, "if to use test input")
var loglevel = flag.String("loglevel", "info", "log level")
var part = flag.Int("part", 1, "part of the task")

var days = map[string][]func(string) string{
	"day1": {day1.Part1, day1.Part2},
	"day2": {day2.Part1, day2.Part2},
}

func main() {
	flag.Parse()

	logger.Init(*loglevel)

	funcs, ok := days["day"+strconv.Itoa(*day)]
	if ok {
		res := funcs[*part-1](input.FilePath(*day, *test))
		if res != "" {
			fmt.Println(res)
		}
	} else {
		fmt.Printf("solution for a dat number %v is not found", *day)
	}
}
