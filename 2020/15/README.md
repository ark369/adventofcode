[Day 15](https://adventofcode.com/2020/day/15)

Part a was straightforward.

Part b kept giving `Program exited: status 2.` on The Go Playground, despite optimizations (uint32s, keeping a lowNums array for much more frequently-accessed elements). Ended up running code in linux with `go run` and the answer came back in seconds.
