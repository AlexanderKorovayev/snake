/*
package core
модуль main_objects
содержит описание типов основных игровых объектов
*/

package game

import (
	"github.com/JoelOtter/termloop"
)

//game основной уровень игры
type game struct {
	termloop.Level
	gameArea    *area
	snake1      *snake
	snake2      *otherSnake
	snake3      *otherSnake
	snake4      *otherSnake
	gameFood    *food
	timeToReady *timeToReady
}

//area объект игрового поля, по которому будет перемещаться змейка
type area struct {
	*termloop.Entity
	areaBorder map[Coordinates]int // для нахождения коллисий с краем поля и отрисовки граней
}

//snake объект змейки
type snake struct {
	*termloop.Entity
	body  []Coordinates
	drctn direction
	name  string
	dead  bool
	color termloop.Attr
}

//snake объект змейки
type otherSnake struct {
	*termloop.Entity
	body  []Coordinates
	name  string
	color termloop.Attr
}

//food объект пищи для змейки
type food struct {
	*termloop.Entity
	coord Coordinates
}

//timeToReady объект для обратного отсчёта до начала игры
type timeToReady struct {
	*termloop.Text
}
