// This app emulated roll a dice.
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Parse number of dice.
	n := numberOfDiceFlag("n", -1, "Example: -n 2")
	flag.Parse()
	// Roll and print result.
	if *n > 0 {
		for _, val := range roll(*n) {
			fmt.Printf("%d\t", val)
		}
	}
	os.Exit(0)
}

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

// Setter of flag value. Implements flag.Value interface.
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
