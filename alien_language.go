package main

import (
	"os"
	"fmt"
	"bufio"
	"log"
	"strings"
	"strconv"
	"regexp"
)

func GrabLines(args string) []string {
	var lines []string
	file, err := os.Open(args)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func findVariables(firstline string) ([]int) {
	var returnints []int
	stringslice := strings.Split(firstline, " ")
	for _, newint := range stringslice {
		stringint, err := strconv.Atoi(newint)
		if err != nil {
			log.Fatal(err)
		}
		returnints = append(returnints, stringint)
	}
	return returnints
}

func returnCount(dictionary []string, testwords []string) {
	var replacer = strings.NewReplacer("(", "[", ")", "]")
	var casenum int
	var count int
	for i, testword := range testwords {
		count = 0
		replacedstring := replacer.Replace(testword)
		expr, err := regexp.Compile(replacedstring)
		if err != nil {
			log.Fatal(err)
		}
		for _, w := range dictionary {
			if err != nil {
				log.Fatal(err)
			}
			if expr.MatchString(w) {
				count += 1
			}
			casenum = i + 1
		}
		fmt.Println("Case #",casenum,":", count)
	}
}

func main() {
	argsWithoutProgram := os.Args[1]; if argsWithoutProgram == ""{
		fmt.Println("Need a single filename argument to run this program")
	}
	lines := GrabLines(argsWithoutProgram)
	firstline := lines[0]
	vars := findVariables(firstline)
	dictionary := lines[1:(vars[1] + 1)]
	testwords := lines[(vars[1] + 1):(vars[1] + 1 + vars[2])]
	returnCount(dictionary, testwords)
}
