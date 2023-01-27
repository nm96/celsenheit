# celsenheit

`celsenheit` is both a very simple tool for converting temperature values
between Celsius and Fahrenheit, and an interactive app for practicing mental
temperature conversion.

## Conversion mode

Usage - three command-line arguments required:
```
$ celsenheit degree_value convert_from convert_to
```
e.g.:
```
$ celsenheit 20.0 C F
Converting 20.0°C to °F:
20°C is equivalent to 68°F.
```

The degree scale input is normalized to its capitalized first letter, so
`celsenheit 20.0 c f` and `celsenheit 20.0 celsius fahrenheit` are also valid.

## Guess mode

Issuing the command `celsenheit` without any arguments will put you into guess
mode, as shown here:

```
Celsenheit guess mode: practice temperature conversions on random values!
=========================================================================

Convert -37.4°C to °F:
```

You are given randomly chosen values to convert from °C to °F or vice-versa,
and issued feedback based on how close your guess was:

```
Convert -37.4°C to °F: -35
Your guess: -35
Astonishing!
*****
-37.4°C is equivalent to -35.4°F.
```
