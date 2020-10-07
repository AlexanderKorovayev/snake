package main

import (
	"fmt"
	"reflect"
)

type Test [][]int

func main() {
	test := [][]int{{1, 1}, {2, 3}}
	fmt.Println(test[1][1])
	fmt.Println(reflect.TypeOf(test))
	//game.Rungame()
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
