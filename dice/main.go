// This app emulated roll a dice.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/randyswine/dice"
)

func main() {
	rf := &rollFlag{}
	flag.CommandLine.Var(rf, "roll", "Example: --roll 2d6")
	flag.Parse()
	var diceSizes []int
	for i := 0; i < rf.diceN; i++ {
		diceSizes = append(diceSizes, rf.diceSize)
	}
	ds := dice.NewDiceSet(diceSizes...)
	fmt.Println(ds)
	os.Exit(0)
}

// This type implements flag.Value interface and define set of dice.
type rollFlag struct {
	diceN, diceSize int //Count of dice
}

// Setter of flag value. Implements flag.Value interface.
func (f *rollFlag) Set(s string) error {
	_, err := fmt.Sscanf(s, "%dd%d", &f.diceN, &f.diceSize)
	if err != nil {
		err = errors.New(fmt.Sprintf("error parse param --roll=%s: %s", s, err.Error()))
		return err
	}
	return nil
}

// *rollFlag.String() represents flag at string.
func (f *rollFlag) String() string {
	return "what"
}
