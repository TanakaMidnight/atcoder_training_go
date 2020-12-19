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

func SplitIntStdin(delim string) (intSlices []int, err error) {
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

func main() {
	l1 := nextLine()
	vals, _ := SplitIntStdin(l1)
	var n int = vals[0]
	var maxN int = n

	var m int = vals[1]
	var t int = vals[2]

	var cafeList [][]int
	for i := 0; i <= m; i++ {
		l2 := nextLine()
		vals2, _ := SplitIntStdin(l2)
		cafeList = append(cafeList, vals2)
	}

	var mCount int = 0
	var flag string = "Yes"
	var i int = 0
	for i < t {
		if i >= cafeList[mCount][0] && i < cafeList[mCount][1] {
			i = i + cafeList[mCount][1] - cafeList[mCount][0]
			n = n + cafeList[mCount][1] - cafeList[mCount][0]
			if n > maxN {
				n = maxN
			}
		} else {
			n--
			i++
		}

		if i > cafeList[mCount][1] {
			if mCount < m-1 {
				mCount++
			}
		}
		if n <= 0 {
			flag = "No"
			break
		}
	}
	fmt.Println(flag)
}
