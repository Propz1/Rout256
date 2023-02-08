C. Pair programming (10 points)

Constraints

|  Constraint        |             Value             |
| ------------------ |:-----------------------------:|
|  Time limit        | 1 seconds                     |
|  Memory limitation | 512Mb                         | 
|  Input             | standard input                |   
|  Output            | standard output               |

The company employs n developers, where n is an even number. The crazy manager decided to split all the developers into teams of two people.

To do this, he compiled a list of all developers and assigned each of them a number on the list (from 1 to n) and the value ai — the skill level of the i-th in the developer's list.

He makes the next command as follows:
 
1. the first developer in the team is the one who goes first in the list;
2. a pair is selected for him such that the difference between their levels is minimal (that is, the minimum value of |ai−aj|, where |x| is the modulus of the number x); if there are several such candidates, then the one who is earlier in the list is selected from them;
3. these two developers form a team and are removed from the list.

For example, if the array a is equal to [2,1,3,1,1,4], then the formation of commands will occur as follows:

1. assign the developers numbers [1,2,3,4,5,6] according to their position in the list, the first among them has number 1, his skill level a1=2, suitable (with a minimum absolute difference) are developers with numbers 2,3,4,5, the first among them is 2, so the first team is developers with numbers 1 and 2;
2. the remaining developers now have numbers [3,4,5,6], the first among them is 3, its level a3=3, there is only one developer with a minimum absolute difference (number 6), so the team is developers with numbers 3 and 6;
3. the remaining developers have numbers [4,5], the first among them is 4, its level a4 = 1, only the developer with number 5 remains, so the third team is developers with numbers 4 and 5.

Your task is to help the crazy manager simulate the process of splitting into teams. Note that the commands must be output in the order described above in the condition.

Input data

The first line contains a single integer t (1≤t≤50) — the number of input data sets.

The first line of each set contains one integer n (2≤n≤50; n is even) — the number of developers.

The second line contains n integers a1,a2,...,an (1≤ai≤100), where ai is the skill level of the i—th developer.

Output data

For each set of input data output n2 lines, the i-th line should contain a pair of numbers — the number of the first and second developer in the i-th team in the order described in the condition.

Output an empty line between the outputs for the input data sets.

Example

|    Input data   |  Output data  |
| --------------- | ------------- |
| 3               | 1 2           |
| 6               | 3 6           |
| 2 1 3 1 1 4     | 4 5           |
| 2               |               |
| 5 5 5           | 1 2           |
| 8               |               |
| 1 4 2 5 4 2 6 3 | 1 3           |
|                 | 2 5           |
|                 | 4 7           |
|                 | 6 8           |

