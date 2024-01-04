package main

import "fmt"

func exist(board [][]byte, word string) bool {
	var path []byte
	var r bool
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != word[0] {
				continue
			}
			trace(board, i, j, path, word, 0, &r)
			if r {
				return r
			}
		}
	}
	return r
}

func trace(board [][]byte, x, y int, path []byte, word string, ind int, result *bool) {
	if x >= len(board) || x < 0 || y >= len(board[0]) || y < 0 || ind >= len(word) {
		return
	}
	if board[x][y] != word[ind] {
		return
	}
	path = append(path, board[x][y])
	board[x][y] = ' '
	trace(board, x+1, y, path, word, ind+1, result)
	trace(board, x-1, y, path, word, ind+1, result)
	trace(board, x, y+1, path, word, ind+1, result)
	trace(board, x, y-1, path, word, ind+1, result)
	board[x][y] = word[ind]
	// 注意判断时机
	if len(path) == len(word) {
		*result = true
	}
	path = path[:len(path)-1]
}

func main() {
	s := exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "BCC")
	// s := exist([][]byte{{'A'}}, "A")
	fmt.Print(s)
}
