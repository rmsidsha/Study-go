package main

import (
	"fmt"
	_ "strconv"
)

func main() {
	Solution(2)
}

func Solution(N int) {
	//fmt.Println(strconv.FormatInt(N, 2)) // int64를 string으로 바꿔줌. 확인하려면 N => int64로 바꿔야함
	start, count, max := 0, 0, 0
	for i := 0; N > 0; i++ {
		if N%2 == 1 {
			start = 1
			if max < count {
				max = count
			}
			count = 0
		} else if start == 1 {
			count++
		}
		N /= 2
	}
	fmt.Println(max)
}
