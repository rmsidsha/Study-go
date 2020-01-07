package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	r := Rect{10., 20.}
	c := Circle{10}
	showArea(r, c)

	var inter interface{}
	inter = 1
	inter = "Char"
	Empty1(inter)

	// Type Assertion
	// Interface type의 x와 type T에 대하여 x.(T)라고 표현을 하였을 때,
	// x가 nil이거나 x의 타입이 T가 아니라면 에러발생
	// x가 T타입인 경우는 T 타입의 x를 리턴한다.

	var a interface{} = 1
	i := a       // a 와 i는 Dynamic type 값은 1이 됨
	j := a.(int) // j는 int타입, 값은 1
	println(i)   // 포인터 주소 출력
	println(j)   // 1출력

	_, err := ontherError()
	switch err.(type) { // error type에 따른 처리 가능
	default: // error가 없을 경우 err == nil
		println("ok")
	case error1:
		println("error1")
	case error2:
		println("error2")
	case error: // 모든 에러는 error interface를 구현하기 때문에 이 문장은 모든 에러에 적용
		log.Fatal(err.Error()) // log.Fatal() 메세지를 출력하고 os.Exit(1)을 호출하여 프로그램을 종료함
	}

}

// struct는 필드의 집합
// interface는 메소드의 집합
type Rect struct {
	width, height float64
}
type Circle struct {
	radius float64
}

// Rect 타입에 대한 Shape interface 구현
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

// Circle 타입에 대한 Shape interface 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// interface 생성
type Shape interface {
	area() float64
	perimeter() float64
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		a := s.area() // 인터페이스 메소드 호출
		fmt.Println(a)
	}
}

// empty interface
// GOlang의 모든 타입은 적어도 0개의 메서드를 구현해야함
// 때문에 모든 Type을 나타내기 위하며 빈 인터페이스를 사용
// 어떠한 타입도 담을 수 있는 컨테이너이며 Dynamic Type이라고 할 수 있음.
// 사용법: interface{}라고 표시함
func Empty1(a interface{}) {
	fmt.Println(a)
}

// Error
// error라는 interface가 내장되어 있다.
// Error() string 이라는 하나의 메서드를 가지며 커스텀 가능함
type error interface {
	Error() string
}


// defer 키워드는 특정 문장 or 함수를 나중에 실행하게 함.(defer을 호출하는 함수가 리턴하기 직전)
// Clean up 작업을 위해 사용함
f, err := os.Open("file_name")
if err != nil {
	panic(err)	// panic은 현재 함수를 즉시 멈추고 현재 함수에 defer함수들을 모두 실행한 후 즉시 리턴한다.
				// panic모드 실행방식은 다시 상위함수에도 똑같이 적용되고 계속 콜스택을 타고 올라가며 적용된다
				// 결국 마지막에 에러를 내고 종료됨
}

defer f.Close() // main 마지막에 파일 close를 실행

bytes := make([]byte, 1024) // 파일을 읽어 옴
f.Read(bytes)
println(len(bytes))

func main() {
	openFile1("file")
	println("Done") // 이 문장은 실행안됨
					// openFile2()일 경우에는 실행됨
}
func openFile1(fn string) {
	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	// 파일 close 실행
	defer f.Close()
	}
}

func openFile2(fn string) {
	defer func() { // panic 실행시 호출됨
		if r := recover(); r != nil {
			fmt.Println("Open error", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

// recover()은 panic상태를 다시 정상으로 되돌리는 함수.
// panic 상태를 제거하고 다음 문장 실행 가능


//================================ go routine
//
// go routine을 사용하여 main routine은 계속 다음 문장을 실행(time.Sleep부분)
// go say()는 별도의 go routine들이 동작함.
func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	say("yes1") // 함수를 동기적으로 실행
	say("yes2")

	go say("Go yes1")	// 함수를 비동기적으로 실행
	go say("Go yes2")
	go say("Go yes3")

	//3초 대기
	time.Sleep(time.Second * 3)
}