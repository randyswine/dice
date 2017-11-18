package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// This function define the flag of number of dice and returns this number.
func numberOfDiceFlag(name string, value int, usage string) *int {
	f := numberFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.n
}

// This type implements flag.Value interface and define number of dice.
type numberFlag struct {
	n int // Number of dice
}

// Setter of numberFlag value. Implements flag.Value interface.
func (f *numberFlag) Set(s string) error {
	_, err := fmt.Sscanf(s, "%d", &f.n)
	if err != nil {
		return fmt.Errorf("error parse param -n=%s: %s", s, err.Error())
	}
	return nil
}

// *numberFlag.String() represents flag at string.
func (f *numberFlag) String() string {
	return "foobar"
}

// This function define the flag of logging of roll a dice and returns this value.
func logValuesFlag(name string, values []int, usage string) *[]int {
	f := logFlag{values}
	flag.CommandLine.Var(&f, name, usage)
	return &f.logValues
}

// This type implements flag.Value interface and define values for logging.
type logFlag struct {
	logValues []int
}

// Setter of logFlag value. Implements flag.Value interface.
func (f *logFlag) Set(s string) error {
	values := strings.Split(s, ",")
	for _, val := range values {
		val = strings.TrimSpace(val)
		ival, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("error of convert %s to int: %s", val, err.Error())
		}
		if ival <= 0 || ival > 6 {
			return fmt.Errorf("cannot log %d at result of roll a d6", ival)
		}
		f.logValues = append(f.logValues, ival)
	}
	return nil
}

// *logFlag.String() represents flag at string.
func (f *logFlag) String() string {
	return "fizzbazz"
}

// This function define the flag of statistics param.
func statValFlag(name string, value int, usage string) *int {
	f := numberFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.n
}

// This flag defined statistics params. Value n defined number of dice.
type statFlag struct {
	n int
}

// Setter of statFlag value. Implements flag.Value interface.
func (f *statFlag) Set(s string) error {
	_, err := fmt.Sscanf(s, "%d", &f.n)
	if err != nil {
		return fmt.Errorf("error parse param -s=%s: %s", s, err.Error())
	}
	return nil
}

// *statFlag.String() represents flag at string.
func (f *statFlag) String() string {
	return "foobar"
}
