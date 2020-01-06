package main

import "fmt"

/*
import _ "other/xlib"

import 하면서 alias를 씀. _alias_ "package_name"
package를 가져오고 사용하지 않을 때 alias를 _로 설정하면 컴파일 에러가 나지 않음

package 이름이 동일하지만 다른 버젼, 다른 위치에서 로딩하고자 할 때에도 alias를 사용해서 구분 가능
ex)
	import (
		mongo "other/mongo/db"
		mysql "other/mysql/db"
	)
*/

/*
func init() {
	패키지를 실행할 시 처음 호출되는 함수임.
	별도의 호출없이 자동으로 실행 됨
	ex) var pop map[int]string
		func init() {
			pop = make(map[int]string)
		}
	패키지 로드 시 pop 초기화
}
*/

func main() {
	//==================================== Slice
	/*
		s := []int{0, 1, 2, 3, 4, 5}
		s = s[2:5]
		s = s[1:]
		fmt.Println(s)
	*/
	/*
		sliceA := []int{1, 2, 3}
		sliceB := []int{4, 5, 6}
		sliceA = append(sliceA, sliceB...) // 그냥 sliceB 해주면 안됨
		fmt.Println(sliceA)
		fmt.Println(sliceB)
	*/

	//===================================== Map
	// Map은 reference 타입으로 nil값을 가짐. 이것을 Nil Map이라 부름
	// Nil Map은 어떤 데이터를 쓸 수 없는데 map을 초기화 하기 위해서 make()를 사용
	// 사용법: map[KEY_TYPE]VALUE_TYPE

	// make를 통한 할당
	/*
		var idMap1 map[int]string
		idMap1 = make(map[int]string) // key = int, value =string
		idMap1[1] = "Map1 ONE"
		idMap1[2] = "Map1 TWO"
		idMap1[3] = "Map1 THREE"

		str1 := idMap1[1]
	*/

	// make(map[KEY_TYPE]VALUE_TYPE)를 지정할 때 make는 해시테이블 자료구조를 메모리에 생성하고
	// 그 메모리를 가리키는 map value를 리턴함. (map value는 내부적으로 runtime.hmap 구조체를 가리키는 포인터)
	// 즉 idMap 변수는 이 해시테이블을 가리키는 map을 가리키게 됨

	// 리터럴을 통한 할당
	/*
		idMap2 := map[int]string{
			1: "ONE",
			2: "TWO",
			3: "THREE",
		}
		str2 := idMap2[1]

		fmt.Println(idMap1)
		fmt.Println(idMap1[1])
		fmt.Println(str1)
		fmt.Println(idMap2)
		fmt.Println(idMap2[1])
		fmt.Println(str2)

		// val에는 value값이 exists에는 bool 값이 담기게 됨. 있으면 true, 없으면 false
		val, exists := idMap2[3]
		if !exists {
			fmt.Println("No")
		} else {
			fmt.Println(val)
		}

		// loop를 사용한 map 열거
		// for range 문을 사용하여 모든 맵 요소 출력
		// map은 unordered이므로 순서는 무작위
		for key, val := range idMap1 {
			fmt.Println(key, val)
		}
		fmt.Println(tp.getString("seo"))
	*/
	//====================================================================================== struct
	// struct는 type문을 사용하여 정의 할 수 있따.
	// struct는 필드의 집합체이며 필드의 컨테이너이다.
	// struct는 필드만을 가지며 메소드는 별도로 분리하여 정의한다
	// struct는 mutable 개체이고 필드 값 변경시 메모리에서 직접 변경됨
	// But parameter로 넘겨지면 Pass by value로 인해 객체를 복사해서 전달함.
	// 때문에 Pass by reference로 전달하려면 struct의 포인터를 전달해야함

	p := Person{}  // Person 객체를 생성
	p.name = "leo" // 필드 값 설정
	p.age = 16
	/* or
	일부 필드가 생략될 경우 생략된 필드들은 Zero value를 갖는다
	int=0,  float=0.0,  string="",  pointer=nil
	new()를 사용하여 생성하면 모든 필드를 Zero value로 초기화하고 해당 객체의 포인터(*Object)를 리턴한다
	포인터에 액세스 할 때에는 .(dot)을 사용하며 이때 포인터는 자동으로 Dereference된다.

	var p1 Person
	p1 = Person{"leo", 16}
	p2 := Person{name: "kim", age: 17}
	p3 := new(Person)
	p3.name = "seo"
	p3.age = 26
	*/

	fmt.Println(p)

	dic := newDict() // 생성자 호출
	dic.data[1] = "A"
	fmt.Println(dic)
}

type Dict struct {
	data map[int]string
}

func newDict() *Dict {
	d := Dict{}
	/* or
	var d Dict
	d = Dict{"val1", "val2", ...}
	d := Dict{feild: "val1", ...}
	d := new(Dict)
	d.feild = ""
	*/
	d.data = map[int]string{}
	return &d // 포인터 전달
}

type Person struct { // 첫 글자를 대문자로 설정하면 외부에서도 사용 가능
	name string
	age  int
}
