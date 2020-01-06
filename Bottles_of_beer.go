package main

import "fmt"

func main() {
	/*
		99 bottles of beer on the wall, 99 bottles of beer.
		Take one down, pass it around, 98 bottles of beer on the wall.
		98 bottles of beer on the wall, 98 bottles of beer.
		Take one down, pass it around, 97 bottles of beer on the wall.
		...(생략)
		2 bottles of beer on the wall, 2 bottles of beer.
		Take one down, pass it around, 1 bottle of beer on the wall.
		1 bottle of beer on the wall, 1 bottle of beer.
		Take one down, pass it around, No more bottles of beer on the wall.
		No more bottles of beer on the wall, No more bottles of beer.
		Go to the store and buy some more, 99 bottles of beer on the wall.
	*/
	i := 99
	for i > 1 {
		str := "bottles"
		fmt.Printf("%d %s of beer on the wall, %d %s of beer. \n", i, str, i, str)
		i--
		if i == 1 {
			str = "bottle"
			fmt.Printf("Take one down, pass it around, %d %s of beer on ther wall \n", i, str)
			fmt.Printf("%d %s of beer on the wall, %d %s of beer. \n", i, str, i, str)
			fmt.Println("No more bottles of beer on the wall, No more bottles of beer.")
			fmt.Println("Go to the store and buy some more, 99 bottles of beer on the wall")
		} else {
			fmt.Printf("Take one down, pass it around, %d %s of beer on ther wall \n", i, str)
		}

	}
}
