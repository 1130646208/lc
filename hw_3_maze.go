package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	var (
		row, column        int
		maze               [][]int
		startRow, startCol = 0, 0
	)

	// read input
	_, _ = fmt.Scanf("%d %d", &row, &column)
	cin := bufio.NewReader(os.Stdin)
	maze = make([][]int, row)

	for r := 0; r < row; r++ {
		maze[r] = make([]int, column)
		line, _ := cin.ReadString('\n')
		split := strings.Split(line, " ")

		for c := 0; c < column; c++ {
			_, _ = fmt.Sscanf(split[c], "%d", &maze[r][c])
		}
	}
	// end read input

	var (
		endRow, endCol = row - 1, column - 1
		visited        = make([][]bool, row)
	)
	for i, _ := range visited {
		visited[i] = make([]bool, column)
	}

	runMaze(maze, visited, row, column, startRow, startCol, endRow, endCol)
}

func runMaze(maze [][]int, visited [][]bool, row, col, startRow, startCol, endRow, endCol int) {

	// 存储准备访问的位置
	queue := make([][2]int, 0, row*col)
	queue = append(queue, [2]int{startRow, startCol})
	footPrint := make([][2]int, 0)

	for len(queue) > 0 {
		nowVisiting := queue[0]
		r, c := nowVisiting[0], nowVisiting[1]
		visited[r][c] = true
		footPrint = append(footPrint, [2]int{r, c})

		if r == endRow && c == endCol {
			printPath(footPrint)
			return
		}

		queue = queue[1:]

		if inMazeAndNotVisited(r-1, c, row, col, visited, maze) {
			queue = append(queue, [2]int{r - 1, c})
		}
		if inMazeAndNotVisited(r+1, c, row, col, visited, maze) {
			queue = append(queue, [2]int{r + 1, c})
		}
		if inMazeAndNotVisited(r, c-1, row, col, visited, maze) {
			queue = append(queue, [2]int{r, c - 1})
		}
		if inMazeAndNotVisited(r, c+1, row, col, visited, maze) {
			queue = append(queue, [2]int{r, c + 1})
		}
	}
}

func inMazeAndNotVisited(r, c, row, col int, visited [][]bool, maze [][]int) bool {
	return r >= 0 && r < row && c >= 0 && c < col && !visited[r][c] && maze[r][c] == 0
}

func printPath(footPrint [][2]int) {
	last := len(footPrint) - 1
	res := make([]int, 0, len(footPrint))
	res = append(res, last)

	for i := last; i >= 0; i-- {
		if math.Abs(float64(footPrint[i][0]-footPrint[last][0])) == 1 && footPrint[i][1] == footPrint[last][1] ||
			math.Abs(float64(footPrint[i][1]-footPrint[last][1])) == 1 && footPrint[i][0] == footPrint[last][0] {
			last = i
			res = append(res, last)
		}
	}

	for i := len(res) - 1; i >= 0; i-- {
		fmt.Printf("(%d,%d)\n", footPrint[res[i]][0], footPrint[res[i]][1])
	}
}
