package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	l1 := nextLine()
	l1 = strings.Replace(l1, "a", "", -1)
	l1 = strings.Replace(l1, "i", "", -1)
	l1 = strings.Replace(l1, "u", "", -1)
	l1 = strings.Replace(l1, "e", "", -1)
	l1 = strings.Replace(l1, "o", "", -1)
	fmt.Println(l1)
}
