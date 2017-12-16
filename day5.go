package main

import (
	"strconv"
	"strings"
)

func advent5A(test string) int {
	steps := 0
	index := 0
	instructionStrings := strings.Split(test, " ")
	var instructions []int
	for _, s := range instructionStrings {
		d, _ := strconv.Atoi(s)
		instructions = append(instructions, d)
	}

	for {
		if index > len(instructions)-1 || index < 0 {
			break
		}
		instruction := instructions[index]
		//fmt.Printf("Current instruction %d index %d\n", instruction, index)
		instructions[index]++
		index += instruction
		steps++
	}
	return steps
}

func advent5B(test string) int {
	steps := 0
	index := 0
	instructionStrings := strings.Split(test, " ")
	var instructions []int
	for _, s := range instructionStrings {
		d, _ := strconv.Atoi(s)
		instructions = append(instructions, d)
	}

	for {
		if index > len(instructions)-1 || index < 0 {
			break
		}
		instruction := instructions[index]
		//fmt.Printf("Current instruction %d index %d\n", instruction, index)
		if instruction >= 3 {
			instructions[index]--
		} else {
			instructions[index]++
		}
		index += instruction
		steps++
	}
	return steps
}
