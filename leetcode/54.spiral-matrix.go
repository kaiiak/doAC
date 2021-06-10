package main

import (
	"fmt"
)

func main() {
	fmt.Println("result:", spiralOrder([][]int{{1, 2, 3} /*, {4, 5, 6}, {7, 8, 9}*/}))
}

func spiralOrder(matrix [][]int) []int {
	var (
		output                                       []int
		limitTop, limitBottom, limitLeft, limitRight int
		indexX, indexY                               int
		spiralFlag                                   int
		isEnd                                        = func() bool {
			fmt.Printf("x:%v,y:%v,top:%v,bottom:%v,left:%v,right:%v\n", indexX, indexY, limitTop, limitBottom, limitLeft, limitRight)
			if limitLeft == limitRight || limitTop == limitBottom {
				return true
			}
			return false
		}
		isBreak bool
	)
	limitRight = len(matrix)
	if limitBottom == 0 {
		return output
	}
	limitBottom = len(matrix[0])
	output = make([]int, 0, limitBottom*limitRight)
	for {
		switch spiralFlag {
		case 0:
			// 上边
			fmt.Println("上边")
			output = append(output, matrix[indexX][indexY])
			if isEnd() {
				isBreak = true
			} else if indexX == limitBottom-1 {
				spiralFlag = 1
			} else {
				indexY++
				limitRight--
			}
		case 1:
			// 右边
			fmt.Println("右边")
			output = append(output, matrix[indexX][indexY])
			if isEnd() {
				isBreak = true
			} else if indexY == limitRight-1 {
				spiralFlag = 2
			} else {
				indexX++
				limitBottom--
			}
		case 2:
			// 下边
			fmt.Println("下边")
			output = append(output, matrix[indexX][indexY])
			if isEnd() {
				isBreak = true
			} else if indexX == limitBottom-1 {
				spiralFlag = 3
			} else {
				indexY--
				limitLeft++
			}
		case 3:
			// 左边
			output = append(output, matrix[indexX][indexY])
			if isEnd() {
				isBreak = true
			} else if indexX == limitBottom-1 {
				spiralFlag = 0
			} else {
				indexX--
				limitTop++
			}
		}
		if isBreak {
			break
		}
    }
	return output
}
