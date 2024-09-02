package main

import (
	"bufio"
	"fmt"
	"os"
)

func CountTreeHits(grid []string, movementDown, movementRight int) int {
	treesHit := 0
	horizontalPosition := 0
	for i := 0; i < len(grid); i += movementDown {
		row := grid[i]
		if row[horizontalPosition] == '#' {
			treesHit++
		}
		horizontalPosition += movementRight
		horizontalPosition = horizontalPosition % len(grid[0])
	}
	return treesHit
}

func GetEntries(fileName string) []string {
	entries := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error in opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			fmt.Println("error in parsing file")
		}
		entries = append(entries, line)
	}
	return entries
}

type Slope struct {
	down  int
	right int
}

func main() {
	grid := GetEntries("dayThreeInput.txt")
	slopesToCheck := []Slope{
		{down: 1, right: 1},
		{down: 1, right: 3},
		{down: 1, right: 5},
		{down: 1, right: 7},
		{down: 2, right: 1},
	}
	treesHit := []int{}
	for _, slope := range slopesToCheck {
		count := CountTreeHits(grid, slope.down, slope.right)
		treesHit = append(treesHit, count)
	}
	finalCount := 1
	for _, count := range treesHit {
		finalCount *= count
	}
	fmt.Println(finalCount)
}
