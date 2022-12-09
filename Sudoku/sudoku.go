package main

func main() {
	var sudokuboard [81]int = [81]int{0,0,0,0,0,0,0,0,0,
		                              0,0,0,0,0,0,0,0,0,
	                                  0,0,0,0,0,0,0,0,0,
                                      0,0,0,0,0,0,0,0,0,
                                      0,0,0,0,0,0,0,0,0,
                                      0,0,0,0,0,0,0,0,0,
                                      0,0,0,0,0,0,0,0,0,
                                      0,0,0,0,0,0,0,0,0,
                                      0,0,0,0,0,0,0,0,0}

}


//constraints of the game
// no repeats in first row sudokuboard[0:8], second row sudokuboard[9:17], third row sudokuboard[18,26], so on...


// no repeats in the first column sudokuboard[0], 9, 18, ... 72, second column sudokuboard[1], 10, 19, ... 73, third column sudokuboard[2], 11, 20, ... 74, so on


// no repeats in the 3x3 squares sudoboard[0], 1, 2, 9, 10, 11, 18, 19, 20, second 3x3 square 3,4,5, 12, 13, 14, 21, 22, 23, so on


