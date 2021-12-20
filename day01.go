package main
// ==================================================================================
// $File: day01.go $
// $Date: 20 December 2021 $
// $Revision: $
// $Creator: Marshel Helsper <helsperm@dnb.com> $
// $Notice: (c) 2021 Dun & Bradstreet Inc. $
// ==================================================================================

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func Day01Part1() {
	inputFile, err := os.Open("input/day01.txt")
	if err != nil {
		log.Fatalln("Could not open input file.", err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	prev := math.MaxInt32
	count := 0
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatalln("Could not convert input string data to an integer.", err)
		}

		if num > prev {
			count++
		}

		prev = num
	}

	fmt.Println("Answer:", count)
}
