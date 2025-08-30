package main

import (
	"testing"
)

func TestNote_Interval(t *testing.T) {
	n1 := Note{3, 9} // La3
	n2 := Note{0, 0} // Do0

	res := n1.IntervalTo(n2)
	exp := Interval(-45)
	if res != exp {
		t.Errorf("interval is %d (should be %d)", res, exp)
	}

}

func TestNote_Frequency(t *testing.T) {
	type fields struct {
		Octave int
		Index  NoteIndex
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"Do0", fields{0, Label2Index("Do")}, 32.703},
		{"Ré1", fields{1, Label2Index("Re")}, 73.416},
		{"Mi2", fields{2, Label2Index("Mi")}, 164.813},
		{"Fa2", fields{2, Label2Index("Fa")}, 174.614},
		{"Sol2", fields{2, Label2Index("Sol")}, 195.997},
		{"La3", fields{3, Label2Index("La")}, 440.000},
		{"Si3", fields{3, Label2Index("Si")}, 493.883},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := Note{
				Octave: tt.fields.Octave,
				Index:  tt.fields.Index,
			}

			if got := n.Frequency(); !almostEqual(got, tt.want, 1e-3) {
				t.Errorf("Note.Frequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
