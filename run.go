package main

import (
	"fmt"
	"reflect"
)

func main() {
	test := 5
	//fmt.Println(&test)
	test5(&test)
	//game.Rungame()
}

func test5(test *int) {
	fmt.Println(reflect.TypeOf(test), test)
	fmt.Println(reflect.TypeOf(*test), *test)
}

/*
есть canvas(полотно) это двумерный массив, который содержит Cell(ячейки),
они могут содержать разные символы и цвета, таким образом рисуется канвас
*/

/*
основные модули это:
	entity.go
	input.go
	level.go
	screen.go
	termloop.go
*/
