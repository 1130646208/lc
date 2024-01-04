package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	col, row := len(matrix[0])-1, 0

	for {
		if target > matrix[row][col] {
			if row+1 <= len(matrix)-1 {
				row++
			} else {
				return false
			}
		} else if target < matrix[row][col] {
			if col-1 >= 0 {
				col--
			} else {
				return false
			}
		} else {
			return true
		}
	}
}

func main() {
	i := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	println(findNumberIn2DArray(i, 5))
}
