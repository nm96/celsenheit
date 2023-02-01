# celsenheit

`celsenheit` is both a simple tool for converting temperature values between
Celsius and Fahrenheit, and an interactive app for practicing mental
temperature conversion.

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
