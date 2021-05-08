package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {

	var n, count int
	fmt.Scan(&n)
	slice := make([]int, 0)
	amari := make([]int, 200)

	sc.Split(bufio.ScanWords)
	for j := 0; j < n; j++ {
		sc.Scan()
		i, e := strconv.Atoi(sc.Text())
		if e != nil {
			panic(e)
		}
		slice = append(slice, i)
		amari[slice[j]%200]++
	}

	for k := 0; k < 200; k++ {
		count += (amari[k] * (amari[k] - 1)) / 2
	}
	fmt.Println(count)

}
