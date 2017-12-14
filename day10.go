package main

import (
	"fmt"
	"strconv"
	"strings"
)

func createNumberedSlice(size int) []int {
	var s []int
	for i := 0; i < size; i++ {
		s = append(s, i)
	}
	return s
}

func reverseSlice(a []int, s int, l int) {
	// build temporary slice from start to start + length
	var temp []int
	if s+l < len(a) {
		temp = a[s : s+l]
	} else {
		temp = a[s:]
		for _, t := range a[:s+l-len(a)] {
			temp = append(temp, t)
		}
	}
	//fmt.Printf("temp: %v\n", temp)
	// reverse the temporary slice
	for i := len(temp)/2 - 1; i >= 0; i-- {
		opp := len(temp) - 1 - i
		temp[i], temp[opp] = temp[opp], temp[i]
	}
	//fmt.Printf("reversed: %v\n", temp)
	// overwrite the original range with the reversed slice
	for i, t := range temp {
		if s+i < len(a) {
			a[s+i] = t
		} else {
			a[s+i-len(a)] = t
		}
	}
}

func advent10A(test string) int {
	// 5 for test, 256 for actual
	size := 5 //256
	list := createNumberedSlice(size)

	inputs := strings.Split(test, ",")
	i := 0
	for skip, inputString := range inputs {
		//fmt.Printf("%v\n", list)
		input, _ := strconv.Atoi(inputString)
		//fmt.Printf("%d: %d [%d]\n", i, input, list[i])
		reverseSlice(list, i, input)
		i += input + skip
		if i > len(list) {
			i -= len(list)
		}
	}
	//fmt.Printf("%v\n", list)
	//fmt.Printf("%s\n", string(rune))
	return list[0] * list[1]
}

func advent10B(test string) string {
	size := 256
	list := createNumberedSlice(size)

	// string to ascii byte slice
	inputs := []byte(test)
	inputs = append(inputs, 17)
	inputs = append(inputs, 31)
	inputs = append(inputs, 73)
	inputs = append(inputs, 47)
	inputs = append(inputs, 23)
	//fmt.Printf("%v\n", inputs)
	skip := 0
	i := 0
	for j := 0; j < 64; j++ {
		//fmt.Printf("loop %d\n", j)
		for _, input := range inputs {
			//fmt.Printf("s:%d, i:%d, input:%d\n", skip, i, input)
			//fmt.Printf("[%d]\n", list[i])
			reverseSlice(list, i, int(input))
			i += int(input) + skip
			for i >= len(list) {
				i -= len(list)
			}
			skip++
		}
	}
	var hash []string
	for k := 0; k < 16; k++ {
		hashbyte := list[k*16]
		for l := (k * 16) + 1; l < (k*16)+16; l++ {
			hashbyte = hashbyte ^ list[l]
		}
		//fmt.Printf("%v\n", hashbyte)
		h := fmt.Sprintf("%02x", hashbyte)
		hash = append(hash, h)
	}
	//fmt.Printf("%v\n", hash)
	return strings.Join(hash, "")
}
