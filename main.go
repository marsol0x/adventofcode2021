package main

// ==================================================================================
// $File: main.go $
// $Date: 20 December 2021 $
// $Revision: $
// $Creator: Marshel Helsper <helsperm@dnb.com> $
// $Notice: (c) 2021 Dun & Bradstreet Inc. $
// ==================================================================================

import (
	"bufio"
	"log"
	"os"
)

func ReadLinesFromFile(path string) []string {
	result := make([]string, 0)
	inputFile, err := os.Open(path)
	if err != nil {
		log.Fatalln("Could not open input file:", path, err)
	}

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	return result
}

func main() {
	//Day01Part1()
	//Day01Part2()
	//Day02Part1()
	//Day02Part2()
	Day03Part1()
}
