package intcode

import (
	"errors"
	"fmt"
	"log"
)

type OneStepProgram struct {
	index, relBase int
	code           []int64
	Finished       bool
}

func NewOneStepProgram(code []int64) OneStepProgram {
	return OneStepProgram{code: code}
}

func (p *OneStepProgram) Run(input int64) int64 {
	var err error
	var output []int64

	if p.Finished {
		return 0
	}
	output, p.index, p.relBase, err = Solve(p.code, []int64{input}, p.index, p.relBase, true)
	if err != nil {
		log.Fatal("failed to solve intcode", err.Error())
	}

	if p.index == -1 {
		p.Finished = true
		return 0
	}
	return output[0]
}

func Solve(a []int64, inputInstructions []int64, startIndex int, startRelBase int, stopOnOutput bool) (output []int64, index int, relBase int, err error) {
	var last, inputIndex int

	relBase = startRelBase
	index = startIndex

	for {
		last = index

		opcode, p1Index, p2Index, p3Index := parsedParameters(a, index, relBase)

		switch opcode {
		case 1:
			write(a, p3Index, a[p1Index]+a[p2Index])
			index += 4
		case 2:
			write(a, p3Index, a[p1Index]*a[p2Index])
			index += 4
		case 3:
			write(a, p1Index, inputInstructions[inputIndex])
			if inputIndex < len(inputInstructions)-1 {
				inputIndex++
			}
			index += 2
		case 4:
			output = append(output, a[p1Index])
			index += 2
			if stopOnOutput {
				return
			}
		case 5:
			if a[p1Index] != 0 {
				index = int(a[p2Index])
			} else {
				index += 3
			}
		case 6:
			if a[p1Index] == 0 {
				index = int(a[p2Index])
			} else {
				index += 3
			}
		case 7:
			if a[p1Index] < a[p2Index] {
				write(a, p3Index, 1)
			} else {
				write(a, p3Index, 0)
			}
			index += 4
		case 8:
			if a[p1Index] == a[p2Index] {
				write(a, p3Index, 1)
			} else {
				write(a, p3Index, 0)
			}
			index += 4
		case 9:
			relBase += int(a[p1Index])
			index += 2
		case 99:
			return output, -1, 0, nil
		default:
			return []int64{}, 0, 0, errors.New(fmt.Sprintf("unknown code %v", opcode))
		}

		if last == index {
			log.Fatal("loop", a[index], a)
		}
	}
}

func write(a []int64, index int, value int64) {
	a[index] = value
}

func parsedParameters(a []int64, index int, relBase int) (int, int, int, int) {
	var p1Index, p2Index, pIndex3 int

	mode3, mode2, mode1, opcode := instructions(int(a[index]))

	if index+1 < len(a) {
		p1Index = getIndexWithMode(a, index+1, mode1, relBase)
	}
	if index+2 < len(a) {
		p2Index = getIndexWithMode(a, index+2, mode2, relBase)
	}
	if index+3 < len(a) {
		pIndex3 = getIndexWithMode(a, index+3, mode3, relBase)
	}

	return opcode, p1Index, p2Index, pIndex3
}

func instructions(k int) (int, int, int, int) {
	a := k / 10000
	b := k / 1000 % 10
	c := k / 100 % 10
	return a, b, c, k % 100
}

func getIndexWithMode(a []int64, index, mode int, relBase int) int {
	switch mode {
	case 1:
		return index
	case 2:
		return int(a[index]) + relBase
	default:
		return int(a[index])
	}
}
