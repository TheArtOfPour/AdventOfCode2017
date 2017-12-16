package main

import "fmt"

func advent1A(test string) (int, error) {
	first := 0
	prev := 0
	total := 0
	for i, rune := range test {
		intRune := int(rune - '0')
		fmt.Printf("%d: %d %d\n", i, intRune, total)
		if i == 0 {
			first = intRune
			prev = intRune
			continue
		}
		if prev == intRune {
			total += intRune
		}
		prev = intRune
	}
	if prev == first {
		total += first
	}

	return total, nil
}

func advent1B(test string) (int, error) {
	fmt.Printf("%s\n", test)
	offset := len(test) / 2
	opposite := 0
	total := 0
	for i, rune := range test[:offset] {
		intRune := int(rune - '0')
		circle := i + offset
		opposite = int(test[circle] - '0')
		if opposite == intRune {
			total += 2 * intRune
		}
		//fmt.Printf("%d: %d %d\n", intRune, opposite, total)
	}

	return total, nil
}
