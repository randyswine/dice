package dice

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
)

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

type diceSet []dice

func NewDiceSet(diceSizes ...int) diceSet {
	var ds diceSet
	for _, size := range diceSizes {
		ds = append(ds, NewDice(size))
	}
	return ds
}

func (ds diceSet) Roll() []int {
	var results []int
	for _, dice := range ds {
		results = append(results, dice.Roll())
	}
	return results
}

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

type dice []int

func NewDice(size int) dice {
	var d dice
	for i := 1; i <= size; i++ {
		d = append(d, i)
	}
	return d
}

func (d dice) Roll() int {
	result := random.Intn(len(d) + 1)
	if result == 0 {
		result = d.Roll()
	}
	return result
}

func (d dice) String() string {
	return fmt.Sprintf("d%d", len(d))
}
