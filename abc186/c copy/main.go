package main

import (
	"bufio"
	"fmt"
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
	vals, _ := splitIntStdin(l1)
	var blocks [][]int
	var sum = 0
	var min = 100
	for i := 0; i < vals[0]; i++ {
		lx := nextLine()
		valx, _ := splitIntStdin(lx)
		min = findMin(min, valx)
		blocks = append(blocks, valx)
	}
	for _, i := range blocks {
		for _, j := range i {
			if j > min {
				sum = sum + (j - min)
			}
		}
	}

	fmt.Println(sum)
}
