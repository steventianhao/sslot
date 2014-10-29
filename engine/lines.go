package engine

/*this is the lines for 3x5 slot, but the matrix is encoded as below
0 0 0 0 0
1 1 1 1 1
2 2 2 2 2
so for example {1, 1, 1, 1, 1} is the line from left to right in the middle, 5 ones
*/
var Lines3x5 = []Line{
	Line{1, 1, 1, 1, 1},
	Line{0, 0, 0, 0, 0},
	Line{2, 2, 2, 2, 2},

	Line{0, 1, 2, 1, 0},
	Line{2, 1, 0, 1, 2},

	Line{0, 0, 1, 0, 0},
	Line{2, 2, 1, 2, 2},

	Line{1, 2, 2, 2, 1},
	Line{1, 0, 0, 0, 1},
	Line{1, 0, 1, 0, 1},

	Line{1, 2, 1, 2, 1},
	Line{0, 1, 0, 1, 0},
	Line{2, 1, 2, 1, 2},
	Line{1, 1, 0, 1, 1},
	Line{1, 1, 2, 1, 1},
}
