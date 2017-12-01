package main

import (
	"fmt"
)

func main() {

}

func advent(test string) (int, error) {
	fmt.Printf("%s\n", test)
	offset := len(test) / 2
	//first := 0
	opposite := 0
	total := 0
	for i, rune := range test {
		intRune := int(rune - '0')
		circle := i + offset
		if i+offset >= len(test) {
			break
		}
		opposite = int(test[circle] - '0')
		if opposite == intRune {
			total += 2 * intRune
		}
		fmt.Printf("%d: %d %d\n", intRune, opposite, total)
	}

	return total, nil
}
