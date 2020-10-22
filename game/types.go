// модуль содержит нестондартные типы
package game

//GameScreen глобальная переменная для получения доступа к основным объектам
var GameScreen *Game

//Coordinates координаты
type Coordinates struct {
	x int
	y int
}

type direction int

const (
	up direction = iota
	down
	left
	right
)
