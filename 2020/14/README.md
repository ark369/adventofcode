[Day 14](https://adventofcode.com/2020/day/14)

All string-to-bit manipulation stuff, but not too hard.

Built an AND mask and an OR mask out of the mask to apply to values in part a.

In part b, used fmt.Sprintf("%036s", strconv.FormatUint(uint64(loc), 2)) to convert a memory location into a string representation that can be paired with the mask. Then just simply iterated over mask and built up a list of addresses to update. Silly bug: needing to initialize the addresses list with "" to seed for the first iteration!
