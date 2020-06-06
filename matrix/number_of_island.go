package matrix

// NumberOfIsland ... returns the number of connected one's called as island.
// 1 is land
// 0 is water
// ISLAND: A group of connected 1s (Horizontally or Vertically or Diagonally) forms an island.
func NumberOfIsland(matrix [][]int) int {

	row := len(matrix)
	if len(matrix) < 1 {
		return 0
	}
	col := len(matrix[0])

	// to check whether a element is visited or not
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}
	count := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 1 && !visited[i][j] {
				visitAllNeighbour(matrix, i, j, visited)
				count++
			}
		}
	}
	return count
}

// visitAllNeighbour ... is to visit
func visitAllNeighbour(matrix [][]int, row int, col int, visited [][]bool) {
	rowNbr := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	colNbr := []int{-1, 0, 1, -1, 1, -1, 0, 1}

	visited[row][col] = true

	for k := 0; k < 8; k++ {
		if isSafe(matrix, row+rowNbr[k], col+colNbr[k], visited) {
			visitAllNeighbour(matrix, row+rowNbr[k], col+colNbr[k], visited)
		}
	}
}

// isSafe ... should return true if the element is within matrix range and not visited.
func isSafe(matrix [][]int, row int, col int, visited [][]bool) bool {
	return (row >= 0) && (row < len(matrix)) && (col >= 0) &&
		(col < len(matrix[0])) && (matrix[row][col] == 1 && !visited[row][col])
}
