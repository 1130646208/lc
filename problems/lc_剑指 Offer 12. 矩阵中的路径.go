package main

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	var visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if backTrace(board, visited, i, j, word, 0) {
				return true
			}
		}
	}

	return false
}

func backTrace(board [][]byte, visited [][]bool, m, n int, word string, currentIdx int) bool {
	if m >= len(board) || m < 0 || n >= len(board[0]) || n < 0 {
		return false
	}

	if word[currentIdx] != board[m][n] {
		return false
	}

	if visited[m][n] {
		return false
	}

	if currentIdx == len(word)-1 {
		return true
	}

	visited[m][n] = true
	res := backTrace(board, visited, m+1, n, word, currentIdx+1) || backTrace(board, visited, m-1, n, word, currentIdx+1) || backTrace(board, visited, m, n+1, word, currentIdx+1) || backTrace(board, visited, m, n-1, word, currentIdx+1)
	visited[m][n] = false

	return res
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	exist(board, "ABCCED")
}
