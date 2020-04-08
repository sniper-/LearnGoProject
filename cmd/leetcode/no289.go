package leetcode

import "fmt"

func GameOfLife(board [][]int) {
	withEdge := make([][]int, len(board)+2)
	for k := 0; k < len(withEdge); k++ {
		withEdge[k] = make([]int, len(board[0])+2)
	}

	//构建一个同board值相同，但是周围都是0边界的镜像数组
	//1、初始化一个比board阵列大一圈的阵列，并赋值为0
	for i := range withEdge {
		for j := range withEdge[0] {
			withEdge[i][j] = 0
		}
	}
	//2、传入的阵列board的值
	for row := range board {
		for column := range board[row] {
			withEdge[row+1][column+1] = board[row][column]
		}
	}

	//以withEdge为参考，直接更新board中的值并输出即可
	for row := range board {
		for column := range board[row] {
			//边界值0并不影响判定结果，不需要考虑四边和四角边界处理
			count := 0
			count += withEdge[row][column]
			count += withEdge[row][column+1]
			count += withEdge[row][column+2]
			count += withEdge[row+1][column]
			count += withEdge[row+1][column+2]
			count += withEdge[row+2][column]
			count += withEdge[row+2][column+1]
			count += withEdge[row+2][column+2]

			switch count {
			//这些状态无论0,1都会变为0
			case 0, 1, 4, 5, 6, 7, 8:
				board[row][column] = 0
			//这些状态无论0,1都会变为1
			case 3:
				board[row][column] = 1
			//其余情况状态不变，不需要处理
			}
		}
	}

	fmt.Println(board)
}
