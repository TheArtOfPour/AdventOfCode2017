package main

import (
	"fmt"
	"strings"
)

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
