package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	l1, l2, l3, l4 := nextLine(), nextLine(), nextLine(), nextLine()
	a, _ := strconv.Atoi(l1)
	b, _ := strconv.Atoi(l2)
	c, _ := strconv.Atoi(l3)
	x, _ := strconv.Atoi(l4)
	var count int = 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				var calc = i*500 + j*100 + k*50
				if calc == x {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
