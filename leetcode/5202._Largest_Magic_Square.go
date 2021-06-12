// 一个 k x k 的 幻方 指的是一个 k x k 填满整数的方格阵，且每一行、每一列以及两条对角线的和 全部相等 。幻方中的整数 不需要互不相同 。显然，每个 1 x 1 的方格都是一个幻方。

// 给你一个 m x n 的整数矩阵 grid ，请你返回矩阵中 最大幻方 的 尺寸 （即边长 k）。

package main

import "fmt"

func main() {
	fmt.Println(largestMagicSquare([][]int{{7, 1, 4, 5, 6}, {2, 5, 1, 6, 4}, {1, 5, 4, 3, 2}, {1, 2, 7, 3, 4}}))
}

func largestMagicSquare(grid [][]int) int {
	c := len(grid)
	r := len(grid[0])
	max := 0
	// 正方形子集
	maxC := c
	if c > r {
		maxC = r
	}
	for d := 1; d <= maxC; d++ {
		// c为边长
		for i := 0; i <= c-d; i++ { // 行
			for j := 0; j <= r-d; j++ { // 列
				tempGrid := [][]int{}
				for m := i; m < i+d; m++ {
					tempGrid = append(tempGrid, grid[m][j:j+d])
				}
				if isMagicSquare(tempGrid) && len(tempGrid) > max{
					max = len(tempGrid)
				}
			}
		}
	}
	return max
}

func isMagicSquare(grid [][]int) bool {
	length := len(grid)
	if length == 1 {
		return true
	}
	if length == 0 || len(grid[0]) == 0 || len(grid) != len(grid[0]) {
		return false
	}
	sum := 0
	for i := 0; i < len(grid[0]); i++ {
		sum += grid[0][i]
	}
	// 横
	for i := 0; i < length; i++ {
		temp := 0
		for j := 0; j < length; j++ {
			temp += grid[i][j]
		}
		if temp != sum {
			return false
		}
	}
	// 竖
	for i := 0; i < length; i++ {
		temp := 0
		for j := 0; j < length; j++ {
			temp += grid[j][i]
		}
		if temp != sum {
			return false
		}
	}
	// \斜
	{
		temp := 0
		for i := 0; i < length; i++ {
			temp += grid[i][i]
		}
		if temp != sum {
			return false
		}
	}
	// /斜
	{
		temp := 0
		for i := 0; i < length; i++ {
			temp += grid[length-i-1][i]
		}
		if temp != sum {
			return false
		}
	}
	return true
}
