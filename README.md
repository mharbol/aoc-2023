# Advent of Code - 2023
Doing the [Advent of Code](https://adventofcode.com/) for [2023](https://adventofcode.com/2023/) in [Go\[lang\]](https://go.dev/).
Probably none of these solutions are optimal... FYI.

Even though this ended a little early for me, please check out the [retrospective](#Retrospective) as I try to complete it.

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

### [Day 04](solution/days/day04.go)
Today felt really good.
Got comforable with some of Go's packages like `regexp` and `strings`.
The have some nifty functions to use... just have to remember there are more functions than methods.
For part one I was able to pull out the card number, separate the left and right sides, and get lists of number (as strings) very easily.
From there the winning number strings were the keys to a `map` with their count (initially 0).
Easy enough to sum them up and determine the score.
Part two I had some time to think of throughout the day.
Count the wins same as before but then add to that number of cards down the games.
I made a slice with all the counts of cards to easily add to and when it was time to add winning cards, I would just add the currently
played card's count. Not sure how graceful that is but it works.

### [Day 05](solution/days/day05.go)
Went a little overboard on the OOP stuff for this one; I got frustrated with it and had to incrementally solve it over the course of a week.
Made an over-engineered [`Span`](util/span/span.go) struct to find the intersections and combine the mapped spans at every iteration.
The first time I solved this, I had to [redline my CPU](https://github.com/mharbol/aoc-2023/blob/8f8400ef44915a283d7313601b4968b1a701ad67/solution/days/day05.go#L30)
for 20 minutes or so to make it work.
After the new solution, it averaged less than 250ms.

### [Day 06](solution/days/day06.go)
For part one I was able to just loop through all the possible races since the numbers were low enough.
For part two I decided to use the quadratic formula.
Given a time $t$ and record distance $d$, I am looking for the roots of $(t-d)x-d=0$.
Using the quadratic formula, the roots are $$r={t \pm \sqrt{t^2+4d} \over 2}$$
The left side if ceilinged and the right side is floored to get the wins from fractions.

### [Day 07](solution/days/day07.go)
Not too bad of a day.
Classifying the hands in part two was a little bit of a spaghetti mashup of `if` statements, but it all was easy to make and follow.
Got comfortable with the `sort.Slice()` function which made ranking hands really easy once the comparison function was working.

### [Day 08](solution/days/day08.go)
This was a fun and math-y one.
For part one it was nothing extreme, just some digraph traversal.
Part two was too big to do with a simple traversal.
By happenstance, I found that the first `XXZ` cycle came right away.
All I had to do was find the cycles, normalize them (with the `gcd`), and compose them together along with the gcd to get the LCM.
I'm happy I found it but I kinda want to go back and make the full implementation more sound and not tied to how
the problem was laid out.

### [Day 09](solution/days/day09.go)
Today looked rough but wasn't too bad.
Part one was made a lot easier realizing the solution was just the sum of the end values.
Part two, just had to work bottom to top up the front of the diff slices and keep track of the previous difference.

### [Day 10](solution/days/day10.go)
There probably was an easier way to traverse the pipe system for part one, but I made it work with a huge `switch`.
Did part two by determining that the area inside could be determined using the shoelace area theorem (pretty sure that's what it's called).
Used the vertices in traversed order to calculate the area.
I subtracted the half length of pipe from part one (since they contribute to the area too) and got the area of tiles inside the polygon (pipe loop).

### [Day 11](solution/days/day11.go)
Lot of fun with this one.
For part one it was easy enough to just add extra rows and columns for the blank spots.
For part two, had to get a bit more creative and used my hacky set to see all the expansion points where I would have to
add distance on the backend.

### [Day 12](solution/days/day12.go)
I did all sorts of silly things on this one before finally getting a good answer.
The secret was caching my answer for the recursion which greatly sped up my recursive algorithm.
I did have to take **extensive** advice from
[this](https://www.reddit.com/r/adventofcode/comments/18hbbxe/2023_day_12python_stepbystep_tutorial_with_bonus/)
Reddit post... but that's what it is there for.
At least I got reps making my own cache.

### [Day 13](solution/days/day13.go)
This was a really fun one to work on.
Being clever in part one paid off for part two.
I represented the lines of `#` and `.` as the bits `1` and `0` respectively.
For part one, this made the comparison a lot quicker and easy since they were just unsigned ints.
This paid off in part two where all I needed to do was let the comparison go through if the lines were equal or `xor`ed to a power of two (one smudged bit).
From there, I knew it was the line of reflection if there was exatly one power of two for that comparison.

### [Day 14](solution/days/day14.go)
Making the stones move around the `platform` reminded me of the Jump Game problems in LeetCode.
Part one was easy enough to move all the rocks off of a given tilt and then sum the weights (I knew all four directions was coming).
For part two I knew the platform would eventully come into a cyclic state.
I used the step it reached that state and the step it started to determine the period.
The remaining steps to go mod the period is the number of steps to go from the current state to the end state without completing one billion iterations.

### [Day 15](solution/days/day15.go)
This was another really fun one; the real challenge was wrapping my head around the box insertion rules after a first read.
Reindeer hash was a good time but had to get creative with my `map`s for the second bit.
For part two, an object like Java's [`LinkedHashMap<>`](https://docs.oracle.com/javase/8/docs/api/java/util/LinkedHashMap.html)
would have made things worlds easier to make boxes which keep track of lenses in order with their labels as keys and focal lengths as values.
Alas, the Go standard library does not have such a structure \[to my knowledge\] so I had to hack it with my own `lensBox`.
Worked perfectly on the first try.

## Retrospective
I like the multithreading, it is very easy to kick off if you are smart about it.
Errors as return types rather than Exceptions coming up from somewhere are a plus.
Visibility by function/method name capitalization is terrible.

Done diferently, I would have made each day its own `interface` to avoid weird naming as the days progressed
(maybe a refactor should be in the works).
