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
	//gameFoot  foot
}

//Area объект арены, по которой будет перемещаться змейка
type Area struct {
	*termloop.Entity
	areaBorder map[coordinates]int // для нахождения коллисий с краем поля и отрисовки граней
}

//Snake объект змейки
type Snake struct {
	*termloop.Entity
	Body      []coordinates
	Direction direction
}

// наполнить структуру
type foot struct {
}
