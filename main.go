package main

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"os"
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
	fileContents, _ := ioutil.ReadFile("./inputs/day14.txt")
	input := string(fileContents)
	out := advent14B(input)
	fmt.Printf("Result %d\n", out)
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

func advent5A(test string) int {
	steps := 0
	index := 0
	instructionStrings := strings.Split(test, " ")
	var instructions []int
	for _, s := range instructionStrings {
		d, _ := strconv.Atoi(s)
		instructions = append(instructions, d)
	}

	for {
		if index > len(instructions)-1 || index < 0 {
			break
		}
		instruction := instructions[index]
		//fmt.Printf("Current instruction %d index %d\n", instruction, index)
		instructions[index]++
		index += instruction
		steps++
	}
	return steps
}

func advent5B(test string) int {
	steps := 0
	index := 0
	instructionStrings := strings.Split(test, " ")
	var instructions []int
	for _, s := range instructionStrings {
		d, _ := strconv.Atoi(s)
		instructions = append(instructions, d)
	}

	for {
		if index > len(instructions)-1 || index < 0 {
			break
		}
		instruction := instructions[index]
		//fmt.Printf("Current instruction %d index %d\n", instruction, index)
		if instruction >= 3 {
			instructions[index]--
		} else {
			instructions[index]++
		}
		index += instruction
		steps++
	}
	return steps
}

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

//Node bi-directional tree node
type Node struct {
	parent   *Node
	name     string
	weight   int
	children []Node
}

func advent7A(test string) string {
	var nodes []Node
	scanner := bufio.NewScanner(strings.NewReader(test))
	var temp Node
	for scanner.Scan() {
		s := scanner.Text()
		nameAndWeight := s
		if strings.Contains(s, "->") {
			stringParts := strings.Split(s, " -> ")
			nameAndWeight = stringParts[0]
			childrenNames := strings.Split(stringParts[1], ", ")
			for _, childName := range childrenNames {
				var childNode Node
				childNode.name = childName
				temp.children = append(temp.children, childNode)
			}
		}
		nameParts := strings.Split(nameAndWeight, " (")
		weight, _ := strconv.Atoi(strings.TrimRight(nameParts[1], ")"))
		temp.name = nameParts[0]
		temp.weight = weight
		nodes = append(nodes, temp)
		//fmt.Printf("%v\n", temp)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	root := findRoot(nodes)
	//fmt.Printf("%v\n", root)

	return root.name
}

func advent7B(test string) int {
	var nodes []Node
	nodeMap := make(map[string]Node)
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		var temp Node
		s := scanner.Text()
		nameAndWeight := s
		if strings.Contains(s, "->") {
			stringParts := strings.Split(s, " -> ")
			nameAndWeight = stringParts[0]
			childrenNames := strings.Split(stringParts[1], ", ")
			for _, childName := range childrenNames {
				var childNode Node
				childNode.name = childName
				temp.children = append(temp.children, childNode)
			}
		}
		nameParts := strings.Split(nameAndWeight, " (")
		weight, _ := strconv.Atoi(strings.TrimRight(nameParts[1], ")"))
		temp.name = nameParts[0]
		temp.weight = weight
		nodes = append(nodes, temp)
		nodeMap[temp.name] = temp
		//fmt.Printf("%v\n", temp)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	root := findRoot(nodes)
	buildTree(&root, nodeMap)
	//fmt.Printf("%v\n", nodeMap)
	result := updateWeights(&root, 0)

	return result
}

func updateWeights(root *Node, result int) int {
	if result != 0 {
		return result
	}
	if len(root.children) == 0 {
		return root.weight
	}
	var weights []int
	sum := 0
	for _, child := range root.children {
		weight := updateWeights(&child, 0)
		weights = append(weights, weight)
		sum += weight
	}
	for i, weight := range weights {
		count := 0
		for _, weightSearch := range weights {
			if weight == weightSearch {
				count++
			}
		}
		if count == 1 {
			offset := 0
			if i == 0 {
				offset = weight - weights[i+1]
			} else {
				offset = weight - weights[i-1]
			}
			//fmt.Printf("\ninvalid weight! : %d %d %d %d %d\n", root.children[i].weight, offset, weight, root.weight, sum)
			panic(root.children[i].weight - offset)
		}
	}
	root.weight += sum
	sum = root.weight
	//fmt.Printf("\nweights : %v\n", weights)
	return sum
}

func buildTree(root *Node, nodeMap map[string]Node) {
	deroot := *root
	newRoot := nodeMap[deroot.name]
	root = &newRoot
	//fmt.Printf("\nroot : %v\n", root)
	if len(root.children) == 0 {
		//fmt.Printf("leaf : %v\n", root)
	} else {
		for i := range root.children {
			//fmt.Printf("child : %v\n", root.children[i])
			root.children[i] = nodeMap[root.children[i].name]
			root.children[i].parent = root
			buildTree(&root.children[i], nodeMap)
		}
	}
}

func findRoot(nodes []Node) Node {
	for _, node := range nodes {
		isParent := true
		for _, parent := range nodes {
			for _, child := range parent.children {
				if child.name == node.name {
					isParent = false
				}
			}
		}
		if isParent {
			return node
		}
	}
	return nodes[0]
}
