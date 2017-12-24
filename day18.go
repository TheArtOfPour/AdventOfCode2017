package main

import (
	"fmt"
	"strconv"
	"strings"
)

func advent18A(test string) int {
	lastSound := 0
	registers := make(map[string]int)
	commands := strings.Split(test, "\r\n")
	i := 0
	for {
		command := commands[i]
		stringParts := strings.Split(command, " ")
		switch stringParts[0] {
		case "snd":
			lastSound = registers[stringParts[1]]
		case "set":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] = registers[stringParts[2]]
			} else {
				registers[stringParts[1]] = numeric
			}
		case "add":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] += registers[stringParts[2]]
			} else {
				registers[stringParts[1]] += numeric
			}
		case "mul":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] *= registers[stringParts[2]]
			} else {
				registers[stringParts[1]] *= numeric
			}
		case "mod":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] %= registers[stringParts[2]]
			} else {
				registers[stringParts[1]] %= numeric
			}
		case "rcv":
			if registers[stringParts[1]] != 0 {
				return lastSound
			}
		case "jgz":
			if registers[stringParts[1]] > 0 {
				numeric, err := strconv.Atoi(stringParts[2])
				if err != nil {
					i += registers[stringParts[2]] - 1
				} else {
					i += numeric - 1
				}
			}
		default:
			panic(fmt.Sprintf("invalid command %s", stringParts[0]))
		}
		fmt.Printf("%d: %s -> %v\n", i, command, registers)
		i++
	}
}

func advent18B(test string) int {
	lastSound := 0
	registers := make(map[string]int)
	commands := strings.Split(test, "\r\n")
	i := 0
	for {
		command := commands[i]
		stringParts := strings.Split(command, " ")
		switch stringParts[0] {
		case "snd":
			lastSound = registers[stringParts[1]]
		case "set":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] = registers[stringParts[2]]
			} else {
				registers[stringParts[1]] = numeric
			}
		case "add":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] += registers[stringParts[2]]
			} else {
				registers[stringParts[1]] += numeric
			}
		case "mul":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] *= registers[stringParts[2]]
			} else {
				registers[stringParts[1]] *= numeric
			}
		case "mod":
			numeric, err := strconv.Atoi(stringParts[2])
			if err != nil {
				registers[stringParts[1]] %= registers[stringParts[2]]
			} else {
				registers[stringParts[1]] %= numeric
			}
		case "rcv":
			if registers[stringParts[1]] != 0 {
				return lastSound
			}
		case "jgz":
			if registers[stringParts[1]] > 0 {
				numeric, err := strconv.Atoi(stringParts[2])
				if err != nil {
					i += registers[stringParts[2]] - 1
				} else {
					i += numeric - 1
				}
			}
		default:
			panic(fmt.Sprintf("invalid command %s", stringParts[0]))
		}
		fmt.Printf("%d: %s -> %v\n", i, command, registers)
		i++
	}
}
