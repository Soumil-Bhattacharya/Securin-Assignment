# Securin-Assignment

<!--toc:start-->
- [Securin-Assignment](#securin-assignment)
  - [Problem Statement: The Doomed Dice Challenge](#problem-statement-the-doomed-dice-challenge)
    - [Part A](#part-a)
      - [Question 1](#question-1)
        - [Solution](#solution)
      - [Question 2](#question-2)
        - [Solution](#solution)
      - [Question 3](#question-3)
        - [Solution](#solution)
    - [Part B](#part-b)
      - [Problem Statement](#problem-statement)
      - [Solution](#solution)
<!--toc:end-->

Solutions for Securin Assignment

## Problem Statement: The Doomed Dice Challenge


### Part A

####  Question 1

Q: **How many total combinations are possible? Show the math along with the code**

##### Solution 

Number of faces in each die: $6$

Hence, number of possible combinations: $6\times 6 = 36$

Ans: $36$

Code:

The above answer can be verified using the following code (written in Go):

```go
import "fmt"
func ques1()  {
    var (
        die1 [6]int = [6]int{1,2,3,4,5,6}
        die2 [6]int = [6]int{1,2,3,4,5,6}
        total int = 0
    )

    for _, i := range die1 {
        for _, j := range die2 {
            total += 1
        }
        fmt.Println()
    }

    fmt.Printf("Total: %d possible combinations\n", total)

}
```

####  Question 2

**Calculate and display the distribution of all possible combinations that can be obtained when rolling both Die A and Die B together. Show the math along with the code!**

##### Solution

   The distribution of all possible can be represented using a 6 x 6 matrix where each element in the matrix represents the sum of the numbers on die A and die B

```
[1,1]   [1,2]   [1,3]   [1,4]   [1,5]   [1,6]
[2,1]   [2,2]   [2,3]   [2,4]   [2,5]   [2,6]
[3,1]   [3,2]   [3,3]   [3,4]   [3,5]   [3,6]
[4,1]   [4,2]   [4,3]   [4,4]   [4,5]   [4,6]
[5,1]   [5,2]   [5,3]   [5,4]   [5,5]   [5,6]
[6,1]   [6,2]   [6,3]   [6,4]   [6,5]   [6,6]
```

```go 
import "fmt"
func ques2()  {
    var (
        dieA [6]int = [6]int{ 1, 2, 3, 4, 5, 6, }
        dieB [6]int = [6]int{ 1, 2, 3, 4, 5, 6}
        freq map[int]int = make(map[int]int)
    )

    for _, i := range dieA {
        for _, j := range dieB {
            freq[i+j] = freq[i+j] + 1
            fmt.Printf("%d\t", i + j)
        }
        fmt.Println()
    }

    fmt.Println("Sum\tFrequency")
    for key := 2; key <= 12; key++ {
        fmt.Printf("%d:\t%d \n", key, freq[key])
    }

}
```

The resulting 6 x 6 matrix looks like this:

```
2       3       4       5       6       7
3       4       5       6       7       8
4       5       6       7       8       9
5       6       7       8       9       10
6       7       8       9       10      11
7       8       9       10      11      12
```

And the frequency distribution looks like this:

| Sum   | Frequency    |
|--------------- | --------------- |
|  2   | 1   |
|  3   | 2   |
|  4   | 3   |
|  5   | 4   |
|  6   | 5   |
|  7   | 6   |
|  8   | 5   |
|  9   | 4   |
|  10  | 3   |
|  11  | 2   |
|  12  | 1   |


The frequency distribution can be mathematically represented in the form of a function

$$
\text{freq}(\text{Sum}) = \begin{cases}
    Sum-1 \text{ if } 2 <= Sum <= 7 \\
    12 - Sum \text{ if } 7 < Sum <= 12\\
    \end{cases}
$$

, and can be represented in the form of code as follows,

```go 

```

#### Question 3

**Q: Calculate the Probability of all Possible Sums occurring among the number of combinations from (2).** 

##### Solution

The Probability distribution can be obtained by dividing the frequency of each sum by the total number of all possible combinations for a pair of dice, i.e. 36

$$
P(\text{ Sum }) = \text{freq(\text{Sum})} / 36
$$

```
Sum     Freq    Probability
2       1       0.027778 
3       2       0.055556 
4       3       0.083333 
5       4       0.111111 
6       5       0.138889 
7       6       0.166667 
8       5       0.138889 
9       4       0.111111 
10      3       0.083333 
11      2       0.055556 
12      1       0.027778 
```

```go
import "fmt"
func freq(Sum int) (int, error) {
    if Sum > 12 || Sum < 2 {
        return 0, errors.New("Sum can only be in range 2 ... 12")
    }

    if Sum <= 7 {
        return Sum-1, nil
    } else {
        return 12 - Sum + 1, nil
    }

}
func ques3()  {

    fmt.Println("Sum\tFreq\tProbability")
    for key := 2; key <= 12; key++ {
        f, _ := freq(key)
        fmt.Printf("%d:\t%d\t%f \n", key, f, float64(f)/36.0)
    }

}
```

### Part B

#### Problem Statement 

Now comes the real challenge. You were happily spending a lazy afternoon playing your board game with your dice when suddenly the mischievous Norse God Loki ( You love Thor too much & Loki didn’t like that much ) appeared.

Loki dooms your dice for his fun removing all the “Spots” off the dice.

No problem! You have the tools to re-attach the “Spots” back on the Dice.
However, Loki has doomed your dice with the following conditions:

- **Die A cannot have more than 4 Spots on a face** 
- **Die A may have multiple faces with the same number of spots.** 
- **Die B can have as many spots on a face as necessary i.e. even more than 6.** 
 
But in order to play your game, the probability of obtaining the Sums must remain the
same!

So if you could only roll P(Sum = 2) = 1/X, the new dice must have the spots reattached
such that those probabilities are not changed

#### Solution

The problem described here can be solved using [Sicherman Dice](https://en.wikipedia.org/wiki/Sicherman_dice). In this pair of dice, the faces are relabelled in such a way that the probability distribution is the same as that of a standard pair of dice.

> Sicherman dice /ˈsɪkərmən/ are a pair of 6-sided dice with non-standard numbers–one with the sides 1, 2, 2, 3, 3, 4 and the other with the sides 1, 3, 4, 5, 6, 8. They are notable as the only pair of 6-sided dice that are not normal dice , bear only positive integers, and have the same probability distribution for the sum as normal dice. They were invented in 1978 by George Sicherman of Buffalo, New York.


So, the solution to the given problem is **New_Die_A = {1,2,2,3,3,4}, New_Die_B = {1,3,4,5,6,8}**

So, the code for the given solution is as follows:

```go 

import (
	"errors"
)
func undoom(dieA, dieB [6]int) ([6]int, [6]int, error) {
	var standard_dice = [6]int{1, 2, 3, 4, 5, 6}

	if (dieA != standard_dice) || (dieB != standard_dice) {
		return [6]int{}, [6]int{}, errors.New("Not standard pair of dice")
	}

	return [6]int{1, 2, 2, 3, 3, 4}, [6]int{1, 3, 4, 5, 6, 8}, nil
}
```
