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
	n, _ := strconv.Atoi(l1)
	vars, _ := splitIntStdin(l2)
	var calcs [][2]int
	for i := 0; i < n-1; i++ {
		x := vars[i]
		for j := i + 1; j < len(vars); j++ {
			calcs = append(calcs, [2]int{x, vars[j]})
		}

	}
	var sum = 0
	for _, y := range calcs {
		var z = y[0] - y[1]
		if z < 0 {
			z = z * -1
		}
		sum += z
	}
	fmt.Println(sum)

}
