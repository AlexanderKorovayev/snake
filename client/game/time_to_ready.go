/*
package core
модуль time_to_ready
содержит объекты, которые змейка может съесть
*/

package game

import (
	"github.com/JoelOtter/termloop"
)

//CreateTimeObj создать отрисовку обратного отсчёта
func CreateTimeObj(val rune) *timeToReady {
	timeObj := new(timeToReady)
	timeObj.Entity = termloop.NewEntity(1, 1, 1, 1)

	//разместим время на игровом поле
	timeObj.coord = Coordinates{(width / 2) - 1, (high / 2) - 1}
	timeObj.value = val
	return timeObj
}

//Draw отвечает за отрисовку пищи на дисплее
func (timeToReady *timeToReady) Draw(screen *termloop.Screen) {
	//отрисовка на экране
	screen.RenderCell(timeToReady.coord.X,
		timeToReady.coord.Y,
		&termloop.Cell{Fg: termloop.ColorWhite, Bg: termloop.ColorBlack,
			Ch: timeToReady.value})
}

//setValue задать значение времени
func (timeToReady *timeToReady) setValue(value rune) {
	//установить новые координаты для еды.
	timeToReady.value = value
}
