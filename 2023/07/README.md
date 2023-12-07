https://adventofcode.com/2023/day/7

* main parts for a:
  * struct representing a hand, including long if-else to determine the type. This was done by storing seen cards into a map, and then just counting number of keys (seen ranks), and the values if necessary (to differentiate 3+1+1 from 2+2+1)
  * struct holding the hands, by using array of size 7, each index being the next higher hand type. Each array entry is a linked list of hands, in increasing stength
  * To calculate score, simply loop through all linked lists in order
* for b, simply change value of 'J' from 11 to 1 (lower than '2'), and then add a bunch of extra hard-coded logic for determining hand type based on number of jokers
* SPENT A LONG TIME DEBUGGING BECAUSE I WROTE 1 (one) INSTEAD OF i in an array lookup
* Forgot to delete 'J' from the map and that caused a slight bug (should have probably special cased it out from the start instead of still adding it to the map).
