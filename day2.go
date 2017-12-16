package main

import (
	"fmt"
	"strconv"
	"strings"
)

func advent2A(test string) int {
	fmt.Printf("%s\n", test)
	min := 999999999
	max := 0
	sum := 0
	rows := strings.Split(test, ".")
	for _, row := range rows {
		min = 999999999
		max = 0
		numbers := strings.Split(row, " ")
		for _, number := range numbers {
			i, _ := strconv.Atoi(number)
			fmt.Printf("%d\n", i)
			if i < min {
				min = i
			}
			if i > max {
				max = i
			}
		}
		sum += max - min
		//fmt.Printf("%s: %d %d %d\n", row, min, max, sum)
	}

	return sum
}

func advent2B(test string) int {
	fmt.Printf("%s\n", test)
	sum := 0
	rows := strings.Split(test, ".")
	for _, row := range rows {
		numbers := strings.Split(row, " ")
		for _, number1 := range numbers {
			found := false
			for _, number2 := range numbers {
				n1, _ := strconv.Atoi(number1)
				n2, _ := strconv.Atoi(number2)
				if n1 == n2 {
					continue
				}
				if n1 > n2 {
					if n1%n2 == 0 {
						sum += n1 / n2
						found = true
						break
					}
				}
				if n2%n1 == 0 {
					sum += n2 / n1
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		//fmt.Printf("%s: %d %d %d\n", row, min, max, sum)
	}

	return sum
}
