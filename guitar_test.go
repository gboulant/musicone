package main

import "testing"

func TestGuitar_Frequency(t *testing.T) {
	tests := []struct {
		name  string
		pluck GuitarNote
		freq  float64
	}{
		{"Mi3", GuitarNote{Mi3, 0}, Note{Octave: 3, Index: Label2Index("Mi")}.Frequency()},
		{"Si2", GuitarNote{Si2, 0}, Note{Octave: 2, Index: Label2Index("Si")}.Frequency()},
		{"Sol2", GuitarNote{Sol2, 0}, Note{Octave: 2, Index: Label2Index("Sol")}.Frequency()},
		{"Ré2", GuitarNote{Re2, 0}, Note{Octave: 2, Index: Label2Index("Re")}.Frequency()},
		{"La1", GuitarNote{La1, 0}, Note{Octave: 1, Index: Label2Index("La")}.Frequency()},
		{"Mi1", GuitarNote{Mi1, 0}, Note{Octave: 1, Index: Label2Index("Mi")}.Frequency()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := tt.pluck
			if got := n.Frequency(); got != tt.freq {
				t.Errorf("Note.Frequency() = %v, want %v", got, tt.freq)
			}
		})
	}
}
