package day3

import (
	"adventofcode2019/input"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

type Operation struct {
	action string
	step   int
}

func Part1(filename string) string {
	wires := input.ReadStringArraysFromFile(filename, 2)

	minDistanse := int64(math.MaxInt64)
	points := make(map[Point]int)

	for _, wire := range wires {
		var point Point
		visitedPoints := make(map[Point]bool)

		for _, op := range wire {
			operation := parsedOperation(op)

			for i := 1; i <= operation.step; i++ {
				newOp := Operation{action: operation.action, step: 1}
				point = applyOperation(point, newOp)

				if !visitedPoints[point] {
					visitedPoints[point] = true
					points[point] += 1

					if points[point] == len(wires) {
						dist := int64(math.Abs(float64(point.x)) + math.Abs(float64(point.y)))
						if dist < minDistanse {
							minDistanse = dist
						}
					}
				}
			}
		}
	}

	if minDistanse == math.MaxInt64 {
		fmt.Println("intersections were not found")
		return ""
	}
	return strconv.FormatInt(minDistanse, 10)
}

func Part2(filename string) string {
	wires := input.ReadStringArraysFromFile(filename, 2)

	minSteps := int64(math.MaxInt64)
	points := make(map[Point][]int)

	for _, wire := range wires {
		var point Point
		visitedPoints := make(map[Point]bool)
		steps := 0

		for _, strOperation := range wire {
			operation := parsedOperation(strOperation)

			for i := 1; i <= operation.step; i++ {
				steps += 1
				newOp := Operation{action: operation.action, step: 1}
				point = applyOperation(point, newOp)

				if !visitedPoints[point] {
					visitedPoints[point] = true
					points[point] = append(points[point], steps)

					if len(points[point]) == len(wires) {
						var sumSteps int64
						for _, item := range points[point] {
							sumSteps += int64(item)
						}
						if sumSteps < minSteps {
							minSteps = sumSteps
						}
					}
				}
			}
		}
	}

	if minSteps == math.MaxInt64 {
		fmt.Println("intersections were not found")
		return ""
	}
	return strconv.FormatInt(minSteps, 10)
}

func parsedOperation(op string) Operation {
	action := strings.Split(op, "")[0]

	num, err := strconv.Atoi(op[1:])
	if err != nil {
		log.Fatal(err.Error())
	}

	return Operation{action, num}
}

func applyOperation(point Point, op Operation) Point {
	switch op.action {
	case "R":
		point.x += op.step
	case "U":
		point.y += op.step
	case "D":
		point.y -= op.step
	case "L":
		point.x -= op.step
	}

	return point
}
