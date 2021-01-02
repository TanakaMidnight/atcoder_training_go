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
	t, _ := strconv.Atoi(l1)
	var blocks [][]int
	for i := 0; i < t; i++ {
		lx := nextLine()
		valx, _ := splitIntStdin(lx)
		blocks = append(blocks, valx)
	}

	for _, i := range blocks {
		var count = 0
		var out = 0
		n := i[0]
		s := i[1]
		k := i[2]
		for {
			if s%n == 0 {
				break
			}
			s = s + k
			if s > n {
				s = s - n

				out++
			}
			if out >= n {

				count = -1
				break
			}
			count++
		}
		fmt.Println(count)

	}

}
