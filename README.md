# MusicOne - Calculating the equal temperament scale frequencies

This project is a sandbox to illustrates how to calculate the
frequencies of the equal temperament music scale. In this example, a
note is characterized by 2 integer numbers: the octave index and the
order index of the note in this octave, considering 12 equally distant
intervals (the definition of the equal temperament scale). We show also
how to determine the note of a guitar pluck (defined by the string index
and the fret index), and then calculate the frequency of this guitar
note.

**contact**: [Guillaume Boulant](mailto:gboulant@gmail.com?subject=musicone)

The project was initially motivated to test an adaptation of the
[timiskhakov music](https://github.com/timiskhakov/music) project from
Timor Iskhakov. Its project shows how to emulate the sound of a guitar
and play songs from a tablature. Timiskhakov uses hard coded values for
the frequencies (using a map that associate each guitar note to its
associated frequency). Indeed, for performance reason, it is surelly a
good choice to get the frequency value from a map, instead of
dynamically calculate the value.

In our example, we show how to calculate the frequency by adding or
substracting intervals to/from a reference frequency (the La3), the only
one to define (to 440 Hz for example). The main program scans all the
guitar notes (by looping on the string index and then the fret index)
and for each note 1/ read the frequency value in the hard coded table
(copied from the Timor project) and 2/ calculate the frequency value
using our algorithm, for finally 3/ compare the values and check that we
obtain the same values with a precision of 1e-2 Hz.

For testing:

```shell
go mod tidy
go build
./musicone
```

Note that in this project, we do not play the note sound, we only
show how to calculate the frequencies. For playing music, go to the
[timiskhakov music](https://github.com/timiskhakov/music) project, or to
my own project XXX (strongly inspired from the Timor project).
