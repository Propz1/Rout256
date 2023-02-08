B. Amount to be paid (10 points)

Constraints

|  Constraint        |             Value             |
| ------------------ |:-----------------------------:|
|  Time limit        | 1 seconds                     |
|  Memory limitation | 256Mb                         | 
|  Input             | standard input                |   
|  Output            | standard output               |

There is an action in the store: "buy three identical goods and pay only for two." Of course, each purchased product can participate in only one promotion. The promotion can be used multiple times.

For example, if goods of one type are purchased at a price per piece and goods of another type at a price per piece, then instead it will be necessary to pay .

Assuming that only the same products have the same prices, find the amount to be paid.

Incomplete solutions to this problem (for example, insufficiently effective) can be assessed with a partial score.

Input data

The first line contains an integer t (1≤t≤10000) — the number of input data sets.

Further sets of input data are recorded. Each begins with a line that contains n (1≤n≤200000) — the number of purchased items. The next line contains their prices p1,p2,...,pn (1≤pi≤10000). If the prices of two products are the same, then we must assume that they are the same product.

It is guaranteed that the sum of the values for all tests does not exceed 200000.

Output data
Print t integers — the amounts to be paid for each of the input data sets.

Example

|         Input data                 |  Output data  |
| ---------------------------------- | ------------- |
| 6                                  | 22            |
| 12                                 | 22            |
| 2 2 2 2 2 2 2 3 3 3 3 3            | 10000         |
| 12                                 | 12            |
| 2 3 2 3 2 2 3 2 3 2 2 3            | 40000         |
| 1                                  | 1100          |
|10000                               |               |
|9                                   |               |
|1 2 3 1 2 3 1 2 3                   |               |
|6                                   |               |
|10000 10000 10000 10000 10000 10000 |               |
|6                                   |               |
|300 100 200 300 200 300             |               |
