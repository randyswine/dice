package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const logFile string = "roll_stat.log"

// This function writing values of roll a dice in log file.
func logResult(values []int) error {
	f, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString(toString(values))
	if err != nil {
		return fmt.Errorf("cannot logging %v: %s", values, err.Error())
	}
	return nil
}

// This function calculated statistics of roll a dice by log file data.
func stat(numberOfDice int) error {
	var (
	//lines []string
	)
	sumCounts := make(map[int]int)
	content, err := ioutil.ReadFile(logFile)
	if err != nil {
		return fmt.Errorf("error of open log: %s", err.Error())
	}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		numberOfDiceInLine := strings.Count(line, "+")
		numberOfDiceInLine += 1
		if numberOfDiceInLine == numberOfDice {
			sides := strings.Split(line, "=")
			sum := sides[len(sides)-1]
			i, err := strconv.Atoi(sum)
			if err != nil {
				return fmt.Errorf("cannot convert %s to int: %s", sum, err.Error())
			}
			sumCounts[i]++
		}
	}
	fmt.Println(sumCounts)
	return nil
}

// This function sums the elements of the []values and returns a string of the form "2+6=8\n".
func toString(values []int) string {
	var (
		terms []string
		sum   int
	)
	for _, val := range values {
		sum += val
		terms = append(terms, strconv.Itoa(val))
	}
	s := strings.Join(terms, "+")
	s = fmt.Sprintf("%s=%d\n", s, sum)
	return s
}
