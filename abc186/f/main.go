package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func splitIntStdin(delim string) (intSlices []int, err error) {
	// 文字列スライスを取得
	stringSplited := strings.Split(delim, " ")

	// 整数スライスに保存
	for i := range stringSplited {
		var iparam int
		iparam, err = strconv.Atoi(stringSplited[i])
		if err != nil {
			return
		}
		intSlices = append(intSlices, iparam)
	}
	return
}

func findMin(min int, valx []int) int {
	var x = min
	for _, i := range valx {
		if i < x {
			x = i
		}
	}
	return x
}
func main() {
	l1 := nextLine()
	vars, _ := splitIntStdin(l1)
	h := vars[0]
	w := vars[1]
	m := vars[2]
	blocks := make([][]int, h)
	for i := range blocks {
		blocks[i] = make([]int, w)
	}

	for i := 0; i < m; i++ {
		lx := nextLine()
		valx, _ := splitIntStdin(lx)
		blocks[valx[0]-1][valx[1]-1] = 9
	}

	// 1手目
	for i := 0; i < h; i++ {
		if blocks[i][0] != 9 {
			blocks[i][0] = 1
		} else {
			break
		}
	}

	for i := 0; i < w; i++ {
		if blocks[0][i] != 9 {
			blocks[0][i] = 1
		} else {
			break
		}
	}

	//2手目

}
