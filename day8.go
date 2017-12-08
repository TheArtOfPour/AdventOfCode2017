package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func advent8A(test string) int {
	registers := make(map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		s := scanner.Text()
		// b inc 5 if a > 1
		// 0  1  2  3 4 5 6
		instructionParts := strings.Split(s, " ")

		// check conditional
		checkRegister := instructionParts[4]
		if _, ok := registers[checkRegister]; !ok {
			registers[checkRegister] = 0
		}
		checkCondition := instructionParts[5]
		checkValue, _ := strconv.Atoi(instructionParts[6])
		checkPass := false
		switch checkCondition {
		case ">":
			checkPass = registers[checkRegister] > checkValue
		case "<":
			checkPass = registers[checkRegister] < checkValue
		case ">=":
			checkPass = registers[checkRegister] >= checkValue
		case "<=":
			checkPass = registers[checkRegister] <= checkValue
		case "==":
			checkPass = registers[checkRegister] == checkValue
		case "!=":
			checkPass = registers[checkRegister] != checkValue
		default:
			panic(fmt.Sprintf("invalid condition %s", checkCondition))
		}
		if !checkPass {
			continue
		}

		// conditional passed, modify register
		register := instructionParts[0]
		if _, ok := registers[register]; !ok {
			registers[register] = 0
		}
		operation := instructionParts[1]
		value, _ := strconv.Atoi(instructionParts[2])
		switch operation {
		case "inc":
			registers[register] += value
		case "dec":
			registers[register] -= value
		default:
			panic(fmt.Sprintf("invalid operation %s", operation))
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	max := 0
	for _, v := range registers {
		if v > max {
			max = v
		}
	}
	return max
}

func advent8B(test string) int {
	max := 0
	registers := make(map[string]int)
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		s := scanner.Text()
		// b inc 5 if a > 1
		// 0  1  2  3 4 5 6
		instructionParts := strings.Split(s, " ")

		// check conditional
		checkRegister := instructionParts[4]
		checkCondition := instructionParts[5]
		checkValue, _ := strconv.Atoi(instructionParts[6])
		checkPass := false
		switch checkCondition {
		case ">":
			checkPass = registers[checkRegister] > checkValue
		case "<":
			checkPass = registers[checkRegister] < checkValue
		case ">=":
			checkPass = registers[checkRegister] >= checkValue
		case "<=":
			checkPass = registers[checkRegister] <= checkValue
		case "==":
			checkPass = registers[checkRegister] == checkValue
		case "!=":
			checkPass = registers[checkRegister] != checkValue
		default:
			panic(fmt.Sprintf("invalid condition %s", checkCondition))
		}
		if !checkPass {
			continue
		}

		// conditional passed, modify register
		register := instructionParts[0]
		operation := instructionParts[1]
		value, _ := strconv.Atoi(instructionParts[2])
		switch operation {
		case "inc":
			registers[register] += value
		case "dec":
			registers[register] -= value
		default:
			panic(fmt.Sprintf("invalid operation %s", operation))
		}
		for _, v := range registers {
			if v > max {
				max = v
			}
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	return max
}
