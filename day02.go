package main

// ==================================================================================
// $File: day02.go $
// $Date: 20 December 2021 $
// $Revision: $
// $Creator: Marshel Helsper <helsperm@dnb.com> $
// $Notice: (c) 2021 Dun & Bradstreet Inc. $
// ==================================================================================

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type position struct {
	x     int
	depth int
	aim   int
}

func NewPosition() position {
	return position{0, 0, 0}
}

func Day02Part1() {
	inputFile, err := os.Open("input/day02.txt")
	if err != nil {
		log.Fatalln("Could not open input file.", err)
	}
	defer inputFile.Close()

	pos := NewPosition()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		direction := parts[0]
		magnatude, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalln("Unable to convert string to int.")
		}

		if direction == "forward" {
			pos.x += magnatude
		} else if direction == "down" {
			pos.depth += magnatude
		} else if direction == "up" {
			pos.depth -= magnatude
		} else {
			log.Fatalln("Unknown direction:", direction)
		}
	}

	fmt.Println(pos.x * pos.depth)
}

func Day02Part2() {
	inputFile, err := os.Open("input/day02.txt")
	if err != nil {
		log.Fatalln("Could not open input file.", err)
	}
	defer inputFile.Close()

	pos := NewPosition()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		direction := parts[0]
		magnatude, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalln("Unable to convert string to int.")
		}

		if direction == "forward" {
			pos.x += magnatude
			pos.depth += pos.aim * magnatude
		} else if direction == "down" {
			pos.aim += magnatude
		} else if direction == "up" {
			pos.aim -= magnatude
		} else {
			log.Fatalln("Unknown direction:", direction)
		}
	}

	fmt.Printf("%+v\n", pos)
	fmt.Println(pos.x * pos.depth)
}
