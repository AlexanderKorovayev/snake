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
	GameArea *Area
	Snake    *Snake
	GameFood *Food
}

//Area объект игрового поля, по которому будет перемещаться змейка
type Area struct {
	*termloop.Entity
	areaBorder map[Coordinates]int // для нахождения коллисий с краем поля и отрисовки граней
}

//Snake объект змейки
type Snake struct {
	*termloop.Entity
	Body      []Coordinates
	Direction direction
}

//Food объект пищи для змейки
type Food struct {
	*termloop.Entity
	coord Coordinates
}
