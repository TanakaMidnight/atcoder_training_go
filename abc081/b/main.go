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

func main() {
	l1, l2 := nextLine(), nextLine()
	_, _ = strconv.Atoi(l1)
	var count int = 0
	var flag bool = false
	vals, _ := splitIntStdin(l2)
	for {
		for i, x := range vals {
			if x%2 != 0 {
				flag = true
				break
			}
			vals[i] = x / 2
		}
		if flag {
			break
		}
		count++
	}

	fmt.Println(count)
}
