package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {

}

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

func advent3A(test int) int {
	fmt.Printf("%d\n", test)
	half := 572 / 2
	var twoD [572][572]int
	x := half
	y := half
	number := 1
	twoD[x][y] = number
	done := false
	ru := true
	for i := 1; i <= test; i++ {
		if ru {
			// R
			for j := 0; j < i; j++ {
				if number >= test {
					done = true
					break
				}
				x++
				twoD[x][y] = number
				number++
				fmt.Printf("R | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			// U
			for j := 0; j < i; j++ {
				if number >= test {
					done = true
					break
				}
				y--
				twoD[x][y] = number
				number++
				fmt.Printf("U | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			ru = false
		} else {
			// L
			for j := 0; j < i; j++ {
				if number >= test {
					done = true
					break
				}
				x--
				twoD[x][y] = number
				number++
				fmt.Printf("L | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			// D
			for j := 0; j < i; j++ {
				if number >= test {
					done = true
					break
				}
				y++
				twoD[x][y] = number
				number++
				fmt.Printf("D | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			ru = true
		}

	}
	// R1 U1 L2 D2 R3 U3 L4 D4
	//fmt.Printf("%s: %d %d %d\n", row, min, max, sum)

	x = x - half
	if x < 0 {
		x *= -1
	}
	y = y - half
	if y < 0 {
		y *= -1
	}

	return x + y
}

func advent3B(test int) int {
	fmt.Printf("%d\n", test)
	half := 574 / 2
	var twoD [574][574]int
	x := half
	y := half
	number := 1
	twoD[x][y] = number
	done := false
	ru := true
	for i := 1; i <= test; i++ {
		if ru {
			// R
			for j := 0; j < i; j++ {
				x++
				number = twoD[x+1][y] + twoD[x+1][y+1] + twoD[x][y+1] + twoD[x-1][y+1] + twoD[x-1][y] + twoD[x-1][y-1] + twoD[x][y-1] + twoD[x+1][y-1]
				if number > test {
					done = true
					break
				}
				twoD[x][y] = number
				fmt.Printf("R | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			// U
			for j := 0; j < i; j++ {
				y--
				number = twoD[x+1][y] + twoD[x+1][y+1] + twoD[x][y+1] + twoD[x-1][y+1] + twoD[x-1][y] + twoD[x-1][y-1] + twoD[x][y-1] + twoD[x+1][y-1]
				if number > test {
					done = true
					break
				}
				twoD[x][y] = number
				fmt.Printf("U | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			ru = false
		} else {
			// L
			for j := 0; j < i; j++ {
				x--
				number = twoD[x+1][y] + twoD[x+1][y+1] + twoD[x][y+1] + twoD[x-1][y+1] + twoD[x-1][y] + twoD[x-1][y-1] + twoD[x][y-1] + twoD[x+1][y-1]
				if number > test {
					done = true
					break
				}
				twoD[x][y] = number
				fmt.Printf("L | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			// D
			for j := 0; j < i; j++ {
				y++
				number = twoD[x+1][y] + twoD[x+1][y+1] + twoD[x][y+1] + twoD[x-1][y+1] + twoD[x-1][y] + twoD[x-1][y-1] + twoD[x][y-1] + twoD[x+1][y-1]
				if number > test {
					done = true
					break
				}
				twoD[x][y] = number
				fmt.Printf("D | number: %d, x: %d, y: %d\n", number, x, y)
			}
			if done {
				break
			}
			ru = true
		}

	}
	// R1 U1 L2 D2 R3 U3 L4 D4
	//fmt.Printf("%s: %d %d %d\n", row, min, max, sum)

	return number
}

func advent4A(test string) int {
	count := 0
	rows := strings.Split(test, ".")
	for _, row := range rows {
		valid := true
		words := strings.Split(row, " ")
		for i, word1 := range words {
			for _, word2 := range words[i+1:] {
				if word1 == word2 {
					fmt.Printf("Duplicate %s\n", word1)
					valid = false
					break
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			fmt.Printf("Valid\n")
			count++
		}
	}
	return count
}

func advent4B(test string) int {
	count := 0
	rows := strings.Split(test, ".")
	for _, row := range rows {
		valid := true
		words := strings.Split(row, " ")
		for i, word1 := range words {
			for _, word2 := range words[i+1:] {
				if len(word1) == len(word2) {
					word1 = SortString(word1)
					word2 = SortString(word2)
					if word1 == word2 {
						fmt.Printf("Duplicate %s %s\n", word1, word2)
						valid = false
						break
					}
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			fmt.Printf("Valid\n")
			count++
		}
	}
	return count
}
