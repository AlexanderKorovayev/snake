/*
package core
модуль main_objects
содержит описание типов основных игровых объектов
*/

package core

import (
	"github.com/JoelOtter/termloop"
)

//Game основной объект игра, который содержит в себе все остальные
type Game struct {
	termloop.Level
	GameArea *area
	Snake    *snake
	GameFood *Food
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

//Food объект пищи для змейки
type Food struct {
	*termloop.Entity
	coord coordinates
}
