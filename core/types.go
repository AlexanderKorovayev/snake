/*
package core
модуль types
модуль содержит новые типы, которые необходимы для удобства в работе
*/

package core

//GameScreen глобальная переменная для получения доступа к основным объектам
var GameScreen *Game

//Coordinates координаты
type Coordinates struct {
	X int
	Y int
}

type direction int

const (
	up direction = iota
	down
	left
	right
)
