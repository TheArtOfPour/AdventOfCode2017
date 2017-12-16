package main

import (
	"crypto/sha1"
	"strconv"
	"strings"
)

func advent6A(test string) int {
	steps := 0
	var combinations [][]byte
	bankStrings := strings.Split(test, " ")
	var banks []int
	for _, s := range bankStrings {
		d, _ := strconv.Atoi(s)
		banks = append(banks, d)
	}

	//fmt.Printf("%v\n", banks)
	for {
		h := sha1.New()
		bankStrings = []string{}
		for _, d := range banks {
			s := strconv.Itoa(d)
			bankStrings = append(bankStrings, s)
		}
		h.Write([]byte(strings.Join(bankStrings, " ")))
		hash := h.Sum(nil)
		//fmt.Printf("%x\n", hash)
		for _, combination := range combinations {
			hashString := string(hash)
			combinationString := string(combination)
			if hashString == combinationString {
				return steps
			}
		}
		combinations = append(combinations, hash)

		maxIndex := 0
		for index, bank := range banks[1:] {
			if bank > banks[maxIndex] {
				maxIndex = index + 1
				//fmt.Printf("%d > %d %d\n", bank, banks[maxIndex], index+1)
			}
		}
		//fmt.Printf("maxIndex: %d\n", maxIndex)
		blocks := banks[maxIndex]
		banks[maxIndex] = 0
		currentIndex := maxIndex + 1
		for i := 0; i < blocks; i++ {
			if currentIndex >= len(banks) {
				currentIndex = 0
			}
			banks[currentIndex]++
			currentIndex++
		}
		//fmt.Printf("%v\n", banks)
		steps++
	}
}

func advent6B(test string) int {
	steps := 0
	var combinations [][]byte
	bankStrings := strings.Split(test, " ")
	var banks []int
	for _, s := range bankStrings {
		d, _ := strconv.Atoi(s)
		banks = append(banks, d)
	}

	//fmt.Printf("%v\n", banks)
	for {
		h := sha1.New()
		bankStrings = []string{}
		for _, d := range banks {
			s := strconv.Itoa(d)
			bankStrings = append(bankStrings, s)
		}
		h.Write([]byte(strings.Join(bankStrings, " ")))
		hash := h.Sum(nil)
		//fmt.Printf("%x\n", hash)
		for i, combination := range combinations {
			//fmt.Printf("%d\n", len(combinations)-i)
			hashString := string(hash)
			combinationString := string(combination)
			if hashString == combinationString {
				return len(combinations) - i
			}
		}
		combinations = append(combinations, hash)

		maxIndex := 0
		for index, bank := range banks[1:] {
			if bank > banks[maxIndex] {
				maxIndex = index + 1
				//fmt.Printf("%d > %d %d\n", bank, banks[maxIndex], index+1)
			}
		}
		//fmt.Printf("maxIndex: %d\n", maxIndex)
		blocks := banks[maxIndex]
		banks[maxIndex] = 0
		currentIndex := maxIndex + 1
		for i := 0; i < blocks; i++ {
			if currentIndex >= len(banks) {
				currentIndex = 0
			}
			banks[currentIndex]++
			currentIndex++
		}
		//fmt.Printf("%v\n", banks)
		steps++
	}
}
