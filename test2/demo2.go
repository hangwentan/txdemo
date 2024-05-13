package main

import (
	"fmt"
)

// 从上到下找到最短路径（n个数字之和最小,n为矩阵的行数），可以从第一行中的任何元素开始，只能往下层走，同时只能走向相邻的节点，例如图中第一排 2 只能走向 第二排的 7、3；第二排的 7 可以走向第三排的 6、2、9
//
// | 5    | 8    | 1    | 2    |
// | 4    | 1    | 7    | 3    |
// | 3    | 6    | 2    | 9    |
//
// Input: [
//     [5, 8, 1, 2],
//     [4, 1, 7, 3],
//     [3, 6, 2, 9]
// ]
// Output: 4

func main() {
	matrix := [][]int{
		{5, 8, 1, 2},
		{4, 1, 7, 3},
		{3, 6, 2, 9},
	}

	//nums := make(map[int]int, 0)
	//
	//for i, ints := range matrix {
	//
	//}
	fmt.Println(matrix)
	fmt.Println(minPathSum(matrix))
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, columns := len(grid), len(grid[0])
	dp := make([][]int, rows)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, columns)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < columns; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[rows-1][columns-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
