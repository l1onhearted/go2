package main

//04.二维数组查找
func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 && matrix[0][0] == target {
		return true
	} else if len(matrix) == 0 {
		return false
	}
	for i, j := 0, len(matrix[0])-1; i <= len(matrix)-1 && j > 0; {
		if target < matrix[i][j] {
			// i = i - 1
			j = j - 1
		} else if target > matrix[i][j] {
			i = i + 1
		} else {
			return true
		}
	}
	return false
}
