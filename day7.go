package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
