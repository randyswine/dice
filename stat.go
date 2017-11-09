package main

import (
	"fmt"
)

// This function writing values in log file.
func logResult(values []int) error {
	fmt.Printf("Log %v\n", values)
	return nil
}

// This function calculated statistics of roll a dice by log file data.
func stat() error {
	fmt.Println("Statistics of roll a dice...\n")
	return nil
}
