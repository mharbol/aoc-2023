# Advent of Code - 2023
Doing the [Advent of Code](https://adventofcode.com/) for [2023](https://adventofcode.com/2023/) in [Go\[lang\]](https://go.dev/).
Probably none of these solutions are optimal... FYI.

Even though this ended a little early for me, please check out the [retrospective](RETROSPECTIVE.md) as I try to complete it.

I went a little overboard with the interfaces and methods; I'm coming from the Java world and can't shake it just yet.

This project only uses the Go standard library so I am forced to use the language as (I guess) the designers intended.
As such, I would love to use a [`Set`](https://pkg.go.dev/github.com/deckarep/golang-set)
type... but they didn't include it in the stdlib so I can't or have to make my own if I really need it.

## Problems and Rationale

### [Day 01](solution/days/day01/day01.go)
Nothing too fancy.
For part one, I just looped forward and back to find the first and last digits.
For part two, I decided not to get fancy with front and back sliders and just made an array with all the "numbers" found in the string.

### [Day 02](solution/days/day02/day02.go)
A little hacky just copying logic from part one to part two... but it works.
Systematically broke down the lines using `strings.Split()` until I got a count and color.
From there just did some easy sums and products to get the desired results.

### [Day 03](solution/days/day03/day03.go)
This was a fun one.
For part one, looped through the rows, parse the number, and look for an adjacent symbol on the border.
Part two was a little more hacky.
Found all the numbers the same as before but did a perimeter search just for `*` characters.
If I encountered one, I would use the coordinates as the key to a map whose values were lists of the adjacent numbers.
Once looping was complete, sum the products of the lists with exactly two items.

### [Day 04](solution/days/day04/day04.go)
Today felt really good.
Got comfortable with some of Go's packages like `regexp` and `strings`.
The have some nifty functions to use... just have to remember there are more functions than methods.
For part one I was able to pull out the card number, separate the left and right sides, and get lists of number (as strings) very easily.
From there the winning number strings were the keys to a `map` with their count (initially 0).
Easy enough to sum them up and determine the score.
Part two I had some time to think of throughout the day.
Count the wins same as before but then add to that number of cards down the games.
I made a slice with all the counts of cards to easily add to and when it was time to add winning cards, I would just add the currently
played card's count. Not sure how graceful that is but it works.

### [Day 05](solution/days/day05/day05.go)
Went a little overboard on the OOP stuff for this one; I got frustrated with it and had to incrementally solve it over the course of a week.
Made an over-engineered [`Span`](util/span/span.go) struct to find the intersections and combine the mapped spans at every iteration.
The first time I solved this, I had to [redline my CPU](https://github.com/mharbol/aoc-2023/blob/8f8400ef44915a283d7313601b4968b1a701ad67/solution/days/day05.go#L30)
for 20 minutes or so to make it work.
After the new solution, it averaged less than 250ms.

### [Day 06](solution/days/day06/day06.go)
For part one I was able to just loop through all the possible races since the numbers were low enough.
For part two I decided to use the quadratic formula.
Given a time $t$ and record distance $d$, I am looking for the roots of $(t-d)x-d=0$.
Using the quadratic formula, the roots are $$r={t \pm \sqrt{t^2+4d} \over 2}$$
The left side if ceilinged and the right side is floored to get the wins from fractions.

### [Day 07](solution/days/day07/day07.go)
Not too bad of a day.
Classifying the hands in part two was a little bit of a spaghetti mashup of `if` statements, but it all was easy to make and follow.
Got comfortable with the `sort.Slice()` function which made ranking hands really easy once the comparison function was working.

### [Day 08](solution/days/day08/day08.go)
This was a fun and math-y one.
For part one it was nothing extreme, just some digraph traversal.
Part two was too big to do with a simple traversal.
By happenstance, I found that the first `XXZ` cycle came right away.
All I had to do was find the cycles, normalize them (with the `gcd`), and compose them together along with the gcd to get the LCM.
I'm happy I found it but I kinda want to go back and make the full implementation more sound and not tied to how
the problem was laid out.

### [Day 09](solution/days/day09/day09.go)
Today looked rough but wasn't too bad.
Part one was made a lot easier realizing the solution was just the sum of the end values.
Part two, just had to work bottom to top up the front of the diff slices and keep track of the previous difference.

### [Day 10](solution/days/day10/day10.go)
There probably was an easier way to traverse the pipe system for part one, but I made it work with a huge `switch`.
Did part two by determining that the area inside could be determined using the shoelace area theorem (pretty sure that's what it's called).
Used the vertices in traversed order to calculate the area.
I subtracted the half length of pipe from part one (since they contribute to the area too) and got the area of tiles inside the polygon (pipe loop).

### [Day 11](solution/days/day11/day11.go)
Lot of fun with this one.
For part one it was easy enough to just add extra rows and columns for the blank spots.
For part two, had to get a bit more creative and used my hacky set to see all the expansion points where I would have to
add distance on the backend.

### [Day 12](solution/days/day12/day12.go)
I did all sorts of silly things on this one before finally getting a good answer.
The secret was caching my answer for the recursion which greatly sped up my recursive algorithm.
I did have to take **extensive** advice from
[this](https://www.reddit.com/r/adventofcode/comments/18hbbxe/2023_day_12python_stepbystep_tutorial_with_bonus/)
Reddit post... but that's what it is there for.
At least I got reps making my own cache.

### [Day 13](solution/days/day13/day13.go)
This was a really fun one to work on.
Being clever in part one paid off for part two.
I represented the lines of `#` and `.` as the bits `1` and `0` respectively.
For part one, this made the comparison a lot quicker and easy since they were just unsigned ints.
This paid off in part two where all I needed to do was let the comparison go through if the lines were equal or `xor`ed to a power of two (one smudged bit).
From there, I knew it was the line of reflection if there was exactly one power of two for that comparison.

### [Day 14](solution/days/day14/day14.go)
Making the stones move around the `platform` reminded me of the Jump Game problems in LeetCode.
Part one was easy enough to move all the rocks off of a given tilt and then sum the weights (I knew all four directions was coming).
For part two I knew the platform would eventually come into a cyclic state.
I used the step it reached that state and the step it started to determine the period.
The remaining steps to go mod the period is the number of steps to go from the current state to the end state without completing one billion iterations.

### [Day 15](solution/days/day15/day15.go)
This was another really fun one; the real challenge was wrapping my head around the box insertion rules after a first read.
Reindeer hash was a good time but had to get creative with my `map`s for the second bit.
For part two, an object like Java's [`LinkedHashMap<>`](https://docs.oracle.com/javase/8/docs/api/java/util/LinkedHashMap.html)
would have made things worlds easier to make boxes which keep track of lenses in order with their labels as keys and focal lengths as values.
Alas, the Go standard library does not have such a structure \[to my knowledge\] so I had to hack it with my own `lensBox`.
Worked perfectly on the first try.

### [Day 16](solution/days/day16/day16.go)
Had a lot of fun on this one.
A few minor confusions on state made it a little challenging but once I got past them it was smooth sailing.
Kept a map to track if a similar beam had passed over a tile before and avoid needless processing and loops.
For part two I spent a longer time than I should have debugging just to realize I needed four distinct copies of the mirror matrix.

### [Day 17](solution/days/day17/day17.go)
To solve this I made a breadth first search which would take into account which positions were already visited but also at what step and direction.
This proved very effective in part one but a little slow in part two (takes about 6-7 seconds).
Will have to go back and fix this error soon.

### [Day 18](solution/days/day18/day18.go)
Today brought back the [shoelace formula](https://en.wikipedia.org/wiki/Shoelace_formula) to great success.
For part two I realized go has a builtin `strconv.ParseInt()` function which was really handy to parse in the hex direction.
Using regular `int` for part one made me repeat most of my logic for part two which required `int64` for accuracy.
Cleaned it up afterwards so I could reuse a lot of the same logic.

### [Day 19](solution/days/day19/day19.go)
Took a bit of a break before starting this one.
The first part was pretty easy and just had to make some good parsers.
Breaking up part two into a few days (weeks) saw the comeback of a lot of structs, methods, and functions to keep the progress for each day in scope.
Originally was going to add to the `Span` struct from day 05 but ended up going with my own inclusive span.
Each part represented a range of parts that might pass the check or fall through to the next job.
Progressively passed spanned parts through jobs and workflows until they were all accepted or rejected.
Definitely rethinking my idea to pass nearly everything as a reference, sometimes a small struct by value *might* end up being better than making
`copy()` methods all over the place.
We'll see if that blends into newer days.

### [Day 20](solution/days/day20/day20.go)
Another long break but finally back at it.
Took some time to get rid of my async tests (with enough days they caused collisions in the maps).
After that I got to solving the problems.
Part one was alright, just made an interface for `module` type and had them pulse through using a homemade queue.
Part two was irritating because it was one of those where the problem design made the math work.
`rx` was fed by a single conjunction whose high pulses were cyclic.
I found the "feeders" to the that one "root" conjunction and found the cycle of their high pulses.
The product of these cycles would be when all of the root's pulses were high and so it sent a low.

### [Day 21](solution/days/day21/day21.go)
Part one was really fun, had to read carefully to see that it is where the elf *ends up* and not where he can walk.
I made a `walkMap` of where he has gone from each new point.
If he walks to a vacant point, I mark it with the number of steps to get there.
Since any plot is an even or odd square as far as steps to get there, any even square within 64 steps is one where you can end up.
All that's left is counting the even stepped squares after 64 steps.
Doing the steps in the map was really easy, surprised this was it.

```go
type coord struct {
    row, col int
}

func walkToPoint(point coord, walkMap map[coord]int, steps int) {
    if _, ok := walkMap[point]; !ok {
        walkMap[point] = steps
    }
}
```

Part two will have to wait.
