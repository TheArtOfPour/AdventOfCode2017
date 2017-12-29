package main

import (
	"sort"
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

// SortString sorts a string
func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func main() {
	// fileContents, _ := ioutil.ReadFile("./inputs/day22.txt")
	// input := string(fileContents)
	// //numeric, _ := strconv.Atoi(input)
	// out := advent22B(input)
	// fmt.Printf("Result %v\n", out)
	justforfun22()
}
