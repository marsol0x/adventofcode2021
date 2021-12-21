package main

// ==================================================================================
// $File: day03.go $
// $Date: 21 December 2021 $
// $Revision: $
// $Creator: Marshel Helsper <helsperm@dnb.com> $
// $Notice: (c) 2021 Dun & Bradstreet Inc. $
// ==================================================================================

import (
	"fmt"
	"strconv"
	"log"
)

func Day03Part1() {
	//mask := int32(0x001F) // NOTE(20211221 helsperm): Sample mask
	mask := int32(0x0FFF) // NOTE(20211221 helsperm): Actual mask

	inputLines := ReadLinesFromFile("input/day03.txt")

	var (
		gamma int32
		epsilon int32

		bitCounts [32]int
	)

	for _, line := range(inputLines) {
		binNum, err := strconv.ParseInt(line, 2, 32)
		if err != nil {
			log.Fatalln("Could not parse string to int16.", err)
		}

		m := int32(1)
		for i := 32; i >= 0; i-- {
			if (int32(binNum) & m) > 0 {
				bitCounts[i-1]++
			}
			m <<= 1
		}
	}

	for _, bc := range(bitCounts) {
		var v int32
		if bc > (len(inputLines) / 2) {
			v = 1
		}

		gamma <<= 1
		gamma |= v
	}

	epsilon = gamma ^ mask

	fmt.Println(gamma, epsilon, gamma * epsilon)
}
