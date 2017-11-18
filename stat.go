package main

import (
	"fmt"
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
func stat() error {
	var err error
	return err
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
