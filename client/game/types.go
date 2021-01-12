/*
package core
модуль types
модуль содержит новые типы, которые необходимы для удобства в работе
*/

package game

// сделан доступным так как для приведения к json нужно
// что бы объекты были открытыми
// Сoordinates координаты
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

// TransportData структура, которая будет передаваться между сервером и клиентом
type TransportData struct {
	Action           interface{}
	MainObjectsCoord map[string][]Coordinates
	Info             interface{}
}
