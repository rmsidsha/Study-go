package main

import "fmt"

func main() {
	var A = []int{1, 2, 3, 4, 5}
	K := 11
	fmt.Println(Solution(A, K))
}

func Solution(A []int, K int) []int {
	length := len(A)
	if length > 0 {
		if K > length {
			K = K % length
		}
		A = append(A[length-K:length], A[0:length-K]...)
	}
	return A
}
