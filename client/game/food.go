/*
package core
модуль food
содержит объекты, которые змейка может съесть
*/

package game

import (
	"github.com/JoelOtter/termloop"
)

//CreateFood создать еду для змейки
func CreateFood(coord Coordinates) *food {
	food := new(food)
	food.Entity = termloop.NewEntity(1, 1, 1, 1)

	//разместим еду на игровом поле
	food.coord = coord

	return food
}

//Draw отвечает за отрисовку пищи на дисплее
func (food *food) Draw(screen *termloop.Screen) {
	//отрисовка на экране
	screen.RenderCell(food.coord.X,
		food.coord.Y,
		&termloop.Cell{Fg: termloop.ColorWhite, Bg: termloop.ColorBlack,
			Ch: rune('€')})
}
