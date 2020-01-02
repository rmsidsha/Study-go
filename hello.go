package main

import "fmt"

func main() {

	/*================================== switch, case
	var name string
	var category = 1

	switch category {
	case 1:
		name = "seo"
	case 2:
		name = "two"
	case 3:
		name = "Three"
	case 4:
		name = "Four"
	default:
		name = "Default"
	}

	println(name)
	=================================== for range
	names := []string{"홍길동", "서은우", "라쿠텐"}

	for index, name := range names {
		println(index, name)
	}
	*/
	/*=================================== break, continue, goto
	  	var a = 1
	  	for a < 15 {
	  		if a == 5 {
	  			a += a
	  			continue //  go to starting point of for loop
	  		}
	  		a++
	  		if a > 10 {
	  			break // escape for loop
	  		}
	  		println(a)
	  	}

	  	if a == 11 {
	  		goto END // go to END
	  	}
	  END:
	  	println("END")
	*/

	/* break label
	i := 0
	   L1:
	   	for {
	   		if i == 0 {
	   			break L1
	   		}
	   	}
	   	println("ok")
	*/

	/*
		msg := "Hello"
		say(&msg) // &를 사용하면 주소를 표시함.
		// 이때 Pass by reference로 문자열이 아니라 그 주소를 갖게 됨
		println(msg)

		say("This", "is", "a", "book")
	*/

	/*
		sum := func(n ...int) int {
			s := 0
			for _, i := range n {
				s += i
			}
			return s
		}
		result := sum(1, 2, 3, 4, 5)
		fmt.Print(reflect.TypeOf(result).Kind())
		println(result)
	*/

	//==========================================함수의 사용
	/*
			add := func(i int, j int) int {
				return i + j
			}

			r1 := cal(add, 10, 20)
			println(r1)

			r2 := cal(func(x int, y int) int { return x - y }, 10, 20)
			println(r2)
		}

		type calculator func(int, int) int //원형 정의 Delegate

		func cal(f calculator, a int, b int) int {
			result := f(a, b)
			return result
		}
	*/

	// sum1 과 sum2는 같다
	/*
	   func sum1(nums ...int) (int, int) {
	   	i := 0
	   	count := 0
	   	for _, a := range nums {
	   		i += a
	   		count++
	   	}
	   	return count, i
	   }

	   func sum2(nums ...int) (count int, total int) { // Named return Parameter
	   	for _, a := range nums {
	   		total += a
	   		count++
	   	}
	   	return //생략하면 에러 발생
	   }
	*/

	//function
	/*
	   func say(msg *string) { // 파라매터가 포인터임을 표시
	   	println(*msg) // *사용으로 주소의 데이터를 사용 이를 Dereferencing이라고 함
	   	*msg = "Changew"

	   }

	   func say(msg ...string) {
	   	for _, s := range msg {
	   		pritnnl(s)
	   	}
	   }
	*/
	//=========================================== closure
	/*
		next := nextValue()

			println(next())
			println(next())
			println(next())

			anotherNext := nextValue()
			println(anotherNext())
			println(anotherNext())

			println(next())
		}

		func nextValue() func() int {
			i := 0
			return func() int {
				i++
				return i
			}
		}
	*/

	/*======================================================== array
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	println(a[1])

	var b = [...]int{1, 2, 3, 4} // 배열 크기가 자동으로
	// c := func(nums ...in`t) (sum int) {
	// 	for _, n := range nums {
	// 		sum += n
	// 	}
	// 	return
	// }
	fmt.Println(b)
	// println(b)

	var multiArray1 [3][4][5]int // 정의
	multiArray1[0][1][2] = 10    // 사용
	fmt.Println(multiArray1)

	var multiArray2 = [2][3]int{
		{1, 2, 3},
		{4, 5, 6}, // 끝에 ,가 있어야함
	}

	println(multiArray2[1][2])
	*/

	슬라이스는 배열에 기초하여 만들어 졌지만 배열과 달리 고정된 크기를 미리 지정하지 않을 수 있고, 크기를 동적으로 변경할 수도 있고
	부분 배열을 발췌할 수도 있다.
	var a []int        // 슬라이스 변수 선언
	a = []int{1, 2, 3} // 슬라이스에 리터럴 값 지정
	a[1] = 10
	fmt.Println(a) // output: [1, 10, 3]

	var s1 = make([]int, 5, 10) // 첫번째 파라미터: 슬라이스 타입, 두번째: 길이, 세번째: 최대길이
	println(len(s1), cap(s1))

	var s2 []int // 슬라이스에 길이와 용량을 지정하지 않을 시, 기본적으로 0,0인 슬라이스를 만들고 이를 Nil Slice라고 한다
	// nil 와 비교를 하면 true를 리턴 함
	if s2 == nil {
		println("Nil Slice")
	}
	println(len(s2), cap(s2))
}
