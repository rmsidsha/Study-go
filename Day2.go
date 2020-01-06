package main

import (
	"fmt"
	"strconv"
)

// strconv: 형변환을 하기 위해 사용하는 패키지.
// Atoi: String to int    /      ItoA: Int to string

func main() {

	//============================================ slice
	// s1 := []int{0, 1, 2, 3, 4, 5}
	// var s2 = s1[2:5]
	// fmt.Println(s1)
	// fmt.Println(s2)

	// s3 := []int{0, 1}
	// s3 = append(s3, 3)
	// fmt.Println(s3)
	// s3 = append(s3, 4, 5)
	// fmt.Println(s3)

	// sliceA := make([]int, 0, 4) // len=0, cap=4
	// for i := 1; i <= 15; i++ {
	// 	sliceA = append(sliceA, i)
	// 	fmt.Println(len(sliceA), cap(sliceA))
	// }
	// fmt.Println(sliceA)

	sliceB := make([]string, 0, 4)
	for i := 1; i <= 15; i++ {
		sliceB = append(sliceB, strconv.Itoa(i))
	}
	fmt.Println(sliceB)

}
