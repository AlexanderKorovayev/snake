//модуль для описания основных объектов
package game

import (
	"github.com/JoelOtter/termloop"
)

//Game основной объект для рисования
type Game struct {
	termloop.Level
	gameArea *Area
	snake    *Snake
	gameFood *Food
}

//Area объект арены, по которой будет перемещаться змейка
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
