/*
package core
модуль area
отвечает за отричовку игрового поля.
небольшое описание для termloop:
	есть canvas(полотно) это двумерный массив, который содержит Cell(ячейки),
	они могут содержать разные символы и цвета, таким образом рисуется канвас
	основные модули это:
		entity.go
		input.go
		level.go
		screen.go
		termloop.go
*/

package core

import "github.com/JoelOtter/termloop"

//CreateArea создать арену, по которой будет перемещаться змейка
func CreateArea() *Area {
	area := new(Area)
	area.Entity = termloop.NewEntity(1, 1, 1, 1)
	border := make(map[Coordinates]int)
	fillAreaBorder(35, 20, &border)
	area.areaBorder = border
	return area
}

//fillAreaBorder заполнить игровую область информацией о её границах
func fillAreaBorder(imax, jmax int, border *map[Coordinates]int) {
	starti := 0
	startj := 0
	for i := starti; i < imax; i++ {
		for j := startj; j < jmax; j++ {
			coord := Coordinates{i, j}
			if i == starti || i == imax-1 {
				(*border)[coord] = 1
			} else if j == startj || j == jmax-1 {
				(*border)[coord] = 1
			} else {
				(*border)[coord] = 0
			}
		}
	}
}

//Draw отвечает за отрисовку змеи на дисплее
func (area *Area) Draw(screen *termloop.Screen) {
	//отрисовка на экране
	for k, v := range area.areaBorder {
		if v == 1 {
			screen.RenderCell(k.X, k.Y, &termloop.Cell{Fg: termloop.ColorWhite,
				Bg: termloop.ColorWhite})
		}
	}
}

//Collision произошло ли косание с змейкой
func (area *Area) Collision(c *Coordinates) bool {
	return area.areaBorder[*c] == 1
}
