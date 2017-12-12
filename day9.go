package main

import (
	"fmt"
)

func advent9A(test string) int {
	skip, garbage := false, false
	multiplier, sum := 0, 0
	for _, rune := range test {
		// !
		if skip {
			skip = false
			continue
		}
		if rune == '!' {
			skip = true
			continue
		}
		// <>
		if rune == '>' {
			garbage = false
			continue
		}
		if garbage {
			continue
		}
		if rune == '<' {
			garbage = true
			continue
		}
		// {}
		if rune == '{' {
			multiplier++
			sum += 1 * multiplier
		}
		if rune == '}' {
			multiplier--
		}
		//fmt.Printf("%s\n", string(rune))
	}
	return sum
}

func advent9B(test string) int {
	skip, garbage := false, false
	sum := 0
	for _, rune := range test {
		// !
		if skip {
			skip = false
			continue
		}
		if rune == '!' {
			skip = true
			continue
		}
		// <>
		if rune == '>' {
			garbage = false
			continue
		}
		if garbage {
			sum++
			continue
		}
		if rune == '<' {
			garbage = true
			continue
		}
		fmt.Printf("%s\n", string(rune))
	}
	return sum
}
