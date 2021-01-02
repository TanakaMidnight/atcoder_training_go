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

func findSumOf7(x int) bool {
	var flag bool = false
	for {
		if x%10 == 7 {
			flag = true
			break
		}
		x = x / 10
		if x <= 0 {
			break
		}
	}

	return flag
}

func check7(x int) int {
	//10進数
	r10 := findSumOf7(x)
	//8進数
	s := fmt.Sprintf("%o", x)
	x8, _ := strconv.Atoi(s)
	r8 := findSumOf7(x8)
	if r10 || r8 {
		return 0
	}
	return 1
}
func main() {
	l1 := nextLine()
	x, _ := strconv.Atoi(l1)
	var sum = 0
	for i := x; i > 0; i-- {
		sum += check7(i)
	}
	fmt.Println(sum)
}
