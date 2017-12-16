package main

import "fmt"

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
