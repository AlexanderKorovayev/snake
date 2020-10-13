//модуль для описания основных объектов
package game

import (
	"github.com/JoelOtter/termloop"
)

//Game основной объект для рисования
type Game struct {
	termloop.Level
	//gameArea  area
	snake *Snake
	//gameFoot  foot
}

// наполнить структуру
type area struct {
	termloop.Entity
	areaBorder map[coordinates]int // для нахождения коллисий с краем поля
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
