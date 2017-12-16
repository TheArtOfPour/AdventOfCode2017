package main

import (
	"fmt"
	"strconv"
)

// HexToBin binary goodness
func HexToBin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return "", err
	}

	format := fmt.Sprintf("%%0%db", len(hex)*4)
	return fmt.Sprintf(format, ui), nil
}

func advent14A(test string) int {
	total := 0
	for row := 0; row < 128; row++ {
		rowString := strconv.Itoa(row)
		toHash := test + "-" + rowString
		hash := advent10B(toHash)
		for _, c := range hash {
			binString, _ := HexToBin(string(c))
			for _, b := range binString {
				if b == '1' {
					total++
				}
			}
			// fmt.Printf("Binary %v\n", binString)
		}
	}
	return total
}

func printDisk(disk [128][128]bool) {
	fmt.Printf("\n")
	diskString := ""
	for i := range disk {
		for j := range disk[i] {
			if disk[i][j] {
				diskString += "#"
			} else {
				diskString += "."
			}
		}
		diskString += "\n"
	}
	fmt.Printf("%s\n", diskString)
}

func defrag(disk *[128][128]bool, row int, col int) {
	disk[row][col] = false
	if row > 0 && disk[row-1][col] {
		defrag(disk, row-1, col)
	}
	if row < 127 && disk[row+1][col] {
		defrag(disk, row+1, col)
	}
	if col > 0 && disk[row][col-1] {
		defrag(disk, row, col-1)
	}
	if col < 127 && disk[row][col+1] {
		defrag(disk, row, col+1)
	}
}

func advent14B(test string) int {
	var disk [128][128]bool
	total := 0
	for row := 0; row < 128; row++ {
		rowString := strconv.Itoa(row)
		toHash := test + "-" + rowString
		hash := advent10B(toHash)
		for ic, c := range hash {
			binString, _ := HexToBin(string(c))
			for ib, b := range binString {
				if b == '1' {
					disk[row][4*ic+ib] = true
				}
			}
		}
	}
	printDisk(disk)
	for i := range disk {
		for j := range disk[i] {
			if disk[i][j] {
				defrag(&disk, i, j)
				total++
			}
		}
		// cmd := exec.Command("cmd", "/c", "cls")
		// cmd.Stdout = os.Stdout
		// cmd.Run()
		printDisk(disk)
	}
	return total
}
