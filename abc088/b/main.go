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

func findSumOfDigits(x int) int {

	var sum int = 0
	for {
		sum += x % 10
		x = x / 10
		if x <= 0 {
			break
		}
	}

	return sum
}

func main() {
	_, l2 := nextLine(), nextLine()
	vals, _ := splitIntStdin(l2)
	var sumX = 0
	var sumY = 0

	var sortVals []int = vals
	// vals 並び替え

	for i, x := range sortVals {
		if i%2 == 0 {
			sumX += x
		} else {
			sumY += x
		}
	}

	fmt.Println(sumY - sumX)
}
