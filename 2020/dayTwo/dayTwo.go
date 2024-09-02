package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Params struct {
	firstNum  int
	secondNum int
	letter    byte
	password  string
}

func GetParams(str string) Params {
	sections := strings.Split(str, " ")
	minMax := strings.Split(sections[0], "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])
	letter := (sections[1][0])
	password := sections[2]
	return Params{min, max, letter, password}
}

func ValidatePassword(params Params) bool {
	count := 0
	for _, letter := range params.password {
		if byte(letter) == params.letter {
			count++
		}
	}
	return count >= params.firstNum && count <= params.secondNum
}

func ValidatePasswordTwo(params Params) bool {
	count := 0
	if params.letter == (params.password[params.firstNum-1]) {
		count++
	}
	if params.letter == (params.password[params.secondNum-1]) {
		count++
	}
	return count == 1
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

func main() {
	entries := GetEntries("dayTwoInput.txt")
	count := 0
	for _, entry := range entries {
		if ValidatePasswordTwo(GetParams(entry)) == true {
			count++
		}
	}
	fmt.Println(count)
}
