# WEIGHTED RANDOM MESSAGE

## How I solve this problem

To get a message from a set of version randomly and make it follow the proportion of its weight in the total weight of all versions, we need use an  auxiliary array which is the accumulated array.

Example:
```
input [ 4, 2, 1, 5, 3, 6]
accumulated [4, 6, 7, 12, 15, 21]
```

Why don't run randomly on input array?

To use this method, we will perform by using its index. In particular, I will pick a number randomly from `0 to 5`, then use this number as array's index to get the value of input[index]. By this method, the frequency of random numbers is equals around in 0.333. Thus, this method is not efficient to solve our problem.

I will try another method that more efficient.

Using an accumulated array that I mentioned above. By this method, I will pick a random number from `0 to 21`. The number recived before will give us an index, then use this index with the input array to find final number.

### Why this method efficiency?

I will devide accumulated array to small pices `(0,4]` - `(4,6]` - `(6,7]` - `(7,12]` - `(12, 15]` - `(15,21]`

Firstly, easy to see that space of 2 pieces `(6,7]`, `(15,21]` is 1, 6 respectively. So, the probability to get a number in `(15,21]` is bigger than `(6,7]` 6 times.

Secondly, the probability to get number from `0 to 21` is equal.

So, we can get random number in the input array with its weight proportion in total weight of version with accepted accuracy.

## How I tested it

To test this function, I will calculate the normal frequency of each number in the input array called `expected` and the appearence of a number after run  1000 times called `result` . I set expected Threshold is `0.05`.
If `abs(actual[i] -expected[i]` > `Threshold`, the test FAIL

Besides, I also test the input data.

## How to run this project

### Installation

```
$ git clone https://github.com/namnhce/weighted-random-messages.git
```

```
$ cd GOPATH/src/github.com/namnhce/weighted-random-messages
```

To run this project
```
$ make run INPUT_PATH="file_path"
```

example: `make run INPUT_PATH="/Users/namnh/Desktop/input.txt"`

`input.txt`
```
0 1 2 3 4 5 6 7 8 9
132413220 121411241 124115122 235124343 53432514 153262365 163343436 446313743 414822459 351615449
232413220 221411241 124115122 235124343 53432524 253262365 263343436 446323743 424822459 351625449
332413220 321411241 124115122 235124343 53432534 353262365 363343436 446333743 434822459 351635449
432413220 421411241 124115122 235124343 53432544 453262365 463343436 446343743 444822459 351645449
532413220 521411241 124115122 235124343 53432554 553262365 563343436 446353743 454822459 351655449
632413220 621411241 124115122 235124343 53432564 653262365 663343436 446363743 464822459 351665449
732413220 721411241 124115122 235124343 53432574 753262365 763343436 446373743 474822459 351675449
832413220 821411241 124115122 235124343 53432584 853262365 863343436 446383743 484822459 351685449
932413220 921411241 124115122 235124343 53432594 953262365 963343436 446393743 494822459 351695449
```

## To use it as a library
```
$ go get -u github.com/namnhce/weighted-random-messages/random
```

Implement:
```
package main

import (
	"fmt"

	"github.com/namnhce/weighted-random-messages/random"
)

func main() {
	input := []int64{50, 30, 60}
	rand := random.F(input)
	fmt.Println(rand)
}

```


