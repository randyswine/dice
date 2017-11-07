// This package implements a dice roll emulator.
package dice

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

var random *rand.Rand //randomizer, use in dice.Roll()

// Init package.
func init() {
	// Init randomizer.
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// Set of dice. This type implements fmt.Stringer.
type diceSet []dice

/*
	NewDiceSet() returns a diceSet. Each element determines the number of
	faces on each diсe of the set (d6, d32, d64, etc).
*/
func NewDiceSet(diceSizes ...int) diceSet {
	var ds diceSet
	for _, size := range diceSizes {
		// Create a dice.
		ds = append(ds, NewDice(size))
	}
	return ds
}

//  Roll() returns random value of roll for each dice of the set.
func (ds diceSet) Roll() []int {
	var results []int
	for _, dice := range ds {
		results = append(results, dice.Roll())
	}
	return results
}

// String() represents a set of dice as a string.
func (ds diceSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Dice set: ")
	buf.WriteByte('{')
	for i, dice := range ds {
		buf.WriteString(dice.String())
		if len(ds) > 1 && i < len(ds)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Dice.
type dice []int

/*
	NewDice() returns a dice. Size determines edges and number
	of faces of a new dice.
*/
func NewDice(size int) dice {
	var d dice
	for i := 1; i <= size; i++ {
		d = append(d, i)
	}
	return d
}

// Roll() returns random value of roll dice.
func (d dice) Roll() int { return d[random.Intn(len(d))] }

// String() represents a dice as a string.
func (d dice) String() string { return fmt.Sprintf("d%d", len(d)) }
