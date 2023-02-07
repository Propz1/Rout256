B. Sea Battle (10 points)

Constraints

|  Constraint        |             Value             |
| ------------------ |:-----------------------------:|
|  Time limit        | 2 seconds                     |
|  Memory limitation | 512Mb                         | 
|  Input             | standard input                |   
|  Output            | standard output               |

You are involved in the development of a subsystem for checking the field for the game "Sea Battle". You need to write a check of the correctness of the number of ships on the field, taking into account their size. Recall that the field should be:

* four single-deck ships,
* three double-decked ships,
* two three-decked ships,
* one four-decked ship.

You are given 10 integers from 1 to 4. Check that the specified sizes meet the requirements above.

Input data
The first line contains an integer t (1≤t≤1000) — the number of input data sets in the test.

The sets of input data in the test are independent. They do not affect each other in any way.

Each set of input data consists of one line, which contains 10 integers a1,a2,...,a10 (1≤ai≤4) — the sizes of ships on the field in any order.

Note that it is already guaranteed that there are exactly 10 ships on the field and their sizes are from 1 to 4, inclusive. You need to check that the number of ships of each type corresponds to the rules of the game.


Output data
For each set of input data in a separate line, output:

* YES, if the specified sizes of ships on the field correspond to the rules of the game;
* NO otherwise. 

You can output YES and NO in any case (for example, the strings yEs, yes, Yes and YES will be recognized as a positive answer).

Example

|         Input data         |  Output data  |
| -------------------------- | ------------- |
| 5                          | YES           |
| 2 1 3 1 2 3 1 1 4 2        | NO            |
| 1 1 1 2 2 2 3 3 3 4        | YES           |
| 1 1 1 1 2 2 2 3 3 4        | YES           |
| 4 3 3 2 2 2 1 1 1 1        | NO            |
| 4 4 4 4 4 4 4 4 4 4        |               |
