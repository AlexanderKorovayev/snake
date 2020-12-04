/*
package core
модуль main_objects
содержит описание типов основных игровых объектов
*/

package game

import (
	"github.com/JoelOtter/termloop"
)

//Game основной уровень игры
type Game struct {
	termloop.Level
	GameArea *area
	Snake    *snake
	GameFood *food
}

//area объект игрового поля, по которому будет перемещаться змейка
type area struct {
	*termloop.Entity
	areaBorder map[coordinates]int // для нахождения коллисий с краем поля и отрисовки граней
}

//snake объект змейки
type snake struct {
	*termloop.Entity
	body  []coordinates
	drctn direction
}

//food объект пищи для змейки
type food struct {
	*termloop.Entity
	coord coordinates
}
