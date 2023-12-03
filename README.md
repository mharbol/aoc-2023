# Advent of Code - 2023
Doing the [Advent of Code](https://adventofcode.com/) for [2023](https://adventofcode.com/2023/) in [Go\[lang\]](https://go.dev/).
Probably none of these solutions are optimal... FYI.

I went a little overboard with the interfaces and methods; I'm coming from the Java world and can't shake it just yet.

This project only uses the Go standard library so I am forced to use the language as (I guess) the designers intended.
As such, I would love to use a [`Set`](https://pkg.go.dev/github.com/deckarep/golang-set)
type... but they didn't include it in the stdlib so I can't or have to make my own if I really need it.

## Problems and Rationale

### [Day 01](solution/days/day01.go)
Nothing too fancy.
For part one, I just looped forward and back to find the first and last digits.
For part two, I decided not to get fancy with front and back sliders and just made an array with all the "numbers" found in the string.

### [Day 02](solution/days/day02.go)
A little hacky just copying logic from part one to part two... but it works.
Systematically broke down the lines using `strings.Split()` until I got a count and color.
From there just did some easy sums and products to get the desired results.

### [Day 03](solution/days/day03.go)
This was a fun one.
For part one, looped through the rows, parse the number, and look for an adjacent symbol on the border.
Part two was a little more hacky.
Found all the numbers the same as before but did a perimeter search just for `*` characters.
If I encountered one, I would use the coordinates as the key to a map whose values were lists of the adjacent numbers.
Once looping was complete, sum the products of the lists with exactly two items.
