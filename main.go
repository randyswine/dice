// This app emulated roll a dice.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Parse number of dice.
	n := numberOfDiceFlag("n", -1, "Example: -n 2")
	logValues := logValuesFlag("l", nil, "Example: -l 2,6")
	flag.Parse()
	// Roll and print result.
	if *n > 0 {
		if *n > 10 {
			log.Fatal("Too many dice (0 < n <= 10).")
		}
		for _, val := range roll(*n) {
			fmt.Printf("%d\t", val)
		}
	}
	if logValues != nil {
		fmt.Println(logValues)
	}
	os.Exit(0)
}
