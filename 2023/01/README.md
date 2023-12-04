[https://adventofcode.com/2023/day/1](https://adventofcode.com/2023/day/1)

a was straightforward.

b had trickiness in word-number tracking:
* "oo" requires resetting the tracker for 1 to be "o" instead of "" (same for all 1-8)
* "nini" requires resetting 9 tracker to "ni" ("nin" is still valid)
