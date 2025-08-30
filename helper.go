package main

import (
	"log"
	"math"
)

// LogError must be used to log error on the terminal. Choose if it must
// stop the process or just print and continue
var LogError = log.Fatalf

//var LogError = fmt.Printf

func almostEqual(a, b float64, accuracy float64) bool {
	return math.Abs(a-b) < accuracy
}
