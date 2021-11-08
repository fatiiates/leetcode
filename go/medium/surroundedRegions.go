package main

func solve(board [][]byte) {

	m := len(board)
	n := len(board[0])
	var onBorder map[int][]int = make(map[int][]int)
	var inBorder map[int][]int = make(map[int][]int)

	if m >= 3 && n >= 3 {

		/* onBorder */

		// first and last row
		for i := 0; i < m; i += m - 1 {
			arr := []int{}
			for j := 1; j < n-1; j++ {
				if board[i][j] == 'O' {
					arr = append(arr, j)
				}
			}
			if len(arr) > 0 {
				onBorder[i] = arr
			}
		}

		// first and last column
		for i := 1; i < m-1; i++ {
			arr := []int{}
			for j := 0; j < n; j += n - 1 {

				if board[i][j] == 'O' {
					arr = append(arr, j)
				}

			}
			if len(arr) > 0 {
				onBorder[i] = arr
			}

		}

		/* inBorder */

		for i := 1; i < m-1; i++ {
			arr := []int{}

			for j := 1; j < n-1; j++ {
				if board[i][j] == 'O' {
					arr = append(arr, j)
				}
			}
			if len(arr) > 0 {
				inBorder[i] = arr
			}

		}

		for x := range onBorder {
			for _, y := range onBorder[x] {
				findOnesWhichConnectedToBorder(x, y, &inBorder, &board)
			}
		}

		for x := range inBorder {
			for _, y := range inBorder[x] {
				board[x][y] = 'X'
			}
		}
	}

}

func main() {
	a := [][]byte{
		{'X', 'X', 'X', 'O', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'O', 'X', 'O', 'O'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'O'},
		{'X', 'X', 'X', 'X', 'O', 'X', 'O'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'O'},
		{'X', 'X', 'X', 'X', 'X', 'X', 'X'},
		{'O', 'X', 'X', 'O', 'O', 'O', 'X'},
	}
	solve(a)
}

// [["X","X","X","O","X","O","X"],
// ["X","O","X","O","X","O","O"],
// ["X","X","X","X","X","X","O"],
// ["X","X","X","X","O","X","O"],
// ["X","X","X","X","X","X","O"],
// ["X","X","X","X","X","X","X"],
// ["O","X","X","O","O","O","X"]]

func findOnesWhichConnectedToBorder(x int, y int, inBorder *map[int][]int, board *[][]byte) {

	if x == 0 {
		isConnectedToBorder(x+1, y, inBorder, board, "t")
	} else if y == 0 {
		isConnectedToBorder(x, y+1, inBorder, board, "l")
	} else if y == len((*board)[0])-1 {
		isConnectedToBorder(x, y-1, inBorder, board, "r")

	} else if x == len(*board)-1 {
		isConnectedToBorder(x-1, y, inBorder, board, "b")
	}
}

func indexOf(inBorder *map[int][]int, x int, y int) int {
	for k, v := range (*inBorder)[x] {
		if v == y {
			return k
		}
	}
	return -1 //not found.
}

func remove(s *map[int][]int, x int, i int) []int {
	(*s)[x][i] = (*s)[x][len((*s)[x])-1]
	return (*s)[x][:len((*s)[x])-1]
}

func isConnectedToBorder(x int, y int, inBorder *map[int][]int, board *[][]byte, from string) bool {

	if (*board)[x][y] == 'O' && indexOf(inBorder, x, y) != -1 {

		if len((*inBorder)[x]) == 1 {
			delete(*inBorder, x)
		} else {
			i := indexOf(inBorder, x, y)
			if i == -1 {
				return true
			}
			(*inBorder)[x] = remove(inBorder, x, i)
		}

	} else {
		return true
	}

	// left

	if from != "l" && indexOf(inBorder, x, y-1) != -1 && y-1 != 0 && (*board)[x][y-1] == 'O' {
		isConnectedToBorder(x, y-1, inBorder, board, "r")
	}

	// top
	if from != "t" && indexOf(inBorder, x-1, y) != -1 && x-1 != 0 && (*board)[x-1][y] == 'O' {
		isConnectedToBorder(x-1, y, inBorder, board, "b")
	}

	// right
	if from != "r" && indexOf(inBorder, x, y+1) != -1 && y+1 != len((*board)[0])-1 && (*board)[x][y+1] == 'O' {
		isConnectedToBorder(x, y+1, inBorder, board, "l")
	}

	// bottom
	if from != "b" && indexOf(inBorder, x+1, y) != -1 && x+1 != len(*board)-1 && (*board)[x+1][y] == 'O' {
		isConnectedToBorder(x+1, y, inBorder, board, "t")
	}

	return true

}
