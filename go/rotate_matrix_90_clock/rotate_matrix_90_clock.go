package algorithms

// Rotate rotates a square matrix 90 degrees clockwise
func RotateMat90Clock(matrix [][]int) {
	n := len(matrix)
	if n <= 1 {
		return
	}

	// Iterate through layers
	for i := 0; i < n/2; i++ {
		// Iterate through elements in current layer
		for j := i; j < n-i-1; j++ {
			// Store top-left element
			temp := matrix[i][j]

			// Move bottom-left to top-left
			matrix[i][j] = matrix[n-1-j][i]

			// Move bottom-right to bottom-left
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]

			// Move top-right to bottom-right
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]

			// Move stored top-left to top-right
			matrix[j][n-1-i] = temp
		}
	}
}
