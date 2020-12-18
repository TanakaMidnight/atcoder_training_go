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
	l1:= nextLine();
	m, _ := strconv.Atoi(l1);
	
	if m < 100{
		fmt.Println("00");
	} else if m <= 5000 {
		fmt.Printf("%02d\n", m/100)
	}else if m <= 30000 {
		fmt.Printf("%02d\n", m/1000+50)
	} else if m <= 70000 {
		fmt.Printf("%02d\n", (m/1000-30)/5+80)
	} else {
		fmt.Println("89");
	}			
}
