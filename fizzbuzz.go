package main

import "fmt"

func main() {
	// 	1에서 100까지 출력
	// 	3의 배수는 Fizz 출력
	// 	5의 배수는 Buzz 출력
	// 	3과 5의 공배수는 FizzBuzz 출력

	for i := 1; i < 101; i++ {
		// if i%3 == 0 && i%5 == 0 {
		// 	fmt.Println(i, "FizzBuzz")
		// 	continue
		// }

		// if i%3 == 0 {
		// 	fmt.Println(i, "Fizz")
		// } else if i%5 == 0 {
		// 	fmt.Println(i, "Buzz")
		// }

		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("fizzbuzz")
		case i%3 == 0:
			fmt.Println("fizz")
		case i%5 == 0:
			fmt.Println("buzz")
		default:
			fmt.Println(i)
		}
	}
}
