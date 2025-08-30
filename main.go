package main

import "fmt"

func getHardCodedFrequency(stringNum int, fretNum int) float64 {
	n := GuitarNote{stringNum, fretNum}
	return frequencies[n]
}

func getCalculatedFrequency(stringNum int, fretNum int) float64 {
	n := GuitarNote{stringNum, fretNum}
	return n.Frequency()
}

// sequenceOfInt generates a list of integer starting from start and
// ending with end (included)
func sequenceOfInt(start int, end int) []int {
	size := end - start + 1
	s := make([]int, int(size))
	for i := range size {
		s[i] = start + i
	}
	return s
}

func TEST01_benchmark() {
	stringNumbers := sequenceOfInt(1, 6)
	fretNumbers := sequenceOfInt(0, 19)

	accuracy := 1e-2

	msgpattern := "(%d, %2d): cf = %6.2f Hz  hf = %6.2f Hz  - %s\n"
	for _, stringNum := range stringNumbers {
		for _, fretNum := range fretNumbers {
			cf := getCalculatedFrequency(stringNum, fretNum)
			hf := getHardCodedFrequency(stringNum, fretNum)
			var msg string
			if !almostEqual(cf, hf, accuracy) {
				msg = "DIFF"
			} else {
				msg = "OK"
			}
			fmt.Printf(msgpattern, stringNum, fretNum, cf, hf, msg)
		}
	}
}

func main() {
	TEST01_benchmark()
}
