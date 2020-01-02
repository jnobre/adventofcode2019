package day4

import (
	"sort"
	"strconv"
	"strings"
)

func solv() (int, int) {
	start, stop := 245318, 765747
	count1, count2 := 0, 0
	for i := start; i <= stop; i++ {
		s := strconv.Itoa(i)
		sorted := sort.SliceIsSorted(s, func(i, j int) bool { return s[i] < s[j] })
		if sorted {
			inc1 := 0
			for i := 0; i < len(s)-1; {
				j := strings.LastIndexByte(s, s[i])
				if i != j {
					inc1 = 1
					if i+1 == j {
						count2++
						break
					}
				}
				i = j + 1
			}
			count1 += inc1
			continue
		}
	}
	return count1, count2
}

func Part1(_ string) string {
	solv, _ := solv()
	return strconv.Itoa(solv)
}

func Part2(_ string) string {
	_, solv := solv()
	return strconv.Itoa(solv)
}
