# celsenheit

`celsenheit` is both a simple tool for converting temperature values between
Celsius and Fahrenheit, and an interactive app for practicing mental
temperature conversion.

## Installation

- Clone or download this repository.
- Install the latest version of Go if not present: https://go.dev/doc/install.
- Run `go test .` to run tests.
- Run `go install` to install celsenheit to the default Go install directory.
- Alternatively, run `go build` to compile the `celsenheit` binary and then
  move it to `~/bin/` or similar.


## Conversion mode

Usage - three command-line arguments required:
```
$ celsenheit degree_value convert_from convert_to
```
e.g.:
```
$ celsenheit 20.0 C F
20°C is equivalent to 68°F.
```

The degree scale input is normalized to its capitalized first letter, so
`celsenheit 20.0 c f` and `celsenheit 20.0 celsius fahrenheit` are also valid.
Temperature values can be written with or without a decimal point but must be
numerical.

## Guess mode

Issuing the command `celsenheit` without any arguments will put you into guess
mode, as shown here:

```
$ celsenheit
Celsenheit guess mode: practice mental conversion of temperature values
=======================================================================

Enter 'Q' to exit.

Convert 17.9°C to °F: 
```

You are given randomly chosen values to convert from °C to °F or vice-versa,
and issued feedback based on how close your guess was:

```
Convert 17.9°C to °F: 65
Very close! You were off by 0.4232°C: 17.9°C is equivalent to 64.2°F.
```

Guess mode repeats infinitely until the user quits by typing `Q` (or `q`,
`quit` etc. - anything beginning with q.) 
