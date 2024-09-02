package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func TwoSum(numList []int) int {
	for i := 0; i < len(numList)-1; i++ {
		for j := i + 1; j < len(numList); j++ {
			if numList[i]+numList[j] == 2020 {
				return numList[i] * numList[j]
			}
		}
	}
	return 0
}

func ThreeSum(numList []int) int {
	for i := 0; i < len(numList)-2; i++ {
		for j := i + 1; j < len(numList)-1; j++ {
			for k := j + 1; k < len(numList); k++ {
				if numList[i]+numList[j]+numList[k] == 2020 {
					return numList[i] * numList[j] * numList[k]
				}
			}
		}
	}
	return 0
}

func getNums(fileName string) []int {
	numList := []int{}
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Println("error in opening file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("error in parsing file")
		}
		numList = append(numList, num)
	}
	slices.Sort(numList)
	return numList
}

func main() {
	numList := getNums("dayOneInput.txt")
	value := ThreeSum(numList)
	fmt.Println(value)
}
