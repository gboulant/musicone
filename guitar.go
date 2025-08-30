package main

import "github.com/timiskhakov/music/guitar"

// The guitar strings are numbered from 1 (the highest sound, Mi3) to 6
// (the lower sound Mi1).
const (
	Mi3 int = iota + 1
	Si2
	Sol2
	Re2
	La1
	Mi1
)

var openStringNotes map[int]Note = map[int]Note{
	Mi3:  {Octave: 3, Index: Label2Index("Mi")},
	Si2:  {Octave: 2, Index: Label2Index("Si")},
	Sol2: {Octave: 2, Index: Label2Index("Sol")},
	Re2:  {Octave: 2, Index: Label2Index("Re")},
	La1:  {Octave: 1, Index: Label2Index("La")},
	Mi1:  {Octave: 1, Index: Label2Index("Mi")},
}

// noteOfOpenString returns the music.Note played when we pluck the
// specified open string (without pressing any fret).
func noteOfOpenString(stringNum int) Note {
	note, ok := openStringNotes[stringNum]
	if !ok {
		LogError("err: (noteOfOpenString) the string number %d is not defined\n", stringNum)
	}
	return note
}

// GuitarNote defines a musical note from a guitar point of view, i.e.
// by specifying the string number to pluck and the fret number to
// press. Concerning the string number, the convention is to designate
// the bass Mi string (string at the top of the guitar neck) as the
// string number 6, and the high Mi string (at the bottom of the neck)
// as the string number 1. To avoid error, you should use the constant
// definitions above (Mi3=1, Si2=2, ..., Mi1=6). A fret number of 0
// means "don't press any fret (play the open string)".
type GuitarNote guitar.Note

// Frequency return the frequency of the sound corresponding to this
// guitar note, defined by its string number and fret number. The
// frequency is computed by adding/substracting intervals to/from a
// reference note (the La3).
func (p GuitarNote) Frequency() float64 {
	// 1. Retrieve the note played by the the open string
	note := noteOfOpenString(p.String)
	// 2. Add the interval corresponding to the fret being pressed (nb of half-tons)
	note.Add(Interval(p.Fret))
	// 3. Compute the frequency of the note
	return note.Frequency()
}
