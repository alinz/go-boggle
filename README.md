## Go-Boggle Solver
This was an attempt in 15 min to solve the Boggle in go.

As soon as you run this code, my simple graph tries to find all possible words and call the callback function for every word it finds.
Use that function to check against your dictionary.


## Input file format
the first line contains all possible node as a label and the rest of the lines are edges.

for example

```
a b c d    //it means we have 4 nodes
a c d      //it means that a connects to c and d
c b        //it means that c connects to b
```
