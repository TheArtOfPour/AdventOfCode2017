package main

import (
	"bufio"
	"fmt"
	"strings"
)

func drawNodes(nodes [][]bool, x int, y int, dir rune) {
	print := "\n"
	for i, row := range nodes {
		for j, node := range row {
			if i == x && j == y {
				switch dir {
				case 'd':
					print += "▼"
				case 'u':
					print += "▲"
				case 'l':
					print += "◄"
				case 'r':
					print += "►"
				}
			} else if node {
				print += "▓"
			} else {
				print += "░"
			}
		}
		print += "\n"
	}
	fmt.Printf("%s\n", print)
}

func drawNodesRunic(nodes [][]rune, x int, y int, dir rune) {
	print := "\n"
	for i, row := range nodes {
		for j, node := range row {
			if i == x && j == y {
				switch dir {
				case 'd':
					print += "▼"
				case 'u':
					print += "▲"
				case 'l':
					print += "◄"
				case 'r':
					print += "►"
				}
			} else {
				print += string(node)
			}
		}
		print += "\n"
	}
	fmt.Printf("%s\n", print)
}

func advent22A(test string) int {
	infections := 0
	iterations := 10000
	var nodes [][]bool
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		s := scanner.Text()
		var row []bool
		for _, rune := range s {
			row = append(row, rune == '#')
		}
		nodes = append(nodes, row)
	}
	var x int
	x = len(nodes) / 2
	y := x
	direction := 'u'
	for i := 0; i < iterations; i++ {
		//fmt.Printf("\n%s (%d, %d)\n%v\n", string(direction), x, y, nodes)
		// prepend column then set x = 0
		if x < 0 {
			var newNodes [][]bool
			var temp []bool
			for j := 0; j < len(nodes[0]); j++ {
				temp = append(temp, false)
			}
			newNodes = append(newNodes, temp)
			for j := 0; j < len(nodes); j++ {
				newNodes = append(newNodes, nodes[j])
			}
			x = 0
			nodes = newNodes
		}
		// prepend row then set y = 0
		if y < 0 {
			var newNodes [][]bool
			for j := 0; j < len(nodes); j++ {
				temp := []bool{false}
				for _, newNode := range nodes[j] {
					temp = append(temp, newNode)
				}
				newNodes = append(newNodes, temp)
			}
			y = 0
			nodes = newNodes
		}
		// add column if oob
		if x >= len(nodes) {
			var temp []bool
			for j := 0; j < len(nodes[0]); j++ {
				temp = append(temp, false)
			}
			nodes = append(nodes, temp)
		}
		// add row if oob
		if y >= len(nodes[0]) {
			for j := 0; j < len(nodes); j++ {
				nodes[j] = append(nodes[j], false)
			}
		}
		drawNodes(nodes, x, y, direction)
		if nodes[x][y] {
			// clean the node
			nodes[x][y] = !nodes[x][y]
			// turn right and move forward
			switch direction {
			case 'd':
				direction = 'l'
				y--
			case 'u':
				direction = 'r'
				y++
			case 'l':
				direction = 'u'
				x--
			case 'r':
				direction = 'd'
				x++
			default:
				panic(fmt.Sprintf("invalid direction %s", string(direction)))
			}
		} else {
			// infect the node
			infections++
			nodes[x][y] = !nodes[x][y]
			// turn left and move forward
			switch direction {
			case 'd':
				direction = 'r'
				y++
			case 'u':
				direction = 'l'
				y--
			case 'l':
				direction = 'd'
				x++
			case 'r':
				direction = 'u'
				x--
			default:
				panic(fmt.Sprintf("invalid direction %s", string(direction)))
			}
		}
	}
	drawNodes(nodes, x, y, direction)
	return infections
}

func advent22B(test string) int {
	infections := 0
	iterations := 10000000
	var nodes [][]rune
	scanner := bufio.NewScanner(strings.NewReader(test))
	for scanner.Scan() {
		s := scanner.Text()
		var row []rune
		for _, rune := range s {
			if rune == '#' {
				row = append(row, '█')
			} else {
				row = append(row, ' ')
			}
		}
		nodes = append(nodes, row)
	}
	var x int
	x = len(nodes) / 2
	y := x
	direction := 'u'
	for i := 0; i < iterations; i++ {
		//fmt.Printf("\n%s (%d, %d)\n%v\n", string(direction), x, y, nodes)
		// prepend column then set x = 0
		if x < 0 {
			var newNodes [][]rune
			var temp []rune
			for j := 0; j < len(nodes[0]); j++ {
				temp = append(temp, ' ')
			}
			newNodes = append(newNodes, temp)
			for j := 0; j < len(nodes); j++ {
				newNodes = append(newNodes, nodes[j])
			}
			x = 0
			nodes = newNodes
		}
		// prepend row then set y = 0
		if y < 0 {
			var newNodes [][]rune
			for j := 0; j < len(nodes); j++ {
				temp := []rune{' '}
				for _, newNode := range nodes[j] {
					temp = append(temp, newNode)
				}
				newNodes = append(newNodes, temp)
			}
			y = 0
			nodes = newNodes
		}
		// add column if oob
		if x >= len(nodes) {
			var temp []rune
			for j := 0; j < len(nodes[0]); j++ {
				temp = append(temp, ' ')
			}
			nodes = append(nodes, temp)
		}
		// add row if oob
		if y >= len(nodes[0]) {
			for j := 0; j < len(nodes); j++ {
				nodes[j] = append(nodes[j], ' ')
			}
		}
		drawNodesRunic(nodes, x, y, direction)
		if nodes[x][y] == ' ' {
			// clean -> weaken
			nodes[x][y] = '▓'
			switch direction {
			case 'd':
				direction = 'r'
			case 'u':
				direction = 'l'
			case 'l':
				direction = 'd'
			case 'r':
				direction = 'u'
			}
		} else if nodes[x][y] == '▓' {
			// weakened -> infect
			nodes[x][y] = '█'
			infections++
		} else if nodes[x][y] == '█' {
			// infected -> flag
			nodes[x][y] = '░'
			switch direction {
			case 'd':
				direction = 'l'
			case 'u':
				direction = 'r'
			case 'l':
				direction = 'u'
			case 'r':
				direction = 'd'
			}
		} else if nodes[x][y] == '░' {
			// flagged -> clean
			nodes[x][y] = ' '
			switch direction {
			case 'd':
				direction = 'u'
			case 'u':
				direction = 'd'
			case 'l':
				direction = 'r'
			case 'r':
				direction = 'l'
			}
		}
		switch direction {
		case 'd':
			x++
		case 'u':
			x--
		case 'l':
			y--
		case 'r':
			y++
		}

	}
	drawNodesRunic(nodes, x, y, direction)
	return infections
}
