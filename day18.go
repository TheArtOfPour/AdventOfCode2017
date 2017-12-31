package main

import (
	"bufio"
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

type program struct {
	index     int
	commands  []string
	registers map[string]int
	queue     []int
}

type result struct {
	wrote      bool
	deadlocked bool
}

func run(program *program, otherProgram *program) result {
	if program.index >= len(program.commands) {
		fmt.Print("OOB")
		return result{false, false}
	}
	command := program.commands[program.index]
	parts := strings.Split(command, " ")
	wrote := false
	deadlocked := false
	switch parts[0] {
	case "set":
		numeric, err := strconv.Atoi(parts[2])
		if err != nil {
			program.registers[parts[1]] = program.registers[parts[2]]
		} else {
			program.registers[parts[1]] = numeric
		}
	case "add":
		numeric, err := strconv.Atoi(parts[2])
		if err != nil {
			program.registers[parts[1]] += program.registers[parts[2]]
		} else {
			program.registers[parts[1]] += numeric
		}
	case "mul":
		numeric, err := strconv.Atoi(parts[2])
		if err != nil {
			program.registers[parts[1]] *= program.registers[parts[2]]
		} else {
			program.registers[parts[1]] *= numeric
		}
	case "mod":
		numeric, err := strconv.Atoi(parts[2])
		if err != nil {
			program.registers[parts[1]] %= program.registers[parts[2]]
		} else {
			program.registers[parts[1]] %= numeric
		}
	case "snd":
		numeric, err := strconv.Atoi(parts[1])
		if err != nil {
			otherProgram.queue = append(otherProgram.queue, program.registers[parts[1]])
		} else {
			otherProgram.queue = append(otherProgram.queue, numeric)
		}
		wrote = true
	case "rcv":
		if len(program.queue) == 0 {
			deadlocked = true
		} else {
			program.registers[parts[1]] = program.queue[0]
			//fmt.Printf("Before %v", program.queue)
			program.queue = program.queue[1:]
			//fmt.Printf("After %v", program.queue)
		}
	case "jgz":
		numeric1, err := strconv.Atoi(parts[1])
		if err != nil {
			if program.registers[parts[1]] > 0 {
				numeric, err := strconv.Atoi(parts[2])
				if err != nil {
					program.index += program.registers[parts[2]] - 1
				} else {
					program.index += numeric - 1
				}
			}
		} else {
			if numeric1 > 0 {
				numeric, err := strconv.Atoi(parts[2])
				if err != nil {
					program.index += program.registers[parts[2]] - 1
				} else {
					program.index += numeric - 1
				}
			}
		}
	}
	//fmt.Printf("%d: %s -> %v\n", program.index, command, program.registers)
	if !deadlocked {
		program.index++
	}
	return result{wrote, deadlocked}
}

func advent18B(test string) int {
	total := 0
	var commands []string
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		s := scanner.Text()
		commands = append(commands, s)
	}
	program0 := program{0, commands, map[string]int{}, []int{}}
	program1 := program{0, commands, map[string]int{"p": 1}, []int{}}
	for {
		result0 := run(&program0, &program1)
		result1 := run(&program1, &program0)
		// fmt.Printf("p0:%v q:%d i:%d c: %s | %v\np1:%v q:%d i:%d c: %s | %v\n\n\n",
		// 	result0, len(program0.queue), program0.index, program0.commands[program0.index], program0.registers,
		// 	result1, len(program1.queue), program1.index, program1.commands[program1.index], program1.registers)
		if result1.wrote {
			total++
		}
		if result0.deadlocked && result1.deadlocked {
			//fmt.Print("DEADLOCKED\n")
			break
		}
	}
	return total
}
