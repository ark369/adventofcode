[Day 8](https://adventofcode.com/2020/day/8)

Part b solution is inefficient since the full list of instructions is copied and copied again on every iteration over path. Minor optimizations: skip nop/jmp 1, immediately return on jmp 0.
