package main

import (
	"fmt"
)

func isValidSudoku(grid [3][3]int) bool {
	seen := make(map[int]bool)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := grid[i][j]
			if num < 1 || num > 9 || seen[num] {
				return false
			}
			seen[num] = true
		}
	}
	return true
}

func main() {
	grid := [3][3]int{
		{2, 7, 6},
		{9, 5, 1},
		{4, 3, 8},
	}

	if isValidSudoku(grid) {
		fmt.Println("Valid")
	} else {
		fmt.Println("Not Valid")
	}
}
