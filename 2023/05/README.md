https://adventofcode.com/2023/day/5

* a was straightforward
* b required a lot of sorted linked lists of ranges. Each current set of ranges would be fed into each subsequent (sorted) mapping, and the new ranges would be sorted again and used as the next round of input. The final sorted ranges would have the answer as the leftmost element of the head of the list.
