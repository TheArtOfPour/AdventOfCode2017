package main

import (
	"fmt"
	"strconv"
	"strings"
)

func advent23A(test string) int {
	mulTime := 0
	registers := make(map[string]int)
	commands := strings.Split(test, "\r\n")
	i := 0
	for {
		if i < 0 || i >= len(commands) {
			break
		}
		command := commands[i]
		stringParts := strings.Split(command, " ")
		switch stringParts[0] {
		case "set":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] = registers[stringParts[2]]
			} else {
				registers[stringParts[1]] = numeric
			}
		case "sub":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] -= registers[stringParts[2]]
			} else {
				registers[stringParts[1]] -= numeric
			}
		case "mul":
			mulTime++
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] *= registers[stringParts[2]]
			} else {
				registers[stringParts[1]] *= numeric
			}
		case "jnz":
			numeric, err := strconv.Atoi(stringParts[1])
			if err != nil {
				numeric = registers[stringParts[1]]
			}
			if numeric != 0 {
				numeric, err := strconv.Atoi(stringParts[2])
				if err != nil {
					numeric = registers[stringParts[2]]
				}
				i += numeric - 1
			}
		default:
			panic(fmt.Sprintf("invalid command %s", stringParts[0]))
		}
		//fmt.Printf("%d: %s -> %v\n", i, command, registers)
		i++
	}
	return mulTime
}

func advent23B(test string) int {
	registers := make(map[string]int)
	registers["a"] = 1
	commands := strings.Split(test, "\r\n")
	i := 0
	for {
		if i < 0 || i >= len(commands) {
			break
		}
		command := commands[i]
		stringParts := strings.Split(command, " ")
		switch stringParts[0] {
		case "set":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] = registers[stringParts[2]]
			} else {
				registers[stringParts[1]] = numeric
			}
		case "sub":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] -= registers[stringParts[2]]
			} else {
				registers[stringParts[1]] -= numeric
			}
		case "mul":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] *= registers[stringParts[2]]
			} else {
				registers[stringParts[1]] *= numeric
			}
		case "jnz":
			numeric, err := strconv.Atoi(stringParts[1])
			if err != nil {
				numeric = registers[stringParts[1]]
			}
			if numeric != 0 {
				numeric, err := strconv.Atoi(stringParts[2])
				if err != nil {
					numeric = registers[stringParts[2]]
				}
				i += numeric - 1
			}
		default:
			panic(fmt.Sprintf("invalid command %s", stringParts[0]))
		}
		//fmt.Printf("%d: %s -> %v h:%d\n", i, command, registers, registers["h"])
		i++
	}
	return registers["h"]
}
