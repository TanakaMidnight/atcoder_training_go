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

func splitFloatStdin(delim string) (floatSlices []float64, err error) {
	// 文字列スライスを取得
	stringSplited := strings.Split(delim, " ")

	// 整数スライスに保存
	for i := range stringSplited {
		var iparam float64
		iparam, err = strconv.ParseFloat(stringSplited[i], 64)
		if err != nil {
			return
		}
		floatSlices = append(floatSlices, iparam)
	}
	return
}

func main() {
	l1 := nextLine()
	vals, _ := splitFloatStdin(l1)
	var x1 = vals[0]
	var x2 = vals[1]
	var y1 = vals[2]
	var y2 = vals[3]
	var z1 = vals[4]
	var z2 = vals[5]

	y1 = y1 - x1
	z1 = z1 - x1
	y2 = y2 - x2
	z2 = z2 - x2

	var result = (y1*z2 - y2*z1) / 2
	if result > 0 {
		fmt.Println(result)
	} else {
		fmt.Println(result * -1)
	}
}
