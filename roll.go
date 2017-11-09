package main

import (
	"math/rand"
	"time"
)

var (
	random *rand.Rand                            // Randomizer, use in roll().
	d6     [6]int     = [6]int{1, 2, 3, 4, 5, 6} // Model of d6.
)

// Init package.
func init() {
	// Init randomizer.
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// This function emulated roll a d6.
func roll(n int) []int {
	var results []int
	for i := 0; i < n; i++ {
		results = append(results, d6[random.Intn(len(d6))])
	}
	return results
}
