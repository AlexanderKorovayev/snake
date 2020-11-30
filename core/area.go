/*
package core
модуль area
отвечает за отричовку игрового поля.
небольшое описание для termloop:
	есть canvas(полотно) это двумерный массив, который содержит Cell(ячейки),
	они могут содержать разные символы и цвета, таким образом рисуется канвас.
	Что бы объект имел возможность попасть в канвас он должен отнаследоваться от
	энтити.
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
func CreateArea() *area {
	area := new(area)
	area.Entity = termloop.NewEntity(1, 1, 1, 1)
	border := make(map[coordinates]int)
	fillAreaBorder(width, high, &border)
	area.areaBorder = border
	return area
}

//fillAreaBorder заполнить игровую область информацией о её границах
func fillAreaBorder(imax, jmax int, border *map[coordinates]int) {
	starti := 0
	startj := 0
	for i := starti; i < imax; i++ {
		for j := startj; j < jmax; j++ {
			coord := coordinates{i, j}
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
func (area *area) Draw(screen *termloop.Screen) {
	//отрисовка на экране
	for k, v := range area.areaBorder {
		if v == 1 {
			screen.RenderCell(k.x, k.y, &termloop.Cell{Fg: termloop.ColorWhite,
				Bg: termloop.ColorWhite})
		}
	}
}

//Collision произошло ли косание с змейкой
func (area *area) collision(c *coordinates) bool {
	return area.areaBorder[*c] == 1
}
