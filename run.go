package main

import (
	"github.com/AlexanderKorovayev/snake/game"
)

func main() {
	//"github.com/AlexanderKorovayev/snake/game"
	game.StartGame()
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
