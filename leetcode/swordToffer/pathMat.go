package main

//12.矩阵路径
var dir = [][]int{
	[]int{-1, 0},
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(board[0]))
	}
	for i, v := range board {
		for j := range v {
			if dfs(board, visited, word, 0, i, j) {
				return true
			}
		}
	}
	return false
}

//判断是否在矩阵内
func inBoard(board [][]byte, x, y int) bool {
	return x >= 0 && x < len(board) && y >= 0 && y < len(board[0])
}

func dfs(board [][]byte, visted [][]bool, word string, index, x, y int) bool {
	if index == len(word)-1 {
		return board[x][y] == word[index] //判断当前矩阵元素是否对应
	}
	if board[x][y] == word[index] {
		visted[x][y] = true //访问过改元素
		for i := 0; i < 4; i++ {
			nx := x + dir[i][0]
			ny := y + dir[i][1]
			if inBoard(board, nx, ny) && !visted[nx][ny] && dfs(board, visted,
				word, index+1, nx, ny) {
				return true
			}
		}
		visted[x][y] = false
	}
	return false
}
