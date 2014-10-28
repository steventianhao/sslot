package engine

/*this is the lines for 3x5 slot, but the matrix is encoded as below
0 0 0 0 0
1 1 1 1 1
2 2 2 2 2
so for example {1, 1, 1, 1, 1} is the line from left to right in the middle, 5 ones
*/
var lines3x5 = [][]int{
	[]int{1, 1, 1, 1, 1},
	[]int{0, 0, 0, 0, 0},
	[]int{2, 2, 2, 2, 2},

	[]int{0, 1, 2, 1, 0},
	[]int{2, 1, 0, 1, 2},

	[]int{0, 0, 1, 0, 0},
	[]int{2, 2, 1, 2, 2},

	[]int{1, 2, 2, 2, 1},
	[]int{1, 0, 0, 0, 1},
	[]int{1, 0, 1, 0, 1},

	[]int{1, 2, 1, 2, 1},
	[]int{0, 1, 0, 1, 0},
	[]int{2, 1, 2, 1, 2},
	[]int{1, 1, 0, 1, 1},
	[]int{1, 1, 2, 1, 1},
}
