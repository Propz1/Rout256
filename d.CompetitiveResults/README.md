D. Competition results (20 points)

Constraints

|  Constraint        |             Value             |
| ------------------ |:-----------------------------:|
|  Time limit        | 3 seconds                     |
|  Memory limitation | 512Mb                         | 
|  Input             | standard input                |   
|  Output            | standard output               |


N athletes took part in the running competition: the i-th of them ran the distance in ti seconds. The jury wants to assign places to participants according to the following rules:

* the seats are numbered from 1 onwards (the best place is the first);
* if two athletes have the same results or differ by one second, then they share a place (in this case, we assume that they share the best of the divided places);
* participants share a place only as a result of applying the previous rule (possibly several times);
* if k participants share a place p, then the places of the following participants are numbered starting from k+p.

Consider the following examples to understand the principle of assigning places:

* let's say n=4 and t=[20,10,20,30], then the places have the form [2,1,2,4] (the second athlete ran first — he has the first place, the first and third shared the second place, the fourth took the last fourth place);
* let's say n=3 and t=[5,7,6], then the places have the form [1,1,1] (since t1=5 and t3=6 differ by 1, then the first and third athletes should take the same place, similarly with the second and third athletes, therefore, all three share the first place);
* let's say n=5 and t=[6,3,4,3,1], then the places have the form [5,2,2,2,1];
* let's say n=5 and t=[200,10,100,11,200], then the places have the form [4,1,3,1,4].

By the given values n and t1,t2,...,tn, output a sequence of places occupied by athletes.

Incomplete solutions to this problem (for example, insufficiently effective) can be assessed with a partial score.

Input data
The first line contains an integer t (1≤t≤1000) — the number of input data sets in the test.

The sets of input data in the test are independent. They do not affect each other in any way.

The first line of each input data set contains an integer n (1≤n≤200000) — the number of athletes.

The second line of the set contains a sequence of integers t1,t2,...,tn (1≤ti≤1000000000), where ti is the time in seconds for which the i—th athlete ran the distance.

The sum of the values of n for all sets of test input data does not exceed 200000.

Output data
For each set of input data, output n positive numbers r1,r2,...,rn, where ri is the place of the i—th athlete.

Example

|         Input data         |  Output data           |
| -------------------------- | ---------------------- |
| 6                          | 2 1 2 4                |
| 4                          | 1 1 1                  |
| 20 10 20 30                | 5 2 2 2 1              |
| 3                          | 4 1 3 1 4              |
| 5 7 6                      | 1                      |
| 5                          | 9 4 9 1 4 7 1 4 7 1 11 |
| 6 3 4 3 1                  |                        |
| 5                          |                        |
| 200 10 100 11 200          |                        |
| 1                          |                        |
| 1000000000                 |                        |
| 11                         |                        |
| 13 8 12 1 7 10 1 8 10 2 17 |                        |

